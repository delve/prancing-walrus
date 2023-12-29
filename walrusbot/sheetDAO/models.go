package sheetDAO

import (
	"context"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/delve/sheetdb"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

var dbClient *sheetdb.Client
var ctx context.Context

// Initialize initializes this package.
func Initialize(spreadsheetID string, saKey []byte) error {
	ctx = context.Background()

	client, err := sheetdb.New(ctx, spreadsheetID, option.WithCredentialsJSON(saKey), option.WithScopes(sheets.SpreadsheetsScope))
	if err != nil {
		return err
	}
	dbClient = client
	loadData := func() error {
		return dbClient.LoadData(ctx)
	}
	err = backoff.Retry(loadData, backoff.WithMaxRetries(backoff.NewConstantBackOff(10*time.Second), 13))
	if err != nil {
		// retrying failed, time to die
		return err
	}

	return err
}

//go:generate sheetdb-modeler -type=Snail -children=SnailStat -test=off

// Snail is a struct of basic snail data
type Snail struct {
	SnailID   int        `json:"SnailID" db:"primarykey"`
	DiscoId   string     `json:"discoId" db:"unique"`
	GameName  string     `json:"gameName"`
	Club      string     `json:"club"`
	Server    ServerName `json:"server"`
	ServerNum int        `json:"serverNum"`
}

//go:generate sheetdb-modeler -type=SnailStat -parent=Snail -test=off

// SnailStat is a struct of snail statistics and is a child of Snail
type SnailStat struct {
	SnailID     int `json:"SnailID" db:"primarykey"`
	SnailStatID int `json:"SnailStatID" db:"primarykey"`
	// unix epoch of last record update
	Updated    int `json:"updated"`
	Leadership int `json:"leadership"`
	Art        int `json:"art"`
	Fth        int `json:"fth"`
	Fame       int `json:"fame"`
	Civ        int `json:"civ"`
	Tech       int `json:"tech"`
	Hp         int `json:"hp"`
	Atk        int `json:"atk"`
	Rush       int `json:"rush"`
	Def        int `json:"def"`
	// accept string (eg 13.0M), convert to number
	TotalPower int `json:"totalPower"`
	// 0 means not unlocked
	ZombieForm int `json:"zombieForm"`
	// 0 means not unlocked
	DemonForm int `json:"demonForm"`
	// 0 means not unlocked
	AngelForm int `json:"angelForm"`
	// 0 means not unlocked
	MutantForm int `json:"mutantForm"`
	// 0 means not unlocked
	MechaForm int `json:"mechaForm"`
	// 0 means not unlocked
	DragonForm         int `json:"dragonForm"`
	SpeciesWarEssences int `json:"speciesWarEssences"`
}
