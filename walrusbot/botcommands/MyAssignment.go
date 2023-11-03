package botcommands

import (
	"fmt"
	"strings"

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
		fmt.Printf("In Handler channel: %v\n", thisChan.Name)
		if thisChan.Name == "prancingwalrustests" {
			// get the username
			name := ctx.Interaction.Member.User.Username
			// fmt.Printf("In MessageHandler: ctx: %v\n", ctx)
			// fmt.Printf("In MessageHandler: ctx.interaction: %v\n", ctx.Interaction)
			// fmt.Printf("In MessageHandler: ctx.interaction.Member: %v\n", ctx.Interaction.Member)
			// fmt.Printf("In MessageHandler: ctx.interaction.Member.User: %v\n", ctx.Interaction.Member.User)
			// fmt.Printf("In MessageHandler: ctx.interaction.Member.Nick: %v\n", ctx.Interaction.Member.Nick)
			_ = ctx.Respond(&discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: getAssignmentMessage(name),
				},
			})
		} else {
			_ = ctx.Respond(&discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "This has been a test. Thank you, and please send fish.",
				},
			})
		}
	}),
	MessageHandler: disgolf.MessageHandlerFunc(func(ctx *disgolf.MessageCtx) {
		_, _ = ctx.Reply(fmt.Sprintf("In MessageHandler: MessageCtx: %v\n", ctx), true)
	}),

	// Middlewares array executes (before? after?) a / command handler. because???
	Middlewares: []disgolf.Handler{
		disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
			fmt.Printf("In Middleware Ctx: %v\n", ctx)
			ctx.Next()
		}),
	},
}

func getAssignmentMessage(name string) (assignMsg string) {
	assignMsg = "Sorry, looks like I have no assignment for you. Go bug the managers and leave me to my fish."
	assign, found := assignments[name]
	if found {
		var sb strings.Builder
		sb.WriteString(fmt.Sprintf("Hi %s! You're assigned to the %s kit. Please gather fossil %s.\n", assign.gameName, assign.role, assign.gather))
		if assign.canUseClamMagic {
			sb.WriteString("If it's a clam war this week please use spells as you see fit.")
		} else {
			sb.WriteString("If it's a clam war this week please use both grow spells, but none of the others.")
		}
		assignMsg = sb.String()
	}
	return
}
