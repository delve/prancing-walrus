package frameworkNotes

import (
	"fmt"
	"walrusbot/utility/config"

	"github.com/FedorLap2006/disgolf"
	"github.com/bwmarrin/discordgo"
)

var MyAssignment = &disgolf.Command{
	Name:        "myassignment",
	Description: "Get your species war assignments",
	Type:        discordgo.ChatApplicationCommand,
	// Handlers handle slash commands
	Handler: disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
		// ctx Respond sends a response constructed as an interaction
		_ = ctx.Respond(&discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf("This command is only available from these channels %v", config.Values.WarPlanningChannels),
			},
		})
	}),

	// Middlewares array executes (before? after?) (Message)Handler. because???
	Middlewares: []disgolf.Handler{
		disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
			fmt.Printf("In Middleware Ctx: %v\n", ctx)
			ctx.Next()
		}),
	},

	// MessageHandlers handle DMs, but only if they contain the configured prefix
	MessageHandler: disgolf.MessageHandlerFunc(func(ctx *disgolf.MessageCtx) {
		// TODO: ctx Reply is different from Respond HOW??
		_, _ = ctx.Reply(fmt.Sprintf("In MessageHandler: MessageCtx: %v\n", ctx), true)
	}),

	// MessageMiddlewares array executes (before? after?) (Message)Handler. because???
	MessageMiddlewares: []disgolf.MessageHandler{
		disgolf.MessageHandlerFunc(func(ctx *disgolf.MessageCtx) {
			fmt.Println("middleware")
			ctx.Next()
		}),
	},
}
