package assignment

import (
	"context"
	"fmt"
	"walrusbot/utility/check"
	"walrusbot/utility/config"
	"walrusbot/utility/log"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

var assignments map[string]assignment

type assignment struct {
	gameName        string
	role            string
	gather          string
	canUseClamMagic bool
	data            interface{}
}

/*
	 consider making this configurable or perhaps dynamic
		 certainly the end column might need to be. not sure what to do with the sheet name
*/

type rosterTab struct {
	club, tabname, headerRange, dataRange string
}

var dataTabs = []rosterTab{
	{club: "Escargot",
		tabname:     "Roster",
		headerRange: "A1:P1",
		dataRange:   "A2:P"},
	{club: "Silken Pagoda",
		tabname:     "Silken Roster",
		headerRange: "A1:P1",
		dataRange:   "A2:P"},
}

func CacheAssignments() {
	log.Infow("Caching assignment data")

	assignments = map[string]assignment{}

	ctx := context.Background()
	// This might be needed for OAUTH2, and the commented NewService line below
	// b, err := os.ReadFile("credentials.json")
	// check.Err(err, "Unable to read client secret file")
	// If modifying these scopes, delete your previously saved token.json.
	// config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets.readonly")
	// check.Err(err, "Unable to parse client secret file to config")
	// client := getClient(config)

	// srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	srv, err := sheets.NewService(ctx, option.WithAPIKey(config.Values.Secrets.GetSheetsApiKey()))
	check.Err(err, "Unable to retrieve Sheets client")

	spreadsheetId := config.Values.SheetId
	for _, tab := range dataTabs {
		headerRange := fmt.Sprintf("'%s'!%s", tab.tabname, tab.headerRange)
		headerData, err := srv.Spreadsheets.Values.Get(spreadsheetId, headerRange).Do()
		check.Err(err, "Unable to retrieve data from sheet")
		if len(headerData.Values) != 1 {
			log.Fatalw("inconcievable! header rows != 1", "headerRowsFound", len(headerData.Values), "data", headerData.Values)
		}

		headers := headerData.Values[0]
		log.Infow("sheet headers retrieved", "club", tab.club, "headers", headerData.Values, "colA", headers[0])

		dataRange := fmt.Sprintf("'%s'!%s", tab.tabname, tab.dataRange)
		playerData, err := srv.Spreadsheets.Values.Get(spreadsheetId, dataRange).Do()
		check.Err(err, "Unable to retrieve data from sheet")

		if len(playerData.Values) == 0 {
			log.Fatalw("inconcievable! no player data", "playerRowsFound", len(playerData.Values), "data", playerData.Values)
		}

		for _, playerRecord := range playerData.Values {
			assignments[playerRecord[0].(string)] = makeAssignmentRecord(playerRecord, headers)
		}
	}

	keys := make([]string, len(assignments))

	i := 0
	for _, k := range assignments {
		keys[i] = k.gameName
		i++
	}
	log.Infow("player data cached", "total players loaded", len(assignments), "players", keys)
}

func makeAssignmentRecord(playerRecord, headers []interface{}) assignment {
	/* abstracting some values away from the data because some of them are empty in the sheet.
	   this causes the array to be truncated so calling playerRecord[1] results in an
	   index out of range error. if it's missing just make it blank for now.
	*/
	name := ""
	if len(playerRecord) >= 2 {
		name = playerRecord[1].(string)
	}

	player := make(map[string]string)
	for element, val := range playerRecord {
		// protect against "unheadered" columns
		if element < len(headers) {
			player[headers[element].(string)] = val.(string)
		} else {
			log.Warnw("unheadered column ignored", "player", name, "headerCount", len(headers), "columnNumZeroBase", element)
		}
	}

	role := ""
	if len(playerRecord) >= 4 {
		role = playerRecord[3].(string)
		if role == "" {
			role = "Prospector" // TODO: Wouldbe better if role was explicit.
		}
	}

	gather := ""
	if len(playerRecord) >= 5 {
		gather = playerRecord[4].(string)
	}

	return assignment{
		gameName:        name,
		role:            role,
		gather:          gather,
		canUseClamMagic: false,
		data:            player,
	}
}
