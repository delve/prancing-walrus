package main

import (
	"encoding/json"
	"os"
	"os/signal"
	"syscall"
	"time"
	"walrusbot/botcommands"

	"github.com/FedorLap2006/disgolf"
	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

const parameterFile = "../config.json"

var (
	Log        *zap.SugaredLogger
	FastLogger *zap.Logger

	config *configStruct

	BotId string
)

type configStruct struct {
	Token     string   `json:"Token"`
	BotPrefix []string `json:"BotPrefix"`
	AppId     string   `json:"AppId"`
	ServerId  string   `json:"ServerId"`
}

func Check(err error) {
	if err != nil {
		Log.Fatalw("unhandled error", "err", err)
	}
}

func getLogger() (err error) {
	err = nil
	// TODO: NewProduction is a canned set of production-ready configs for the logger.
	//       expand to customizable configs and enable changing log level. from env var? updating via chat cmds?
	FastLogger, err = zap.NewProduction()
	if err != nil {
		return
	}
	Log = FastLogger.Sugar()
	return
}

func ReadConfig() (err error) {
	err = nil
	Log.Infow("reading config", "file", parameterFile)
	file, err := os.ReadFile(parameterFile)
	if err != nil {
		return
	}

	err = json.Unmarshal(file, &config)
	if err != nil {
		return
	}

	if config.Token == "BOT_TOKEN" {
		Log.Infow("token not found in config.json; reading from environment")
		config.Token = os.Getenv("BOT_TOKEN")
	}

	return
}

func init() {
	err := getLogger()
	if err != nil {
		panic(err)
	}
	Log.Infow("starting up...")

	err = ReadConfig()
	Check(err)
}

func main() {
	defer FastLogger.Sync() // flushes buffer, if any

	bot, err := disgolf.New(config.Token)
	// bot, err := disgolf.New("fQ9joR5zEH1xKEupw7ylSzivQCtnIoJh")
	if err != nil {
		Log.Fatalw("failed to init disgolf", "err", err)
	}
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
				Log.Infow("Middleware worked!")
				ctx.Next()
			}),
		},

		// MessageMiddlewares array executes (before? after?) a @ mention or prefixed message handler. because???
		MessageMiddlewares: []disgolf.MessageHandler{
			disgolf.MessageHandlerFunc(func(ctx *disgolf.MessageCtx) {
				Log.Infow("Message niddleware worked!", "command args", ctx.Arguments)
				ctx.Next()
			}),
		},
	})

	bot.Router.Register(botcommands.MyAssignment)

	bot.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		Log.Infow("Bot is up!")
	})
	bot.AddHandler(bot.Router.HandleInteraction)
	bot.AddHandler(bot.Router.MakeMessageHandler(&disgolf.MessageHandlerConfig{
		Prefixes:      []string{"d.", "dis.", "disgolf."},
		MentionPrefix: true,
	}))

	err = bot.Open()
	if err != nil {
		Log.Fatalw("bot open exited with a error", "err", err)
	}
	defer bot.Close()
	err = bot.Router.Sync(bot.Session, config.AppId, config.ServerId)
	if err != nil {
		Log.Fatalw("cannot publish commands", "err", err)
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
