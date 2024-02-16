package snail

import (
	"fmt"
	"walrusbot/utility/helpers"

	"github.com/FedorLap2006/disgolf"
	"github.com/bwmarrin/discordgo"
)

var autocomplete *discordgo.ApplicationCommandOption = &discordgo.ApplicationCommandOption{Autocomplete: true}

var subCommandList []string = []string{"add", "list", "delete", "show", "update", "help"}

var Snail = &disgolf.Command{
	Name:        "snail",
	Description: "Manage your snail(s)",
	Type:        discordgo.ChatApplicationCommand,
	MessageHandler: disgolf.MessageHandlerFunc(func(ctx *disgolf.MessageCtx) {
		// Default handler, no subcommand selected
		_, _ = ctx.Reply("hi (default)", false)
	}),
	Handler: disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
		_ = ctx.Respond(helpers.GetDefaultResponse(fmt.Sprintf("You have to use a subcommand your snailness. %v", subCommandList), true, ctx))
	}),
	// MessageMiddlewares: []disgolf.MessageHandler{
	// 	disgolf.MessageHandlerFunc(func(ctx *disgolf.MessageCtx) {
	// 		fmt.Println("middleware")
	// 		ctx.Next()
	// 	}),
	// },
	// Middlewares: []disgolf.Handler{
	// 	disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
	// 		fmt.Println("middleware")
	// 		ctx.Next()
	// 	}),
	// },
	SubCommands: disgolf.NewRouter([]*disgolf.Command{
		{
			Name:        "group",
			Description: "Subcommand group",
			// MessageMiddlewares: []disgolf.MessageHandler{
			// 	disgolf.MessageHandlerFunc(func(ctx *disgolf.MessageCtx) {
			// 		fmt.Println("group middleware")
			// 		ctx.Next()
			// 	}),
			// },
			// Middlewares: []disgolf.Handler{
			// 	disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
			// 		fmt.Println("group middleware")
			// 		ctx.Next()
			// 	}),
			// },
			SubCommands: disgolf.NewRouter([]*disgolf.Command{
				{
					Name:        "subcommand",
					Description: "Subcommand in a subcommand group",
					Handler: disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
						_ = ctx.Respond(&discordgo.InteractionResponse{
							Type: discordgo.InteractionResponseChannelMessageWithSource,
							Data: &discordgo.InteractionResponseData{Content: "hi (group)"},
						})
					}),
					MessageHandler: disgolf.MessageHandlerFunc(func(ctx *disgolf.MessageCtx) {
						_, _ = ctx.Reply("hi (group)", false)
					}),
					// MessageMiddlewares: []disgolf.MessageHandler{
					// 	disgolf.MessageHandlerFunc(func(ctx *disgolf.MessageCtx) {
					// 		fmt.Println("individual middleware")
					// 		ctx.Next()
					// 	}),
					// },
					// Middlewares: []disgolf.Handler{
					// 	disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
					// 		fmt.Println("individual middleware")
					// 		ctx.Next()
					// 	}),
					// },
				},
			}),
			MessageHandler: disgolf.MessageHandlerFunc(func(ctx *disgolf.MessageCtx) {
				_, _ = ctx.Reply("hi (group default)", false)
			}),
		},
		{
			Name:        "subcommand",
			Description: "Just a subcommand",
			Handler: disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
				_ = ctx.Respond(&discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{Content: "hi"},
				})
			}),
			MessageHandler: disgolf.MessageHandlerFunc(func(ctx *disgolf.MessageCtx) {
				_, _ = ctx.Reply("hi", false)
			}),
			MessageMiddlewares: []disgolf.MessageHandler{
				disgolf.MessageHandlerFunc(func(ctx *disgolf.MessageCtx) {
					fmt.Println("individual middleware (2nd level)")
					ctx.Next()
				}),
			},
			Middlewares: []disgolf.Handler{
				disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
					fmt.Println("individual middleware (2nd level)")
					ctx.Next()
				}),
			},
		},
	}),
}
