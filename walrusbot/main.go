package main

import (
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
	"walrusbot/bot/assignment"
	botcommands "walrusbot/bot/commands"
	"walrusbot/sheetDAO"
	"walrusbot/utility/check"
	"walrusbot/utility/config"
	"walrusbot/utility/log"

	"github.com/FedorLap2006/disgolf"
	"github.com/bwmarrin/discordgo"
)

var (
	BotId string
)

func tidy() {
	log.FastLogger.Sync() // flushes buffer, if any
	config.Cleanup()      // cleans up SA key
}

func main() {
	log.Infow("config loaded", "config", config.Values)
	log.Infow("testing sheetdb for sheet DAO")
	err := sheetDAO.Initialize(config.Values.SheetId, config.Values.Secrets.GetServiceAccountKey())
	check.Err(err)

	users, err := sheetDAO.GetSnails()
	check.Err(err)
	log.Infow("Found", "users", users)

	// log.Infow("testing sheet edit from brainstorm")
	// ctx := context.Background()
	// srv, err := sheets.NewService(ctx, option.WithCredentialsJSON(config.Values.Secrets.GetServiceAccountKey()), option.WithScopes(sheets.SpreadsheetsScope))
	// check.Err(err, "Unable to retrieve Sheets client")
	// spreadsheetId := config.Values.SheetId
	// dataRange := fmt.Sprintf("'%s'!%s", "users", "A1:B1")
	// data, err := srv.Spreadsheets.Values.Get(spreadsheetId, dataRange).Do()
	// check.Err(err, "Unable to retrieve data from sheet")
	// if len(data.Values) != 1 {
	// 	log.Fatalw("inconcievable! header rows != 1", "rowsFound", len(data.Values), "data", data.Values)
	// }
	// log.Infow("retrieved sheet data", "data", data)
	//
	//
	//
	defer os.Exit(0)
	defer tidy()
	runtime.Goexit()
	//
	//
	//
	log.Infow("Inited, main starting up...")
	defer tidy()

	// check the bot is minimally functional before loading any data
	bot, err := disgolf.New(config.Values.Secrets.GetBotToken())
	check.Err(err, "failed to init disgolf")

	// initial cache of assignment data
	assignment.CacheAssignments()

	bot.Router.Register(botcommands.MyAssignment)
	bot.Router.Register(botcommands.MyAss)
	bot.Router.Register(botcommands.RefreshAssignment)

	bot.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Infow("Bot is up!")
	})
	bot.AddHandler(bot.Router.HandleInteraction)
	// lets just not respond to DMs at all for now.
	// bot.AddHandler(bot.Router.MakeMessageHandler(&disgolf.MessageHandlerConfig{
	// 	// TODO: tidy this
	// 	Prefixes:      []string{"w.", "walrus."},
	// 	MentionPrefix: true,
	// }))

	err = bot.Open()
	if err != nil {
		log.Fatalw("bot open exited with a error", "err", err)
	}
	defer bot.Close()

	err = bot.Router.Sync(bot.Session, config.Values.AppId, config.Values.ServerId)
	if err != nil {
		log.Fatalw("cannot publish commands", "err", err)
	}
	stchan := make(chan os.Signal, 1)
	signal.Notify(stchan, syscall.SIGTERM, os.Interrupt, syscall.SIGSEGV)
end:
	for {
		select {
		case <-stchan:
			break end
		default:
		}
		time.Sleep(time.Second)
	}
}
