package assignment

import (
	"fmt"
	"strings"
	"walrusbot/utility/config"
	"walrusbot/utility/log"

	"golang.org/x/exp/slices"

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

		name := ctx.Interaction.Member.User.Username
		log.Infow("In Handler", "command", "myassignment", "channel", thisChan.Name, "user", name)
		if slices.Contains(config.Values.WarPlanningChannels, thisChan.Name) {
			// get the username
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
					Content: fmt.Sprintf("This command is only available from these channels %v", config.Values.WarPlanningChannels),
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

var RefreshAssignment = &disgolf.Command{
	Name:        "refreshassignments",
	Description: "Refresh the species war assignment cache from the data source",
	Type:        discordgo.ChatApplicationCommand,
	Handler: disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
		// TODO: replace this with a query against the managers role
		canRefresh := []string{"bionic_turkey", "iknyc", "dustinj", "gaze3", "fedcode", "mehhhhhhhhhhhhhhhhhhhhhhhhhhhhh", "vinnydev", "na000."}
		thisChan, err := ctx.Channel(ctx.Interaction.ChannelID)
		// TODO: replace this when Check is available everywhere
		if err != nil {
			panic(err)
		}

		name := ctx.Interaction.Member.User.Username
		log.Infow("In Handler", "command", "refreshassignments", "channel", thisChan.Name, "user", name)
		if slices.Contains(canRefresh, name) {
			_ = ctx.Respond(&discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Refreshing assignment cache.",
				},
			})
			CacheAssignments()
			// TODO: this doesn't come out if placed here. Should exist at the end of CacheAssignments() anyway. Make it happen.
			// _ = ctx.Respond(&discordgo.InteractionResponse{
			// 	Type: discordgo.InteractionResponseChannelMessageWithSource,
			// 	Data: &discordgo.InteractionResponseData{
			// 		Content: "Assignment cache updated.",
			// 	},
			// })
		} else {
			_ = ctx.Respond(&discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "This command is not available to this user in this context.",
				},
			})
		}
	}),
}

func getAssignmentMessage(name string) (assignMsg string) {
	// TODO: grab the manager roleid dynamically. allow config for role(s?) to mention
	assignMsg = "Sorry bud, I don't know who you are. Paging <@&1154960752004845646> for assistance."
	assign, found := assignments[name]
	if found {
		var sb strings.Builder
		sb.WriteString(fmt.Sprintf("Hi %s! ", assign.gameName))
		if assign.role == "" {
			// TODO: grab the manager roleid dynamically. allow config for role(s?) to mention
			sb.WriteString("Looks like you don't have a kit assignment yet. Paging <@&1154960752004845646> for assistance. ")
		} else {
			sb.WriteString(fmt.Sprintf("You're assigned to the %s kit. ", assign.role))
		}

		if assign.gather == "" {
			// TODO: grab the manager roleid dynamically. allow config for role(s?) to mention
			sb.WriteString("Looks like you haven't been assigned a fossil to gather yet. Please gather whatever fossil type we seem have the fewest of.\n")
		} else {
			sb.WriteString(fmt.Sprintf("Please gather fossil %s.\n", assign.gather))
		}

		if assign.canUseClamMagic {
			sb.WriteString("If it's a clam war this week please use spells as you see fit.")
		} else {
			sb.WriteString("If it's a clam war this week please use both grow spells, but none of the others.")
		}
		assignMsg = sb.String()
	}
	return
}
