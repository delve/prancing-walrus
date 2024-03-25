package snail

import (
	"fmt"
	"regexp"
	"strings"
	"time"
	"walrusbot/sheetDAO"
	"walrusbot/utility/helpers"
	"walrusbot/utility/log"

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
		{ // add command
			Name:        "add",
			Description: "Add a new snail.",
			Options: []*discordgo.ApplicationCommandOption{{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "name",
				Description: "Name of the snail you want to add",
				Required:    true,
			}},
			Handler: disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
				thisChan, _ := ctx.Channel(ctx.Interaction.ChannelID)
				log.Infow("In Handler", "command", "snail", "subcommand", "add", "channel", thisChan.Name, "user", ctx.Interaction.Member.User.Username, "options", ctx.Options)
				if match, err := regexp.MatchString("^[0-9a-zA-Z_-]+$", ctx.Options["name"].StringValue()); !match || err != nil {
					_ = ctx.Respond(helpers.GetDefaultResponse(fmt.Sprintf("%s doesn't look like a valid name. Just what stunt are you trying to pull here?\nOnly allowing letters, numbers, _, and -. Because I couldn't find a list of characters the game consideres valid.", ctx.Options["name"].StringValue()), true, ctx))
					return
				}
				player, err := sheetDAO.GetPlayerByDiscoId(ctx.Interaction.Member.User.Username)
				if _, isErr := err.(*sheetdb.NotFoundError); isErr {
					// create a new player
					player, err = sheetDAO.AddPlayer(ctx.Interaction.Member.User.Username)
					if err != nil {
						_ = ctx.Respond(helpers.GetDefaultResponse("Sorry, there was a problem adding you to the player table. Paging <pingCaretakerRole> to review the log", true, ctx))
						log.Errorw("Error retrieving player in /snail add", "err", err)
						return
					}
				}
				snail, err := getSnail(ctx, ctx.Options["name"].StringValue())
				if err != nil { // logged and responded in function
					return
				}
				if snail != nil {
					_ = ctx.Respond(helpers.GetDefaultResponse(fmt.Sprintf("Sorry, you already told me about that snail. Try `/snail show %s` to check up on them.\nIf this is a new snail in a different server then you can give me a nickname for it.", ctx.Options["name"].StringValue()), true, ctx))
					return
				}

				snail, err = player.AddSnail(int(time.Now().Unix()), ctx.Options["name"].StringValue(), 0, "", 0, "", 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0)
				if err != nil {
					_ = ctx.Respond(helpers.GetDefaultResponse("Sorry, there was a problem adding your snail. Paging <pingCaretakerRole> to review the log", true, ctx))
					log.Errorw("Error adding snail in /snail add", "err", err)
					return
				}
				_ = ctx.Respond(helpers.GetDefaultResponse(fmt.Sprintf("Thanks for telling me about %s! Please use `/snail update %s` to tell me more about them.", snail.SnailName, snail.SnailName), true, ctx))
			}),
		},
		{ // list command
			Name:        "list",
			Description: "Get a list of all your snails.",
			Handler: disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
				thisChan, _ := ctx.Channel(ctx.Interaction.ChannelID)
				log.Infow("In Handler", "command", "snail", "subcommand", "list", "channel", thisChan.Name, "user", ctx.Interaction.Member.User.Username)
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
				thisChan, _ := ctx.Channel(ctx.Interaction.ChannelID)
				log.Infow("In Handler", "command", "snail", "subcommand", "show", "channel", thisChan.Name, "user", ctx.Interaction.Member.User.Username, "options", ctx.Options)
				snail, err := getSnail(ctx, ctx.Options["name"].StringValue())
				if err != nil { // logged and responded in function
					return
				}
				if snail == nil {
					_ = ctx.Respond(helpers.GetDefaultResponse(fmt.Sprintf("Sorry, looks like you haven't told me about %s. Maybe you could `/snail add` them.", ctx.Options["name"].StringValue()), true, ctx))
				}
				_ = ctx.Respond(helpers.GetDefaultResponse(formatSnailStats(snail), true, ctx))
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
		{ // update command
			Name:        "update",
			Description: "Update your snails' stats.",
			Options:     snailStatsOptions,
			Handler: disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
				thisChan, _ := ctx.Channel(ctx.Interaction.ChannelID)
				log.Infow("In Handler", "command", "snail", "subcommand", "update", "channel", thisChan.Name, "user", ctx.Interaction.Member.User.Username, "options", ctx.Options)
				err := updateSnail(ctx)
				if _, isErr := err.(alreadyResponded); isErr { // response has already been sent
					return
				}
				if err != nil {
					_ = ctx.Respond(helpers.GetDefaultResponse("Sorry, there was a problem updating your snail. Paging <pingCaretakerRole> to review the log", true, ctx))
					// error already logged in updateSnail()
					return
				}
				_ = ctx.Respond(helpers.GetDefaultResponse(fmt.Sprintf("Thanks! I've updated what I know about %s. You can use `/snail show %s` to confirm it.", ctx.Options["name"].StringValue(), ctx.Options["name"].StringValue()), true, ctx))
			}),
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

func getSnails(ctx *disgolf.Ctx) ([]*sheetDAO.Snail, error) {
	var responded alreadyResponded = ""
	// get list of snails from DB by disco username ctx.Interaction.Member.User.Username
	player, err := sheetDAO.GetPlayerByDiscoId(ctx.Interaction.Member.User.Username)
	if err != nil {
		_ = ctx.Respond(helpers.GetDefaultResponse("Sorry, there was a problem looking up your player data. Paging <pingCaretakerRole> to review the log", true, ctx))
		log.Errorw("Error retrieving player in getSnails()", "err", err)
		return nil, responded
	}
	snails, err := player.GetSnails()
	if err != nil {
		_ = ctx.Respond(helpers.GetDefaultResponse("Sorry, there was a problem looking up your snail data. Paging <pingCaretakerRole> to review the log", true, ctx))
		log.Errorw("Error retrieving snails in getSnails()", "err", err)
		return nil, responded
	}
	if len(snails) == 0 {
		_ = ctx.Respond(helpers.GetDefaultResponse("Sorry, I don't know any of your snails. Maybe you should `/snail add` one.", true, ctx))
		return nil, responded
	}

	return snails, nil
}

func getSnail(ctx *disgolf.Ctx, name string) (*sheetDAO.Snail, error) {
	snails, err := getSnails(ctx)
	if err != nil {
		return nil, err
	}
	for _, snail := range snails {
		if strings.EqualFold(snail.SnailName, name) {
			return snail, nil
		}
	}
	return nil, nil
}

func formatSnailStats(snail *sheetDAO.Snail) string {
	club, _ := sheetDAO.GetClub(snail.Club)
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("This is everything I know about %s.\n", snail.SnailName))
	sb.WriteString(fmt.Sprintf("Server: %s %d\tClub: %s\n", snail.Server, snail.ServerNum, club.Name))
	sb.WriteString(fmt.Sprintf("Leadership: %d\tHoarded SW Essences: %d\n", snail.Leadership, snail.SpeciesWarEssences))
	sb.WriteString(fmt.Sprintf("Total Power: %d\n", snail.TotalPower))
	sb.WriteString(fmt.Sprintf("__AFFCT__\nArt \t%d\tFaith\t%d\nFame\t%d\tCiv\t%d\nTech \t%d\n", snail.Art, snail.Fth, snail.Fame, snail.Civ, snail.Tech))
	sb.WriteString(fmt.Sprintf("__HARD__\nHP \t%d\tAtk\t%d\nRush\t%d\tDef\t%d\n", snail.Hp, snail.Atk, snail.Rush, snail.Def))
	// custom emoji in the Snailverse server
	sb.WriteString("__Form Tiers__\n")
	sb.WriteString(fmt.Sprintf("<:zombie:1211480757382418443>\t%d\t<:demon:1211480750457888798>\t%d\n", snail.ZombieForm, snail.DemonForm))
	sb.WriteString(fmt.Sprintf("<:angel:1211480749195264010>\t%d\t<:mutant:1211480755935514674>\t%d\n", snail.AngelForm, snail.MutantForm))
	sb.WriteString(fmt.Sprintf("<:mecha:1211480754228305990>\t%d\t<:dragon:1211480752311509042>\t%d", snail.MechaForm, snail.DragonForm))

	return sb.String()
}

func updateSnail(ctx *disgolf.Ctx) error {
	var responded alreadyResponded = ""
	snail, err := getSnail(ctx, ctx.Options["name"].StringValue())
	if err != nil {
		return err
	}
	for key, option := range ctx.Options {
		switch key {
		case "name":
			// do nothing with name
		case "leadership":
			snail.Leadership = int(option.IntValue())
		case "speciesessences":
			snail.SpeciesWarEssences = int(option.IntValue())
		case "totalpower":
			value, err := helpers.DebreviateNumber(option.StringValue())
			if err != nil {
				_ = ctx.Respond(helpers.GetDefaultResponse(fmt.Sprintf("%s doesn't look like a valid number: %s", option.StringValue(), err), true, ctx))
				return responded
			}
			// discard any remaining fraction, there shouldn't be any in this context anyway
			snail.TotalPower = int(value)
		case "art":
			snail.Art = int(option.IntValue())
		case "fth":
			snail.Fth = int(option.IntValue())
		case "fame":
			snail.Fame = int(option.IntValue())
		case "civ":
			snail.Civ = int(option.IntValue())
		case "tech":
			snail.Tech = int(option.IntValue())
		case "hp":
			snail.Hp = int(option.IntValue())
		case "atk":
			snail.Atk = int(option.IntValue())
		case "rush":
			snail.Rush = int(option.IntValue())
		case "def":
			snail.Def = int(option.IntValue())
		case "zombie":
			snail.ZombieForm = int(option.IntValue())
		case "demon":
			snail.DemonForm = int(option.IntValue())
		case "angel":
			snail.AngelForm = int(option.IntValue())
		case "mutant":
			snail.MutantForm = int(option.IntValue())
		case "mecha":
			snail.MechaForm = int(option.IntValue())
		case "dragon":
			snail.DragonForm = int(option.IntValue())
		case "newname":
			//lint:ignore SA6000 looping over a map, this only triggers once
			if match, err := regexp.MatchString("^[0-9a-zA-Z_-]+$", option.StringValue()); !match || err != nil {
				_ = ctx.Respond(helpers.GetDefaultResponse(fmt.Sprintf("%s doesn't look like a valid name. Just what stunt are you trying to pull here?\nOnly allowing letters, numbers, _, and -. Because I couldn't find a list of characters the game considers valid.", option.StringValue()), true, ctx))
				return responded
			}
			sn, _ := getSnail(ctx, option.StringValue())
			if sn != nil {
				_ = ctx.Respond(helpers.GetDefaultResponse(fmt.Sprintf("Sorry, you already have a snail named %s so I can't perform your update. You can use `/snail list` to see all your snails.", option.StringValue()), true, ctx))
				return responded
			}
			snail.SnailName = option.StringValue()
		}
	}
	err = snail.UpdateThisSnail()
	if err != nil {
		log.Errorw("error updating snail", "err", err, "snail", snail, "changes", ctx.Options)
		return err
	}
	return nil
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
