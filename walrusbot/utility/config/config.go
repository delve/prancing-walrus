package config

import (
	"encoding/json"
	"os"
	"walrusbot/utility/check"
	"walrusbot/utility/log"
)

const defaultParameterFile = "../config.json"

var Values *ConfigStruct

type ConfigStruct struct {
	Token               string   `json:"Token"`
	BotPrefix           []string `json:"BotPrefix"`
	AppId               string   `json:"AppId"`
	WarPlanningChannels []string `json:"WarPlanningChannels"`
	ServerId            string   `json:"ServerId"` // Only set this if you want to limit the servers the bot talks to
	SheetId             string   `json:"SheetId"`  // google spreadsheet data source
	APIKey              string   `json:"APIKey"`   // get yer own!
}

func init() {
	err := readConfig()
	check.Err(err)
}

func readConfig() (err error) {
	err = nil
	confFile := ""
	if os.Getenv("CONFIG") == "" {
		confFile = defaultParameterFile
	} else {
		confFile = os.Getenv("CONFIG")
	}

	log.Infow("reading config", "file", confFile)
	file, err := os.ReadFile(confFile)
	if err != nil {
		return
	}

	err = json.Unmarshal(file, &Values)
	if err != nil {
		return
	}

	if Values.Token == "" {
		log.Infow("token not found in config.json; reading from environment")
		Values.Token = os.Getenv("BOT_TOKEN")
	}
	log.Infow("found token", "length", len(Values.Token))

	if Values.APIKey == "" {
		log.Infow("API key not found in config.json; reading from environment")
		Values.APIKey = os.Getenv("APIKey")
	}
	log.Infow("found key", "length", len(Values.APIKey))

	return
}
