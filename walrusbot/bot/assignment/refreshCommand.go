package assignment

import (
	"walrusbot/utility/config"
	"walrusbot/utility/helpers"
	"walrusbot/utility/log"

	"github.com/bwmarrin/discordgo"
	"github.com/zekrotja/ken"
)

type RefreshAssignment struct{}

var (
	_ ken.SlashCommand = (*RefreshAssignment)(nil)
	_ ken.DmCapable    = (*RefreshAssignment)(nil)
)

func (c *RefreshAssignment) Name() string {
	return "refreshassignments"
}

func (c *RefreshAssignment) Description() string {
	return "Refresh the species war assignment cache from the data source"
}

func (c *RefreshAssignment) Version() string {
	return "1.0.0"
}

func (c *RefreshAssignment) Type() discordgo.ApplicationCommandType {
	return discordgo.ChatApplicationCommand
}

func (c *RefreshAssignment) IsDmCapable() bool {
	return true
}

func (c *RefreshAssignment) Options() []*discordgo.ApplicationCommandOption {
	return []*discordgo.ApplicationCommandOption{}
}

func (c *RefreshAssignment) Run(ctx ken.Context) (err error) {
	refreshRoleId, err := helpers.GetRoleId(ctx, config.Values.Roles["CanRefresh"])
	if refreshRoleId == "" || err != nil {
		log.Errorw("error finding role ID", "roleTag", "CanRefresh", "configuredRole", config.Values.Roles["CanRefresh"], "error", err)
		return
	}

	if helpers.CheckRoleMembership(ctx, refreshRoleId) {
		_ = ctx.Respond(helpers.GetDefaultResponse("Refreshing assignment cache.", false, ctx))
		CacheAssignments()
		err = ctx.FollowUpMessage("Assignments refreshed from spreadsheet.").Send().Error
	} else {
		_ = ctx.Respond(helpers.GetDefaultResponse("This command is not available to this user in this context.", true, ctx))
	}

	return
}
