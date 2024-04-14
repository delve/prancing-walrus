package snail

import (
	"walrusbot/utility/helpers"

	"github.com/FedorLap2006/disgolf"
	"github.com/bwmarrin/discordgo"
)

type alreadyResponded string

func (e alreadyResponded) Error() string {
	return "already responded"
}

var integerOptionZeroValue = 0.0

var addSnail = func(ctx *disgolf.Ctx) {
	helpers.HandlerWrapper("snail", ctx, snailAdd)
}

var listSnail = func(ctx *disgolf.Ctx) {
	helpers.HandlerWrapper("snail", ctx, snailList)
}

var showSnail = func(ctx *disgolf.Ctx) {
	helpers.HandlerWrapper("snail", ctx, snailShow)
}

var updateSnail = func(ctx *disgolf.Ctx) {
	helpers.HandlerWrapper("snail", ctx, snailUpdate)
}

var Snail = &disgolf.Command{
	Name:        "snail",
	Description: "Manage your snail(s)",
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
		{ // add command
			Name:        "add",
			Description: "Add a new snail.",
			Options: []*discordgo.ApplicationCommandOption{{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "name",
				Description: "Name of the snail you want to add",
				Required:    true,
			}},
			Handler: disgolf.HandlerFunc(addSnail),
		},
		{ // list command
			Name:        "list",
			Description: "Get a list of all your snails.",
			Handler:     disgolf.HandlerFunc(listSnail),
		},
		/* Delete requires a confirmation step, making it a more complex interaction, worry about it later
		{
			Name:        "delete",
			Description: "Delete one of your snails. CAUTION: Cannot be undone.",
			Handler: disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
				_ = ctx.Respond(helpers.GetDefaultResponse("Sorry, delete isn't implemented yet. Ask Mehh for assistance.", true, ctx))
			}),
		},*/
		{ // show command
			Name:        "show",
			Description: "Show the stats of one of your snails",
			Options: []*discordgo.ApplicationCommandOption{{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "name",
				Description: "Name of the snail you want to see",
				Required:    true,
			}},
			Handler: disgolf.HandlerFunc(showSnail),
		},
		/* help subcommand is unimplemented
		{
			Name:        "help",
			Description: "Get help with the `snail` command",
			Handler: disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
				_ = ctx.Respond(helpers.GetDefaultResponse("help subcommand.", true, ctx))
			}),
		},*/
		{ // update command
			Name:        "update",
			Description: "Update your snails' stats.",
			Options:     snailStatsOptions,
			Handler:     disgolf.HandlerFunc(updateSnail),
		},
		/* setserver requires another interaction step, making it a more complex interaction, worry about it later
		{
			Name:        "setserver",
			Description: "Set what server your snail is on.",
			Handler: disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
				_ = ctx.Respond(helpers.GetDefaultResponse("Sorry, setserver isn't implemented yet. Ask Mehh for assistance.", true, ctx))
			}),
		},*/
	}),
}

var snailStatsOptions = []*discordgo.ApplicationCommandOption{
	{
		Type:        discordgo.ApplicationCommandOptionString,
		Name:        "name",
		Description: "Name of the snail to update",
		Required:    true,
	},
	{
		Type:        discordgo.ApplicationCommandOptionInteger,
		Name:        "leadership",
		Description: "Leadership",
		MinValue:    &integerOptionZeroValue,
		// MaxValue:    10,
		Required: false,
	},
	{
		Type:        discordgo.ApplicationCommandOptionInteger,
		Name:        "sw_essences",
		Description: "Species war essence count",
		MinValue:    &integerOptionZeroValue,
		// MaxValue:    10,
		Required: false,
	},
	{
		Type:        discordgo.ApplicationCommandOptionString,
		Name:        "totalpower",
		Description: "Snail's total power (EG 3.2M)",
		Required:    false,
	},
	{
		Type:        discordgo.ApplicationCommandOptionString,
		Name:        "art",
		Description: "Art",
		Required:    false,
	},
	{
		Type:        discordgo.ApplicationCommandOptionString,
		Name:        "fth",
		Description: "Faith",
		Required:    false,
	},
	{
		Type:        discordgo.ApplicationCommandOptionString,
		Name:        "fame",
		Description: "Fame",
		Required:    false,
	},
	{
		Type:        discordgo.ApplicationCommandOptionString,
		Name:        "civ",
		Description: "Civ",
		Required:    false,
	},
	{
		Type:        discordgo.ApplicationCommandOptionString,
		Name:        "tech",
		Description: "Tech",
		Required:    false,
	},
	{
		Type:        discordgo.ApplicationCommandOptionString,
		Name:        "hp",
		Description: "Hit points",
		Required:    false,
	},
	{
		Type:        discordgo.ApplicationCommandOptionString,
		Name:        "atk",
		Description: "Attack",
		Required:    false,
	},
	{
		Type:        discordgo.ApplicationCommandOptionString,
		Name:        "rush",
		Description: "Rush",
		Required:    false,
	},
	{
		Type:        discordgo.ApplicationCommandOptionString,
		Name:        "def",
		Description: "Defense",
		Required:    false,
	},
	{
		Type:        discordgo.ApplicationCommandOptionInteger,
		Name:        "zombie",
		Description: "Zombie form tier",
		MinValue:    &integerOptionZeroValue,
		// MaxValue:    10,
		Required: false,
	},
	{
		Type:        discordgo.ApplicationCommandOptionInteger,
		Name:        "demon",
		Description: "Demon form tier",
		MinValue:    &integerOptionZeroValue,
		// MaxValue:    10,
		Required: false,
	},
	{
		Type:        discordgo.ApplicationCommandOptionInteger,
		Name:        "angel",
		Description: "Angel form tier",
		MinValue:    &integerOptionZeroValue,
		// MaxValue:    10,
		Required: false,
	},
	{
		Type:        discordgo.ApplicationCommandOptionInteger,
		Name:        "mutant",
		Description: "Mutant form tier",
		MinValue:    &integerOptionZeroValue,
		// MaxValue:    10,
		Required: false,
	},
	{
		Type:        discordgo.ApplicationCommandOptionInteger,
		Name:        "mecha",
		Description: "Mecha form tier",
		MinValue:    &integerOptionZeroValue,
		// MaxValue:    10,
		Required: false,
	},
	{
		Type:        discordgo.ApplicationCommandOptionInteger,
		Name:        "dragon",
		Description: "Dragon form tier",
		MinValue:    &integerOptionZeroValue,
		// MaxValue:    10,
		Required: false,
	},
	{
		Type:        discordgo.ApplicationCommandOptionString,
		Name:        "newname",
		Description: "Change your snail's name",
		Required:    false,
	},
}
