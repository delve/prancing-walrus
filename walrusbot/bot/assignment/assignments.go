package assignment

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"walrusbot/utility/check"
	"walrusbot/utility/config"
	"walrusbot/utility/log"

	"golang.org/x/oauth2"
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
const (
	headerRange = "Roster!A1:P1"
	dataRange   = "Roster!A2:P"
)

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
	srv, err := sheets.NewService(ctx, option.WithAPIKey(config.Values.APIKey))
	check.Err(err, "Unable to retrieve Sheets client")

	// Prints the names and majors of students in a sample spreadsheet:
	// https://docs.google.com/spreadsheets/d/1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms/edit
	spreadsheetId := config.Values.SheetId
	headerData, err := srv.Spreadsheets.Values.Get(spreadsheetId, headerRange).Do()
	check.Err(err, "Unable to retrieve data from sheet")
	if len(headerData.Values) != 1 {
		log.Fatalw("inconcievable! header rows != 1", "headerRowsFound", len(headerData.Values), "data", headerData.Values)
	}
	headers := headerData.Values[0]
	log.Infow("sheet headers retrieved", "headers", headerData.Values, "colA", headers[0])

	playerData, err := srv.Spreadsheets.Values.Get(spreadsheetId, dataRange).Do()
	check.Err(err, "Unable to retrieve data from sheet")

	if len(playerData.Values) == 0 {
		log.Fatalw("inconcievable! no player data", "playerRowsFound", len(playerData.Values), "data", playerData.Values)
	} else {
		for _, playerRecord := range playerData.Values {
			player := make(map[string]string)
			for element, val := range playerRecord {
				player[headers[element].(string)] = val.(string)
			}

			/* abstracting player name away from the data because some of them are empty in the sheet.
			this causes the array to be truncated so calling playerRecord[1] results in an
			index out of range error. if it's missing just make it blank for now.
			*/
			name := ""
			if len(playerRecord) >= 2 {
				name = playerRecord[1].(string)
			}

			assignments[playerRecord[0].(string)] = assignment{
				gameName:        name,
				role:            "",
				gather:          "",
				canUseClamMagic: false,
				data:            player,
			}
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

// OAUTH2 shite below here. may not need it.
// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalw("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalw("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalw("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}
