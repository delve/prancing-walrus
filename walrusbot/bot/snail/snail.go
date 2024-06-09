package snail

import (
	"strings"
	"walrusbot/sheetDAO"
	"walrusbot/utility/helpers"

	"github.com/bwmarrin/discordgo"
	"github.com/zekrotja/ken"
)

type Snail struct{}

var (
	_ ken.SlashCommand        = (*Snail)(nil)
	_ ken.DmCapable           = (*Snail)(nil)
	_ ken.AutocompleteCommand = (*Snail)(nil)
)

func (c *Snail) Name() string {
	return "snail"
}

func (c *Snail) Description() string {
	return "Manage your snail(s)"
}

func (c *Snail) Version() string {
	return "1.0.0"
}

func (c *Snail) Type() discordgo.ApplicationCommandType {
	return discordgo.ChatApplicationCommand
}

func (c *Snail) IsDmCapable() bool {
	return true
}

func (c *Snail) Options() []*discordgo.ApplicationCommandOption {
	return []*discordgo.ApplicationCommandOption{
		{
			Type:        discordgo.ApplicationCommandOptionSubCommand,
			Name:        "add",
			Description: "Add a new snail.",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "name",
					Description: "Name of the snail you want to add",
					Required:    true,
				},
			},
		},
		{
			Type:        discordgo.ApplicationCommandOptionSubCommand,
			Name:        "list",
			Description: "Get a list of all your snails.",
			Options:     []*discordgo.ApplicationCommandOption{},
		},
		{
			Type:        discordgo.ApplicationCommandOptionSubCommand,
			Name:        "show",
			Description: "Show the stats of one of your snails",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:         discordgo.ApplicationCommandOptionString,
					Name:         "name",
					Description:  "Name of the snail you want to see",
					Required:     true,
					Autocomplete: true,
				},
			},
		},
		{
			Type:        discordgo.ApplicationCommandOptionSubCommand,
			Name:        "update",
			Description: "Update your snails' stats.",
			Options:     snailStatsOptions,
		},
		/* Delete requires a confirmation step, making it a more complex interaction, worry about it later
		{
			Name:        "delete",
			Description: "Delete one of your snails. CAUTION: Cannot be undone.",
			Handler: disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
				_ = ctx.Respond(helpers.GetDefaultResponse("Sorry, delete isn't implemented yet. Ask Mehh for assistance.", true, ctx))
			}),
		},*/
		/* setserver requires another interaction step, making it a more complex interaction, worry about it later
		{
			Name:        "setserver",
			Description: "Set what server your snail is on.",
			Handler: disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
				_ = ctx.Respond(helpers.GetDefaultResponse("Sorry, setserver isn't implemented yet. Ask Mehh for assistance.", true, ctx))
			}),
		},*/
		/* help subcommand is unimplemented
		{
			Name:        "help",
			Description: "Get help with the `snail` command",
			Handler: disgolf.HandlerFunc(func(ctx *disgolf.Ctx) {
				_ = ctx.Respond(helpers.GetDefaultResponse("help subcommand.", true, ctx))
			}),
		},*/
	}
}

func (c *Snail) Autocomplete(ctx *ken.AutocompleteContext) ([]*discordgo.ApplicationCommandOptionChoice, error) {
	input, ok := ctx.SubCommand().GetInput("name")

	if !ok {
		return nil, nil
	}

	player, err := sheetDAO.GetPlayerByDiscoId(ctx.User().Username)
	if err != nil {
		return nil, nil
	}

	snails, err := sheetDAO.GetSnails(player.PlayerID)
	if err != nil {
		return nil, nil
	}

	choices := make([]*discordgo.ApplicationCommandOptionChoice, 0, len(snails))
	input = strings.ToLower(input)

	for _, snail := range snails {
		if strings.HasPrefix(strings.ToLower(snail.SnailName), input) {
			choices = append(choices, &discordgo.ApplicationCommandOptionChoice{
				Name:  snail.SnailName,
				Value: snail.SnailName,
			})
		}
	}

	return choices, nil
}

func (c *Snail) Run(ctx ken.Context) (err error) {
	err = ctx.HandleSubCommands(
		ken.SubCommandHandler{Name: "add", Run: c.add},
		ken.SubCommandHandler{Name: "list", Run: c.list},
		ken.SubCommandHandler{Name: "show", Run: c.show},
		ken.SubCommandHandler{Name: "update", Run: c.update},
	)

	return
}

func (c *Snail) add(ctx ken.SubCommandContext) (err error) {
	snailAdd(ctx)
	return
}

func (c *Snail) list(ctx ken.SubCommandContext) (err error) {
	snailList(ctx)
	return
}

func (c *Snail) show(ctx ken.SubCommandContext) (err error) {
	snailShow(ctx)
	return
}

func (c *Snail) update(ctx ken.SubCommandContext) (err error) {
	snailUpdate(ctx)
	return
}

var snailStatsOptions = []*discordgo.ApplicationCommandOption{
	{
		Type:         discordgo.ApplicationCommandOptionString,
		Name:         "name",
		Description:  "Name of the snail to update",
		Required:     true,
		Autocomplete: true,
	},
	{
		Type:        discordgo.ApplicationCommandOptionInteger,
		Name:        "leadership",
		Description: "Leadership",
		MinValue:    &helpers.IntegerOptionZeroValue,
		// MaxValue:    10,
		Required: false,
	},
	{
		Type:        discordgo.ApplicationCommandOptionInteger,
		Name:        "sw_essences",
		Description: "Species war essence count",
		MinValue:    &helpers.IntegerOptionZeroValue,
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
		MinValue:    &helpers.IntegerOptionZeroValue,
		// MaxValue:    10,
		Required: false,
	},
	{
		Type:        discordgo.ApplicationCommandOptionInteger,
		Name:        "demon",
		Description: "Demon form tier",
		MinValue:    &helpers.IntegerOptionZeroValue,
		// MaxValue:    10,
		Required: false,
	},
	{
		Type:        discordgo.ApplicationCommandOptionInteger,
		Name:        "angel",
		Description: "Angel form tier",
		MinValue:    &helpers.IntegerOptionZeroValue,
		// MaxValue:    10,
		Required: false,
	},
	{
		Type:        discordgo.ApplicationCommandOptionInteger,
		Name:        "mutant",
		Description: "Mutant form tier",
		MinValue:    &helpers.IntegerOptionZeroValue,
		// MaxValue:    10,
		Required: false,
	},
	{
		Type:        discordgo.ApplicationCommandOptionInteger,
		Name:        "mecha",
		Description: "Mecha form tier",
		MinValue:    &helpers.IntegerOptionZeroValue,
		// MaxValue:    10,
		Required: false,
	},
	{
		Type:        discordgo.ApplicationCommandOptionInteger,
		Name:        "dragon",
		Description: "Dragon form tier",
		MinValue:    &helpers.IntegerOptionZeroValue,
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
		Name:        "simpower",
		Description: "Snail's minion sim power (EG 3.2M)",
		Required:    false,
	},
}
