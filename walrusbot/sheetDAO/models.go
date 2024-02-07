package sheetDAO

import (
	"context"
	"time"
	"walrusbot/utility/log"

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
	logRetry := func(err error, delay time.Duration) {
		log.Infow("snailDB loaddata failed, will retry", "error", err, "retryDelay", delay)
	}
	err = backoff.RetryNotify(loadData, backoff.WithMaxRetries(backoff.NewConstantBackOff(10*time.Second), 13), logRetry)
	if err != nil {
		// retrying failed, time to die
		return err
	}

	return err
}

//go:generate sheetdb-modeler -type=Player -children=Snail -test=off

// Snail is a struct of basic snail data
type Player struct {
	PlayerID int    `json:"playerID" db:"primarykey"`
	DiscoId  string `json:"discoId" db:"unique"`
}

//go:generate sheetdb-modeler -type=Snail -parent=Player -test=off

// SnailStat is a struct of snail statistics and is a child of Snail
type Snail struct {
	PlayerID int `json:"playerID" db:"primarykey"`
	SnailID  int `json:"SnailID" db:"primarykey"`
	// unix epoch of last record update
	Updated    int        `json:"updated"`
	SnailName  string     `json:"snailName"`
	Club       string     `json:"club"`
	Server     ServerName `json:"server"`
	ServerNum  int        `json:"serverNum"`
	Leadership int        `json:"leadership"`
	Art        int        `json:"art"`
	Fth        int        `json:"fth"`
	Fame       int        `json:"fame"`
	Civ        int        `json:"civ"`
	Tech       int        `json:"tech"`
	Hp         int        `json:"hp"`
	Atk        int        `json:"atk"`
	Rush       int        `json:"rush"`
	Def        int        `json:"def"`
	// accept string (eg 13.0M), convert to number in frontend
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

func (s *Snail) AddThisSnail() error {
	_, err := AddSnail(s.PlayerID,
		int(time.Now().Unix()),
		s.SnailName,
		s.Club,
		ServerName(s.Server),
		s.ServerNum,
		s.Leadership,
		s.Art,
		s.Fth,
		s.Fame,
		s.Civ,
		s.Tech,
		s.Hp,
		s.Atk,
		s.Rush,
		s.Def,
		s.TotalPower,
		s.ZombieForm,
		s.DemonForm,
		s.AngelForm,
		s.MutantForm,
		s.MechaForm,
		s.DragonForm,
		s.SpeciesWarEssences,
	)
	return err
}

func (s *Snail) UpdateThisSnail() error {
	_, err := UpdateSnail(s.PlayerID,
		s.SnailID,
		int(time.Now().Unix()),
		s.SnailName,
		s.Club,
		ServerName(s.Server),
		s.ServerNum,
		s.Leadership,
		s.Art,
		s.Fth,
		s.Fame,
		s.Civ,
		s.Tech,
		s.Hp,
		s.Atk,
		s.Rush,
		s.Def,
		s.TotalPower,
		s.ZombieForm,
		s.DemonForm,
		s.AngelForm,
		s.MutantForm,
		s.MechaForm,
		s.DragonForm,
		s.SpeciesWarEssences,
	)
	return err
}
