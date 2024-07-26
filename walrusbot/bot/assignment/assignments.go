package assignment

import (
	"context"
	"fmt"
	"slices"
	"walrusbot/sheetDAO"
	"walrusbot/utility/check"
	"walrusbot/utility/config"
	"walrusbot/utility/log"

	"golang.org/x/oauth2/google"
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
	key                                   int
}

var dataTabs = []rosterTab{
	{club: "The One Shell",
		tabname:     "OS Roster",
		headerRange: "A1:P1",
		dataRange:   "A2:P",
		key:         0},
	{club: "You Shell Not Pass",
		tabname:     "YSNP Roster",
		headerRange: "A1:P1",
		dataRange:   "A2:P",
		key:         0},
	{club: "Zenith",
		tabname:     "Zenith Roster",
		headerRange: "A1:P1",
		dataRange:   "A2:P",
		key:         0},
}

func calculateAssignments() error {
	log.Infow("calculating all kit assignments")
	errs := []string{}

	for _, club := range dataTabs {
		clubRec, err := sheetDAO.GetClubByName(club.club)
		if err != nil {
			errs = append(errs, club.club)
			continue
		}
		err = calculateClubAssignments(clubRec.ClubID, club.club)
		if err != nil {
			errs = append(errs, club.club)
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf("club assignments failed for %v", errs)
	}
	log.Infow("kit assignments complete")
	return nil
}

func calculateClubAssignments(club int, clubname string) error {
	log.Infow("calculating kit assignments for club", "clubId", club, "clubName", clubname)
	// get club snails, sorted by leadership ascending
	clubFilter := func(snail *sheetDAO.Snail) bool { return snail.Club == club }
	simPowerSort := func(snails []*sheetDAO.Snail) {
		slices.SortStableFunc(snails,
			func(a, b *sheetDAO.Snail) int { return a.MinionSimPower - b.MinionSimPower })
	}

	snails, err := sheetDAO.GetAllSnails(sheetDAO.SnailFilter(clubFilter), sheetDAO.SnailSort(simPowerSort))
	if err != nil {
		return err
	}
	if len(snails) < 1 {
		log.Infow("no snails found to assign kits", "clubId", club, "clubName", clubname)
		return nil
	}
	// define kit counts
	kitCount := map[rune]int{
		'l': 0,
		'p': 0,
		'v': 0,
	}
	kitCount['l'] = len(snails) - 50
	if kitCount['l'] < 0 {
		kitCount['l'] = 0
	}

	if len(snails) > 25 {
		kitCount['p'] = 25
	} else {
		kitCount['p'] = len(snails)
	}

	kitCount['v'] = len(snails) - kitCount['p'] - kitCount['l']

	// make assignments
	snailCursor := 0
	rank := len(snails)
	for i := kitCount['l']; i > 0; i, snailCursor, rank = i-1, snailCursor+1, rank-1 {
		//assign to l
		snails[snailCursor].SWKit = "laborer"
		snails[snailCursor].SWKitRank = rank
		snails[snailCursor].UpdateThisSnail()
		// fmt.Printf("%s laborer %d: %d - %s\n", clubname, i, snails[snailCursor].Leadership, snails[snailCursor].SnailName)
	}
	for i := kitCount['p']; i > 0; i, snailCursor, rank = i-1, snailCursor+1, rank-1 {
		//assign to p
		snails[snailCursor].SWKit = "prospector"
		snails[snailCursor].SWKitRank = rank
		snails[snailCursor].UpdateThisSnail()
		// fmt.Printf("%s prospector %d: %d - %s\n", clubname, i, snails[snailCursor].Leadership, snails[snailCursor].SnailName)
	}
	for i := kitCount['v']; i > 0; i, snailCursor, rank = i-1, snailCursor+1, rank-1 {
		//assign to v
		snails[snailCursor].SWKit = "vanguard"
		snails[snailCursor].SWKitRank = rank
		snails[snailCursor].UpdateThisSnail()
		// fmt.Printf("%s vanguard %d: %d - %s\n", clubname, i, snails[snailCursor].Leadership, snails[snailCursor].SnailName)
	}
	return nil
}

// TODO: this should be private and part of an init function instead of called directly from main
func CacheAssignments() {
	log.Infow("Caching assignment data")

	assignments = map[string]assignment{}

	scopes := []string{
		"https://www.googleapis.com/auth/spreadsheets.readonly",
	}
	jwtConfig, err := google.JWTConfigFromJSON(config.Values.Secrets.GetServiceAccountKey(), scopes...)
	check.Err(err, "failed to create jwt config")

	ctx := context.Background()
	srv, err := sheets.NewService(ctx, option.WithHTTPClient(jwtConfig.Client(ctx)))
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
			if len(playerRecord) < 1 {
				continue
			}
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
