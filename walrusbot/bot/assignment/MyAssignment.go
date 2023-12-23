package assignment

import (
	"fmt"
	"strings"
	"walrusbot/utility/check"
	"walrusbot/utility/config"
	"walrusbot/utility/helpers"
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
		thisChan, _ := ctx.Channel(ctx.Interaction.ChannelID)
		log.Infow("In Handler", "command", "myassignment", "channel", thisChan.Name, "user", ctx.Interaction.Member.User.Username)
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
		thisChan, _ := ctx.Channel(ctx.Interaction.ChannelID)
		log.Infow("In Handler", "command", "myass", "channel", thisChan.Name, "user", ctx.Interaction.Member.User.Username)
		requestAssignment(ctx, true)
	}),
	MessageHandler: disgolf.MessageHandlerFunc(func(ctx *disgolf.MessageCtx) {
		_, _ = ctx.Reply(fmt.Sprintf("In MessageHandler: MessageCtx: %v\n", ctx), true)
	}),
}

var RefreshAssignment = &disgolf.Command{
	Name:        "refreshassignments",
	Description: "Refresh the species war assignment cache from the data source",
	Type:        discordgo.ChatApplicationCommand,
	Handler: disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
		thisChan, err := ctx.Channel(ctx.Interaction.ChannelID)
		check.Err(err)
		log.Infow("In Handler", "command", "refreshassignments", "channel", thisChan.Name, "user", ctx.Interaction.Member.User.Username)

		refreshRoleId, err := helpers.GetRoleId(ctx, config.Values.Roles["CanRefresh"])
		if refreshRoleId == "" || err != nil {
			log.Errorw("error finding role ID", "roleTag", "CanRefresh", "configuredRole", config.Values.Roles["CanRefresh"], "error", err)
			return
		}

		if helpers.CheckroleMembership(ctx, refreshRoleId) {
			_ = ctx.Respond(helpers.GetDefaultResponse("Refreshing assignment cache.", false, ctx))
			CacheAssignments()
			// TODO: this doesn't come out if placed here. Make it happen.
			// _ = ctx.Respond(&discordgo.InteractionResponse{
			// 	Type: discordgo.InteractionResponseChannelMessageWithSource,
			// 	Data: &discordgo.InteractionResponseData{
			// 		Content: "Assignment cache updated.",
			// 	},
			// })
		} else {
			_ = ctx.Respond(helpers.GetDefaultResponse("This command is not available to this user in this context.", true, ctx))
		}
	}),
}

func requestAssignment(ctx *disgolf.Ctx, sass bool) {
	thisChan, err := ctx.Channel(ctx.Interaction.ChannelID)
	check.Err(err)

	if slices.Contains(config.Values.WarPlanningChannels, thisChan.Name) {
		_ = ctx.Respond(helpers.GetDefaultResponse(getAssignmentMessage(ctx.Interaction.Member.User.Username, sass), true, ctx))
	} else {
		_ = ctx.Respond(helpers.GetDefaultResponse(fmt.Sprintf("This command is only available from these channels %v", config.Values.WarPlanningChannels), true, ctx))
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
		sb.WriteString(closing[tone])
		assignMsg = sb.String()
	}
	return
}
