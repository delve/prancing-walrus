package assignment

import (
	"walrusbot/utility/config"
	"walrusbot/utility/helpers"
	"walrusbot/utility/log"

	"github.com/bwmarrin/discordgo"
	"github.com/zekrotja/ken"
)

type CalculateAssignment struct{}

var (
	_ ken.SlashCommand = (*CalculateAssignment)(nil)
	_ ken.DmCapable    = (*CalculateAssignment)(nil)
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

func (c *CalculateAssignment) Run(ctx ken.Context) (err error) {
	refreshRoleId, err := helpers.GetRoleId(ctx, config.Values.Roles["CanRefresh"])
	if refreshRoleId == "" || err != nil {
		log.Errorw("error finding role ID", "roleTag", "CanRefresh", "configuredRole", config.Values.Roles["CanRefresh"], "error", err)
		return
	}

	if helpers.CheckRoleMembership(ctx, refreshRoleId) {
		_ = ctx.Respond(helpers.GetDefaultResponse("Calculating kit assignments.", false, ctx))
		calculateAssignments()
		err = ctx.FollowUpMessage("Calculation complete").Send().Error
	} else {
		err = ctx.Respond(helpers.GetDefaultResponse("This command is not available to this user in this context.", true, ctx))
	}

	return
}
