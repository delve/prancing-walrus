package config

import (
	"encoding/json"
	"os"
	"walrusbot/utility/check"
	"walrusbot/utility/log"
)

const defaultParameterFile = "./config.json"

type ConfigStruct struct {
	Secrets             secrets
	GcpProject          string            `json:"GcpProject"`
	BotPrefix           []string          `json:"BotPrefix"`
	AppId               string            `json:"AppId"`
	WarPlanningChannels []string          `json:"WarPlanningChannels"`
	ServerId            string            `json:"ServerId"`  // Only set this if you want to limit the servers the bot talks to
	SheetId             string            `json:"SheetId"`   // google spreadsheet data source maintained by humans
	DbSheetId           string            `json:"DbSheetId"` // google spreadsheet data source maintained by bot
	Roles               map[string]string `json:"Roles"`
	Debug               bool              `json:"Debug"` // enables debug logging everywhere i know about
}

var Values *ConfigStruct

func init() {
	err := readConfig()
	check.Err(err)
}

func readConfig() (err error) {
	err = nil
	confFile := defaultParameterFile
	if os.Getenv("CONFIG") != "" {
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

	err = get_secrets()
	check.Err(err)
	return
}

func Cleanup() {
	// tidy up by deleting the SA key we created.
	deleteKey(Values.Secrets.serviceAccountKey.Name)
}
