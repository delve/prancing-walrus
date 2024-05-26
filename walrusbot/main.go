package main

import (
	"os"
	"os/signal"
	"syscall"
	"walrusbot/bot/assignment"
	botcommands "walrusbot/bot/commands"
	"walrusbot/sheetDAO"
	"walrusbot/utility/check"
	"walrusbot/utility/config"
	"walrusbot/utility/log"

	"github.com/bwmarrin/discordgo"
	"github.com/zekrotja/ken"
	"github.com/zekrotja/ken/store"
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
	if config.Values.Debug["appLogs"] {
		log.SetLevelDebug()
	}
	log.Infow("loading sheet DAO")
	check.Err(sheetDAO.Initialize(config.Values.DbSheetId, config.Values.Secrets.GetServiceAccountKey()))

	log.Infow("Inited, main starting up...")
	defer tidy()

	// check the bot is minimally functional before loading any data
	session, err := discordgo.New("Bot " + config.Values.Secrets.GetBotToken())
	check.Err(err, "failed to init disgolf")
	defer session.Close()

	session.Debug = config.Values.Debug["discgoLogs"]
	k, err := ken.New(session, ken.Options{
		// this errors because the directory doesn't exist. but i don't care right now.
		CommandStore: store.NewLocalCommandStore("./.tmp/.commandCache.json"),
		// OnCommandError is called when an error occurs
		// during middleware or command execution.
		OnCommandError: func(err error, ctx *ken.Ctx) {
			log.Errorw("Recovered panic in handler", "error", err, "Ctx", ctx)
		},

		// i will also need these!
		// OnEventError is called when any other user
		// event based error occured.
		// OnEventError func(context string, err error)
	})
	check.Err(err)

	// initial cache of assignment data
	assignment.CacheAssignments()

	check.Err(k.RegisterCommands(botcommands.Commands...))
	defer k.Unregister()

	// bot.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
	// 	log.Infow("Bot session opened")
	// })
	// bot.AddHandler(bot.Router.HandleInteraction)

	check.Err(session.Open())
	defer session.Close()

	// err = bot.Router.Sync(bot.Session, config.Values.AppId, config.Values.ServerId)
	// if err != nil {
	// 	log.Fatalw("cannot publish commands", "err", err)
	// }
	log.Infow("Bot is up!")

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sigchan
}
