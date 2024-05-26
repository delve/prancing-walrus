package club

import (
	"github.com/bwmarrin/discordgo"
	"github.com/zekrotja/ken"
)

type Club struct{}

var (
	_ ken.SlashCommand = (*Club)(nil)
	_ ken.DmCapable    = (*Club)(nil)
	// _ ken.AutocompleteCommand = (*Club)(nil)
)

func (c *Club) Name() string {
	return "club"
}

func (c *Club) Description() string {
	return "Manage your club(s)"
}

func (c *Club) Version() string {
	return "1.0.0"
}

func (c *Club) Type() discordgo.ApplicationCommandType {
	return discordgo.ChatApplicationCommand
}

func (c *Club) IsDmCapable() bool {
	return true
}

func (c *Club) Run(ctx ken.Context) (err error) {
	err = ctx.HandleSubCommands(
		ken.SubCommandHandler{Name: "members", Run: c.memberList},
		ken.SubCommandHandler{Name: "induct", Run: c.snailInduct},
		ken.SubCommandHandler{Name: "kick", Run: c.snailKick},
	)

	return
}

func (c *Club) Options() []*discordgo.ApplicationCommandOption {
	return []*discordgo.ApplicationCommandOption{
		{
			Type:        discordgo.ApplicationCommandOptionSubCommand,
			Name:        "members",
			Description: "Get a list of all the snails in your club(s).",
		},
		{
			Type:        discordgo.ApplicationCommandOptionSubCommand,
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
		},
		{
			Type:        discordgo.ApplicationCommandOptionSubCommand,
			Name:        "kick",
			Description: "Officers only. Kick a snail out of your club.",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "snail",
					Description: "Name of the snail you want to induct",
					Required:    true,
				},
			},
		},
	}
}
