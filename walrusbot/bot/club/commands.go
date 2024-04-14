package club

import (
	"walrusbot/utility/helpers"

	"github.com/FedorLap2006/disgolf"
	"github.com/bwmarrin/discordgo"
)

type alreadyResponded string

func (e alreadyResponded) Error() string {
	return "already responded"
}

var listMembers = func(ctx *disgolf.Ctx) {
	helpers.HandlerWrapper("club", ctx, memberList)
}

var inductSnail = func(ctx *disgolf.Ctx) {
	helpers.HandlerWrapper("club", ctx, snailInduct)
}

var kickSnail = func(ctx *disgolf.Ctx) {
	helpers.HandlerWrapper("club", ctx, snailKick)
}

var Club = &disgolf.Command{
	Name:        "club",
	Description: "Manage your club(s)",
	Type:        discordgo.ChatApplicationCommand,
	/* handlers for the base command don't appear to be necessary
	MessageHandler: disgolf.MessageHandlerFunc(func(ctx *disgolf.MessageCtx) {
		// Default handler, no subcommand selected
		_, _ = ctx.Reply("hi (default)", false)
	}),*/
	/* a handlers for the base command doesn't appear to be necessary
	Handler: disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
		_ = ctx.Respond(helpers.GetDefaultResponse(fmt.Sprintf("You have to use a subcommand your snailness. %v", subCommandList), true, ctx))
	}),*/
	SubCommands: disgolf.NewRouter([]*disgolf.Command{
		{ // list members command
			Name:        "members",
			Description: "Get a list of all the snails in your club(s).",
			Handler:     disgolf.HandlerFunc(listMembers),
		},
		{ // induct command
			Name:        "induct",
			Description: "Officers only. Induct a snail into your club.",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "snail",
					Description: "Name of the snail you want to induct",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "club",
					Description: "Optional. If you're officer of more than one club you must specify which one to induct to.",
					Required:    false,
				},
			},
			Handler: disgolf.HandlerFunc(inductSnail),
		},
		/* help subcommand is unimplemented
		{
			Name:        "help",
			Description: "Get help with the `snail` command",
			Handler: disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
				_ = ctx.Respond(helpers.GetDefaultResponse("help subcommand.", true, ctx))
			}),
		},*/
		{ // kick command
			Name:        "kick",
			Description: "Officers only. Kick a snail out of your club.",
			Options: []*discordgo.ApplicationCommandOption{{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "snail",
				Description: "Name of the snail you want to induct",
				Required:    true,
			}},
			Handler: disgolf.HandlerFunc(kickSnail),
		},
	}),
}
