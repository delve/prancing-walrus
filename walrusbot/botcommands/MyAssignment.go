package botcommands

import (
	"fmt"
	"regexp"
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
		// TODO: replace this when Check is available everywhere
		if err != nil {
			panic(err)
		}
		// TODO: replace this when config is available everywhere
		match, err := regexp.MatchString("species-war$", thisChan.Name)
		// match, _ := regexp.MatchString(Config.WarPlanningChannelRegex, thisChan)
		// TODO: replace this when Check is available everywhere
		if err != nil {
			panic(err)
		}

		fmt.Printf("In Handler channel: %v\n", thisChan.Name)
		if match {
			// get the username
			name := ctx.Interaction.Member.User.Username
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
					Content: fmt.Sprintf("This command is only available from the %s channel.", "species-war"), // TODO: replace this when config is available everywhere (Config.WarPlanningChannelName)
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
		// sb.WriteString(fmt.Sprintf("Hi %s! You're assigned to the %s kit.\n", assign.gameName, assign.role))
		if assign.canUseClamMagic {
			sb.WriteString("If it's a clam war this week please use spells as you see fit.")
		} else {
			sb.WriteString("If it's a clam war this week please use both grow spells, but none of the others.")
		}
		assignMsg = sb.String()
	}
	return
}
