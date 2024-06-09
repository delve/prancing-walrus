package assignment

import (
	"walrusbot/bot/middleware/beforeMiddleware"
	"walrusbot/utility/config"
	"walrusbot/utility/helpers"

	"github.com/bwmarrin/discordgo"
	"github.com/zekrotja/ken"
)

type CalculateAssignment struct{}

var (
	_ ken.SlashCommand                 = (*CalculateAssignment)(nil)
	_ ken.DmCapable                    = (*CalculateAssignment)(nil)
	_ beforeMiddleware.RequiresOneRole = (*CalculateAssignment)(nil)
)

func (c *CalculateAssignment) Name() string {
	return "calculateassignments"
}

func (c *CalculateAssignment) Description() string {
	return "Calculate the species war assignments based on player entered data"
}

func (c *CalculateAssignment) Version() string {
	return "1.0.0"
}

func (c *CalculateAssignment) Type() discordgo.ApplicationCommandType {
	return discordgo.ChatApplicationCommand
}

func (c *CalculateAssignment) IsDmCapable() bool {
	return true
}

func (c *CalculateAssignment) Options() []*discordgo.ApplicationCommandOption {
	return []*discordgo.ApplicationCommandOption{}
}

func (c *CalculateAssignment) AllowedRoles(ctx *ken.Ctx) (roles []string) {
	roles = []string{config.Values.Roles["CanRefresh"]}
	return
}

func (c *CalculateAssignment) Run(ctx ken.Context) (err error) {
	_ = ctx.Respond(helpers.GetDefaultResponse("Calculating kit assignments.", false, ctx))
	calculateAssignments()
	err = ctx.FollowUpMessage("Calculation complete").Send().Error

	return
}
