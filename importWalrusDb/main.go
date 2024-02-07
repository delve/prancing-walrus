package main

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/delve/prancing-walrus/importWalrusDb/sheetDAO"
	"github.com/delve/prancing-walrus/importWalrusDb/utility/check"
	"github.com/delve/prancing-walrus/importWalrusDb/utility/config"
	"github.com/delve/prancing-walrus/importWalrusDb/utility/log"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type rosterTab struct {
	club, tabname, headerRange, dataRange string
}

var dataTabs = []rosterTab{
	{club: "Escargot",
		tabname:     "Roster",
		headerRange: "A1:P1",
		dataRange:   "A2:C"},
	{club: "Silken Pagoda",
		tabname:     "Silken Roster",
		headerRange: "A1:P1",
		dataRange:   "A2:C"},
}

func main() {
	defer os.Exit(0)
	defer tidy()
	// manual roster sheet "1kulg2agXhCRbrdYDtYf5ggq_8U2yWkib_gxUx7mV2II"
	// test DB sheet "18pqTrKEncHH2RJ9-R3wUKnN4gEjKaXzE0KvxUaxkiXo"
	sourceSheet := "1kulg2agXhCRbrdYDtYf5ggq_8U2yWkib_gxUx7mV2II"
	destinationSheet := "18pqTrKEncHH2RJ9-R3wUKnN4gEjKaXzE0KvxUaxkiXo"

	err := sheetDAO.Initialize(destinationSheet, config.Values.Secrets.GetServiceAccountKey())
	check.Err(err)

	players, err := sheetDAO.GetPlayers()
	check.Err(err)
	log.Infow("Found DB players", "count", len(players))

	// now get a data connection to manual roster
	ctx := context.Background()
	srv, err := sheets.NewService(ctx, option.WithAPIKey(config.Values.Secrets.GetSheetsApiKey()))
	check.Err(err, "Unable to retrieve Sheets client")

	logRetry := func(err error, delay time.Duration) {
		log.Infow("snailDB operation failed, will retry", "error", err, "retryDelay", delay)
	}
	for _, tab := range dataTabs {
		headerRange := fmt.Sprintf("'%s'!%s", tab.tabname, tab.headerRange)
		headerData, err := srv.Spreadsheets.Values.Get(sourceSheet, headerRange).Do()
		check.Err(err, "Unable to retrieve data from sheet")
		if len(headerData.Values) != 1 {
			log.Fatalw("inconcievable! header rows != 1", "headerRowsFound", len(headerData.Values), "data", headerData.Values)
		}

		headers := headerData.Values[0]
		log.Infow("sheet headers retrieved", "club", tab.club, "headers", headerData.Values, "colA", headers[0])

		dataRange := fmt.Sprintf("'%s'!%s", tab.tabname, tab.dataRange)
		playerData, err := srv.Spreadsheets.Values.Get(sourceSheet, dataRange).Do()
		check.Err(err, "Unable to retrieve data from sheet")

		if len(playerData.Values) == 0 {
			log.Fatalw("inconcievable! no player data", "playerRowsFound", len(playerData.Values), "data", playerData.Values)
		}
		log.Infow("Found SS players", "count", len(playerData.Values))

		for _, playerRecord := range playerData.Values {
			if len(playerRecord[0].(string)) == 0 {
				log.Infow("player has no disco", "playerData", playerData.Values[0])
				continue
			}

			discoId := playerRecord[0].(string)
			insertPlayer := func() error {
				p, err := sheetDAO.AddPlayer(
					discoId,
				)
				log.Infow("reread player", "playerRecord", p)

				return err
			}
			err = backoff.RetryNotify(insertPlayer, backoff.WithMaxRetries(backoff.NewExponentialBackOff(), 13), logRetry)
			check.Err(err)

			p, _ := sheetDAO.GetPlayerByDiscoId(playerRecord[0].(string))
			log.Infow("reread player", "playerRecord", p)

			snailName := playerRecord[1].(string)
			leadership, err := strconv.Atoi(playerRecord[2].(string))
			if err != nil {
				leadership = 0
			}
			insertSnail := func() error {
				s, err := sheetDAO.AddSnail(
					p.PlayerID,
					int(time.Now().Unix()),
					snailName,
					tab.club,
					sheetDAO.DeepFried,
					18,
					leadership,
					0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, // values not recorded
				)
				log.Infow("inserted snail", "snailRecord", s)

				return err
			}
			err = backoff.RetryNotify(insertSnail, backoff.WithMaxRetries(backoff.NewExponentialBackOff(), 13), logRetry)
			check.Err(err)
		}
	}

	runtime.Goexit()

}

func tidy() {
	log.FastLogger.Sync() // flushes buffer, if any
	config.Cleanup()      // cleans up SA key
}
