package assignment

import (
	"walrusbot/utility/helpers"

	"github.com/bwmarrin/discordgo"
	"github.com/zekrotja/ken"
)

type MyAssignment struct{}

var (
	_ ken.SlashCommand = (*MyAssignment)(nil)
	_ ken.DmCapable    = (*MyAssignment)(nil)
)

func (c *MyAssignment) Name() string {
	return "myassignment"
}

func (c *MyAssignment) Description() string {
	return "Get your species war assignments"
}

func (c *MyAssignment) Version() string {
	return "1.0.0"
}

func (c *MyAssignment) Type() discordgo.ApplicationCommandType {
	return discordgo.ChatApplicationCommand
}

func (c *MyAssignment) IsDmCapable() bool {
	return true
}

func (c *MyAssignment) Options() []*discordgo.ApplicationCommandOption {
	return []*discordgo.ApplicationCommandOption{}
}

func (c *MyAssignment) Run(ctx ken.Context) (err error) {
	sass := false
	_ = ctx.Respond(helpers.GetDefaultResponse(getAssignmentMessage(ctx.User().Username, sass), true, ctx))

	return
}
