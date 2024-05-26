package assignment

import (
	"walrusbot/utility/helpers"

	"github.com/bwmarrin/discordgo"
	"github.com/zekrotja/ken"
)

type MyAss struct{}

var (
	_ ken.SlashCommand = (*MyAss)(nil)
	_ ken.DmCapable    = (*MyAss)(nil)
)

func (c *MyAss) Name() string {
	return "myass"
}

func (c *MyAss) Description() string {
	return "Get your species war assignments"
}

func (c *MyAss) Version() string {
	return "1.0.0"
}

func (c *MyAss) Type() discordgo.ApplicationCommandType {
	return discordgo.ChatApplicationCommand
}

func (c *MyAss) IsDmCapable() bool {
	return true
}

func (c *MyAss) Options() []*discordgo.ApplicationCommandOption {
	return []*discordgo.ApplicationCommandOption{}
}

func (c *MyAss) Run(ctx ken.Context) (err error) {
	sass := true
	_ = ctx.Respond(helpers.GetDefaultResponse(getAssignmentMessage(ctx.User().Username, sass), true, ctx))

	return
}
