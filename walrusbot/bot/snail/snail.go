package snail

import (
	"walrusbot/utility/helpers"

	"github.com/FedorLap2006/disgolf"
	"github.com/bwmarrin/discordgo"
)

// var autocomplete *discordgo.ApplicationCommandOption = &discordgo.ApplicationCommandOption{Autocomplete: true}

// var subCommandList []string = []string{"add", "list", "delete", "show", "update", "help"}

var integerOptionZeroValue = 0.0

var Snail = &disgolf.Command{
	Name:        "snail",
	Description: "Manage your snail(s)",
	Type:        discordgo.ChatApplicationCommand,
	/* handlers for the base command don't appear to be necessary
	MessageHandler: disgolf.MessageHandlerFunc(func(ctx *disgolf.MessageCtx) {
		// Default handler, no subcommand selected
		_, _ = ctx.Reply("hi (default)", false)
	}),*/
	/* a handler for the base command doesn't appear to be necessary
	Handler: disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
		_ = ctx.Respond(helpers.GetDefaultResponse(fmt.Sprintf("You have to use a subcommand your snailness. %v", subCommandList), true, ctx))
	}),*/
	SubCommands: disgolf.NewRouter([]*disgolf.Command{
		{
			Name:        "add",
			Description: "Add a new snail.",
			Handler: disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
				_ = ctx.Respond(helpers.GetDefaultResponse("add subcommand.", true, ctx))
			}),
		},
		{
			Name:        "list",
			Description: "Get a list of all your snails.",
			Handler: disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
				_ = ctx.Respond(helpers.GetDefaultResponse("list subcommand.", true, ctx))
			}),
		},
		{
			Name:        "delete",
			Description: "Delete one of your snails. CAUTION: Cannot be undone.",
			Handler: disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
				_ = ctx.Respond(helpers.GetDefaultResponse("delete subcommand.", true, ctx))
			}),
		},
		{
			Name:        "show",
			Description: "Show the stats of one of your snails",
			Handler: disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
				_ = ctx.Respond(helpers.GetDefaultResponse("show subcommand.", true, ctx))
			}),
		},

		/* help subcommand is unimplemented
		{
			Name:        "help",
			Description: "Get help with the `snail` command",
			Handler: disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
				_ = ctx.Respond(helpers.GetDefaultResponse("help subcommand.", true, ctx))
			}),
		},*/
		{
			Name:        "update",
			Description: "Update your snails' stats.",
			Options:     updateOptions,
			Handler: disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
				_ = ctx.Respond(helpers.GetDefaultResponse("show subcommand.", true, ctx))
			}),
		},

		{
			Name:        "olddate",
			Description: "Update your snails' stats.",
			SubCommands: disgolf.NewRouter([]*disgolf.Command{
				{
					Name:        "leadership",
					Description: "Update your snail's leadership",
					Handler: disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
						_ = ctx.Respond(helpers.GetDefaultResponse("update leadership subcommand.", true, ctx))
					}),
				},
				{
					Name:        "speciesessences",
					Description: "Update how many species war essences your snail has",
					Handler: disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
						_ = ctx.Respond(helpers.GetDefaultResponse("update s_war_essences subcommand.", true, ctx))
					}),
				},
				{
					Name:        "totalpower",
					Description: "Update your snail's total power",
					Handler: disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
						// accept string (eg 13.0M), convert to number in frontend
						_ = ctx.Respond(helpers.GetDefaultResponse("update totalpower subcommand.", true, ctx))
					}),
				},
				{
					Name:        "art",
					Description: "Update your snail's Art",
					Handler: disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
						_ = ctx.Respond(helpers.GetDefaultResponse("update Art subcommand.", true, ctx))
					}),
				},
				{
					Name:        "fth",
					Description: "Update your snail's Fth",
					Handler: disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
						_ = ctx.Respond(helpers.GetDefaultResponse("update Fth subcommand.", true, ctx))
					}),
				},
				{
					Name:        "fame",
					Description: "Update your snail's Fame",
					Handler: disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
						_ = ctx.Respond(helpers.GetDefaultResponse("update Fame subcommand.", true, ctx))
					}),
				},
				{
					Name:        "civ",
					Description: "Update your snail's Civ",
					Handler: disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
						_ = ctx.Respond(helpers.GetDefaultResponse("update Civ subcommand.", true, ctx))
					}),
				},
				{
					Name:        "tech",
					Description: "Update your snail's Tech",
					Handler: disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
						_ = ctx.Respond(helpers.GetDefaultResponse("update Tech subcommand.", true, ctx))
					}),
				},
				{
					Name:        "hp",
					Description: "Update your snail's Hp",
					Handler: disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
						_ = ctx.Respond(helpers.GetDefaultResponse("update Hp subcommand.", true, ctx))
					}),
				},
				{
					Name:        "atk",
					Description: "Update your snail's Atk",
					Handler: disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
						_ = ctx.Respond(helpers.GetDefaultResponse("update Atk subcommand.", true, ctx))
					}),
				},
				{
					Name:        "rush",
					Description: "Update your snail's Rush",
					Handler: disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
						_ = ctx.Respond(helpers.GetDefaultResponse("update Rush subcommand.", true, ctx))
					}),
				},
				{
					Name:        "def",
					Description: "Update your snail's Def",
					Handler: disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
						_ = ctx.Respond(helpers.GetDefaultResponse("update Def subcommand.", true, ctx))
					}),
				},

				/* Club update disabled until i work out a way to deal with it
				{
					Name:        "Club",
					Description: "Update your snail's Club membership",
					Handler: disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
						_ = ctx.Respond(helpers.GetDefaultResponse("update Club subcommand.", true, ctx))
					}),
				},*/

				{
					Name:        "zombie",
					Description: "Update your snail's Zombie form; 0 means it isn't unlocked yet",
					Handler: disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
						_ = ctx.Respond(helpers.GetDefaultResponse("update Zombie subcommand.", true, ctx))
					}),
				},
				{
					Name:        "demon",
					Description: "Update your snail's Demon form; 0 means it isn't unlocked yet",
					Handler: disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
						_ = ctx.Respond(helpers.GetDefaultResponse("update Demon subcommand.", true, ctx))
					}),
				},
				{
					Name:        "angel",
					Description: "Update your snail's Angel form; 0 means it isn't unlocked yet",
					Handler: disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
						_ = ctx.Respond(helpers.GetDefaultResponse("update Angel subcommand.", true, ctx))
					}),
				},
				{
					Name:        "mutant",
					Description: "Update your snail's Mutant form; 0 means it isn't unlocked yet",
					Handler: disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
						_ = ctx.Respond(helpers.GetDefaultResponse("update Mutant subcommand.", true, ctx))
					}),
				},
				{
					Name:        "mecha",
					Description: "Update your snail's Mecha form; 0 means it isn't unlocked yet",
					Handler: disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
						_ = ctx.Respond(helpers.GetDefaultResponse("update Mecha subcommand.", true, ctx))
					}),
				},
				{
					Name:        "dragon",
					Description: "Update your snail's Zombie form; 0 means it isn't unlocked yet",
					Handler: disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
						_ = ctx.Respond(helpers.GetDefaultResponse("update Dragon subcommand.", true, ctx))
					}),
				},
				{
					Name:        "name",
					Description: "Update your snail's name",
					Handler: disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
						_ = ctx.Respond(helpers.GetDefaultResponse("update SnailName subcommand.", true, ctx))
					}),
				},
				{
					Name:        "server",
					Description: "Move your snail to a different game server",
					Handler: disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
						_ = ctx.Respond(helpers.GetDefaultResponse("update Server subcommand.", true, ctx))
					}),
				},
				{
					Name:        "servernum",
					Description: "Move your snail to a different game server number",
					Handler: disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
						_ = ctx.Respond(helpers.GetDefaultResponse("update ServerNum subcommand.", true, ctx))
					}),
				},
			}),
		},
	}),
}

var updateOptions = []*discordgo.ApplicationCommandOption{
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
		Name:        "speciesessences",
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
		Type:        discordgo.ApplicationCommandOptionInteger,
		Name:        "art",
		Description: "Art",
		MinValue:    &integerOptionZeroValue,
		// MaxValue:    10,
		Required: false,
	},
	{
		Type:        discordgo.ApplicationCommandOptionInteger,
		Name:        "fth",
		Description: "Faith",
		MinValue:    &integerOptionZeroValue,
		// MaxValue:    10,
		Required: false,
	},
	{
		Type:        discordgo.ApplicationCommandOptionInteger,
		Name:        "fame",
		Description: "Fame",
		MinValue:    &integerOptionZeroValue,
		// MaxValue:    10,
		Required: false,
	},
	{
		Type:        discordgo.ApplicationCommandOptionInteger,
		Name:        "civ",
		Description: "Civ",
		MinValue:    &integerOptionZeroValue,
		// MaxValue:    10,
		Required: false,
	},
	{
		Type:        discordgo.ApplicationCommandOptionInteger,
		Name:        "tech",
		Description: "Tech",
		MinValue:    &integerOptionZeroValue,
		// MaxValue:    10,
		Required: false,
	},
	{
		Type:        discordgo.ApplicationCommandOptionInteger,
		Name:        "hp",
		Description: "Hit points",
		MinValue:    &integerOptionZeroValue,
		// MaxValue:    10,
		Required: false,
	},
	{
		Type:        discordgo.ApplicationCommandOptionInteger,
		Name:        "atk",
		Description: "Attack",
		MinValue:    &integerOptionZeroValue,
		// MaxValue:    10,
		Required: false,
	},
	{
		Type:        discordgo.ApplicationCommandOptionInteger,
		Name:        "rush",
		Description: "Rush",
		MinValue:    &integerOptionZeroValue,
		// MaxValue:    10,
		Required: false,
	},
	{
		Type:        discordgo.ApplicationCommandOptionInteger,
		Name:        "def",
		Description: "Defense",
		MinValue:    &integerOptionZeroValue,
		// MaxValue:    10,
		Required: false,
	},
	{
		Type:        discordgo.ApplicationCommandOptionInteger,
		Name:        "zombie",
		Description: "Zombie form tier)",
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
		Required: false,
	},
	{
		Type:        discordgo.ApplicationCommandOptionString,
		Name:        "server",
		Description: "Game server name",
		Required: false,
	},
	{
		Type:        discordgo.ApplicationCommandOptionInteger,
		Name:        "servernum",
		Description: "Game server number",
		MinValue:    &integerOptionZeroValue,
		// MaxValue:    10,
		Required: false,
	},

}
