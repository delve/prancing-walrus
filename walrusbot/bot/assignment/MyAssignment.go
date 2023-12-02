package assignment

import (
	"fmt"
	"strings"
	"walrusbot/utility/check"
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
		requestAssignment(ctx, false)
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

var MyAss = &disgolf.Command{
	Name:        "myass",
	Description: "Get your species war assignments",
	Type:        discordgo.ChatApplicationCommand,
	Handler: disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
		requestAssignment(ctx, true)
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

func requestAssignment(ctx *disgolf.Ctx, sass bool) {
	thisChan, err := ctx.Channel(ctx.Interaction.ChannelID)
	check.Err(err)

	name := ctx.Interaction.Member.User.Username
	log.Infow("In Handler", "command", "myassignment", "channel", thisChan.Name, "user", name)
	if slices.Contains(config.Values.WarPlanningChannels, thisChan.Name) {
		// get the username
		_ = ctx.Respond(&discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: getAssignmentMessage(name, sass),
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
}

func getAssignmentMessage(name string, sass bool) (assignMsg string) {
	tone := 0
	if sass {
		tone = 1
	}
	assignMsg = noAssignmentTxt[tone]
	assign, found := assignments[name]
	if found {
		var sb strings.Builder
		sb.WriteString(fmt.Sprintf("Hi %s! ", assign.gameName))
		if assign.role == "" {
			sb.WriteString(noKitAssignmentTxt[tone])
		} else {
			sb.WriteString(fmt.Sprintf(kitAssignmentTxt[tone], assign.role))
		}

		if assign.gather == "" {
			sb.WriteString(noFossilAssignmentTxt[tone])
		} else {
			sb.WriteString(fmt.Sprintf(fossilAssignmentTxt[tone], assign.gather))
		}

		if assign.canUseClamMagic {
			sb.WriteString(canUseClamMagicTxt[tone])
		} else {
			sb.WriteString(canNotUseClamMagicTxt[tone])
		}
		assignMsg = sb.String()
	}
	return
}
