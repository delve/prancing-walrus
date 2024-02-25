package snail

import (
	"fmt"
	"regexp"
	"strings"
	"time"
	"walrusbot/sheetDAO"
	"walrusbot/utility/helpers"

	"github.com/FedorLap2006/disgolf"
	"github.com/bwmarrin/discordgo"
	"github.com/delve/sheetdb"
)

type alreadyResponded string

func (e alreadyResponded) Error() string {
	return "already responded"
}

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
			Options: []*discordgo.ApplicationCommandOption{{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "name",
				Description: "Name of the snail you want to add",
				Required:    true,
			}},
			Handler: disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
				if match, err := regexp.MatchString("^[0-9a-zA-Z_-]$", ctx.Options["name"].Value.(string)); !match || err != nil {
					_ = ctx.Respond(helpers.GetDefaultResponse(fmt.Sprintf("%s doesn't look like a valid name. Just what stunt are you trying to pull here?", ctx.Options["name"].Value), true, ctx))
					return
				}
				player, err := sheetDAO.GetPlayerByDiscoId(ctx.Interaction.Member.User.Username)
				if _, isErr := err.(*sheetdb.NotFoundError); isErr {
					// create a new player
					player, err = sheetDAO.AddPlayer(ctx.Interaction.Member.User.Username)
					if err != nil {
						_ = ctx.Respond(helpers.GetDefaultResponse("Sorry, there was a problem adding you to the player table. Paging <pingCaretakerRole> for assistance.", true, ctx))
						return
					}
				}
				snails, err := player.GetSnails()
				if err != nil {
					_ = ctx.Respond(helpers.GetDefaultResponse("Sorry, there was a problem looking up your snail data. Paging <pingCaretakerRole> for assistance.", true, ctx))
					return
				}
				for _, snail := range snails {
					if snail.SnailName == ctx.Options["name"].Value {
						_ = ctx.Respond(helpers.GetDefaultResponse(fmt.Sprintf("Sorry, you already told me about that snail. Try `/snail show %s` to check up on them.\nIf this is a new snail in a different server then you can give me a nickname for it.", ctx.Options["name"].Value), true, ctx))
						return
					}
				}

				snail, err := player.AddSnail(int(time.Now().Unix()), ctx.Options["name"].Value.(string), "", "", 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0)
				_ = ctx.Respond(helpers.GetDefaultResponse("add subcommand.", true, ctx))
			}),
		},
		{ // list command
			Name:        "list",
			Description: "Get a list of all your snails.",
			Handler: disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
				snails, err := getSnails(ctx)
				if _, isErr := err.(alreadyResponded); isErr { // response has already been sent
					return
				}

				var sb strings.Builder
				sb.WriteString(fmt.Sprintf("Hi %s! These are the snails I know about.\n", ctx.Interaction.Member.User.Username))
				for _, snail := range snails {
					sb.WriteString(fmt.Sprintf("%s on %s %d\n", snail.SnailName, snail.Server, snail.ServerNum))
				}
				_ = ctx.Respond(helpers.GetDefaultResponse(sb.String(), true, ctx))
			}),
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
			Handler: disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
				snails, err := getSnails(ctx)
				if _, isErr := err.(alreadyResponded); isErr { // response has already been sent
					return
				}
				for _, snail := range snails {
					if snail.SnailName == ctx.Options["name"].Value {
						_ = ctx.Respond(helpers.GetDefaultResponse(formatSnailStats(snail), true, ctx))
						return
					}
				}
				_ = ctx.Respond(helpers.GetDefaultResponse(fmt.Sprintf("Sorry, looks like you haven't told me about %s. Maybe you could `/snail add` them.", ctx.Options["name"].Value), true, ctx))
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
	}),
}

func getSnails(ctx *disgolf.Ctx) ([]*sheetDAO.Snail, error) {
	var responded alreadyResponded = ""
	// get list of snails from DB by disco username ctx.Interaction.Member.User.Username
	player, err := sheetDAO.GetPlayerByDiscoId(ctx.Interaction.Member.User.Username)
	if err != nil {
		_ = ctx.Respond(helpers.GetDefaultResponse("Sorry, there was a problem looking up your player data. Paging <pingCaretakerRole> for assistance.", true, ctx))
		return nil, responded
	}
	snails, err := player.GetSnails()
	if err != nil {
		_ = ctx.Respond(helpers.GetDefaultResponse("Sorry, there was a problem looking up your snail data. Paging <pingCaretakerRole> for assistance.", true, ctx))
		return nil, responded
	}
	if len(snails) == 0 {
		_ = ctx.Respond(helpers.GetDefaultResponse("Sorry, I don't know any of your snails. Maybe you should `/snail add` one.", true, ctx))
		return nil, responded
	}

	return snails, nil
}

func formatSnailStats(snail *sheetDAO.Snail) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("This is everything I know about %s.\n", snail.SnailName))
	sb.WriteString(fmt.Sprintf("Server: %s %d\tClub: %s\n", snail.Server, snail.ServerNum, snail.Club))
	sb.WriteString(fmt.Sprintf("Leadership: %d\tHoarded SW Essences: %d\n", snail.Leadership, snail.SpeciesWarEssences))
	sb.WriteString(fmt.Sprintf("Total Power: %d\n", snail.TotalPower))
	sb.WriteString(fmt.Sprintf("__AFFCT__\nArt \t%d\tFaith\t%d\nFame\t%d\tCiv\t%d\nTech \t%d\n", snail.Art, snail.Fth, snail.Fame, snail.Civ, snail.Tech))
	sb.WriteString(fmt.Sprintf("__HARD__\nHP \t%d\tAtk\t%d\nRush\t%d\tDef\t%d\n", snail.Hp, snail.Atk, snail.Rush, snail.Def))
	// custom emoji in the Snailverse server
	sb.WriteString(fmt.Sprintf("__Form Tiers__\n:zombie~1:\t%d\t:demon:\t%d\n:angel~1:\t%d\t:mutant:\t%d\n:mecha:\t%d\t:dragon~1:\t%d", snail.ZombieForm, snail.DemonForm, snail.AngelForm, snail.MutantForm, snail.MechaForm, snail.DragonForm))

	return sb.String()
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
		Required:    false,
	},
	{
		Type:        discordgo.ApplicationCommandOptionString,
		Name:        "server",
		Description: "Game server name",
		Required:    false,
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
