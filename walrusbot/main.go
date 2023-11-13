package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"
	"walrusbot/bot/assignment"
	botcommands "walrusbot/bot/commands"
	"walrusbot/utility/check"
	"walrusbot/utility/config"
	"walrusbot/utility/log"

	"github.com/FedorLap2006/disgolf"
	"github.com/bwmarrin/discordgo"
)

var (
	BotId string
)

func main() {
	log.Infow("Inited, main starting up...")
	defer log.FastLogger.Sync() // flushes buffer, if any

	// check the bot is minimally functional before loading any data
	bot, err := disgolf.New(config.Values.Token)
	check.Err(err, "failed to init disgolf")

	// initial cache of assignment data
	assignment.CacheAssignments()

	bot.Router.Register(&disgolf.Command{
		Name:        "ping",
		Description: "Ping it!",
		Type:        discordgo.ChatApplicationCommand,
		// Handler responds to / commands directly
		Handler: disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
			_ = ctx.Respond(&discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Hi, I'm a bot built on Disgolf library. :O",
				},
			})
		}),
		// MessageHandler responds to @ mention or prefixed message commands
		MessageHandler: disgolf.MessageHandlerFunc(func(ctx *disgolf.MessageCtx) {
			_, _ = ctx.Reply("Hi, I'm a bot built on Disgolf library", true)
		}),

		// Middlewares array executes (before? after?) a / command handler. because???
		Middlewares: []disgolf.Handler{
			disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
				log.Infow("Middleware worked!")
				ctx.Next()
			}),
		},

		// MessageMiddlewares array executes (before? after?) a @ mention or prefixed message handler. because???
		MessageMiddlewares: []disgolf.MessageHandler{
			disgolf.MessageHandlerFunc(func(ctx *disgolf.MessageCtx) {
				log.Infow("Message niddleware worked!", "command args", ctx.Arguments)
				ctx.Next()
			}),
		},
	})

	bot.Router.Register(botcommands.MyAssignment)

	bot.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Infow("Bot is up!")
	})
	bot.AddHandler(bot.Router.HandleInteraction)
	bot.AddHandler(bot.Router.MakeMessageHandler(&disgolf.MessageHandlerConfig{
		Prefixes:      []string{"d.", "dis.", "disgolf."},
		MentionPrefix: true,
	}))

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
