package botcommands

import (
	"fmt"

	"github.com/FedorLap2006/disgolf"
	"github.com/bwmarrin/discordgo"
)

var MyAssignment = &disgolf.Command{
	Name:        "myassignment",
	Description: "Get your species war assignments",
	Type:        discordgo.ChatApplicationCommand,
	Handler: disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
		// fmt.Printf("In Handler CTX: %v\n", ctx)
		// fmt.Printf("In Handler interaction: %v\n", ctx.Interaction.ChannelID)
		thisChan, err := ctx.Channel(ctx.Interaction.ChannelID)
		if err != nil {
			panic(err)
		}
		// fmt.Printf("In Handler err: %v\n", err)
		// fmt.Printf("In Handler channel: %v\n", thisChan.Name)
		if thisChan.Name == "prancingwalrustests" {
			_ = ctx.Respond(&discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Hi meh! You're assigned to the Prospector team and gather fossil type γ",
				},
			})
		} else {
			_ = ctx.Respond(&discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "This command is only available from the species-war channel.",
				},
			})
		}
	}),
	MessageHandler: disgolf.MessageHandlerFunc(func(ctx *disgolf.MessageCtx) {
		fmt.Printf("In MessageHandler CTX: %v\n", ctx)
		_, _ = ctx.Reply("Hi meh! You're assigned to the Prospector team and gather fossil type γ", true)
	}),

	// Middlewares array executes (before? after?) a / command handler. because???
	Middlewares: []disgolf.Handler{
		disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
			fmt.Printf("In Middleware CTX: %v\n", ctx)
			ctx.Next()
		}),
	},
}

// func getAssignmentMessage(ctx)
