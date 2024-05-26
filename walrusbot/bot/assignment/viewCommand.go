package assignment

import (
	"walrusbot/sheetDAO"
	"walrusbot/utility/helpers"
	"walrusbot/utility/log"

	"github.com/bwmarrin/discordgo"
	"github.com/zekrotja/ken"
)

type ViewAssignment struct{}

var (
	_ ken.SlashCommand = (*ViewAssignment)(nil)
	_ ken.DmCapable    = (*ViewAssignment)(nil)
)

func (c *ViewAssignment) Name() string {
	return "viewassignments"
}

func (c *ViewAssignment) Description() string {
	return "Calculate the species war assignments based on player entered data"
}

func (c *ViewAssignment) Version() string {
	return "1.0.0"
}

func (c *ViewAssignment) Type() discordgo.ApplicationCommandType {
	return discordgo.ChatApplicationCommand
}

func (c *ViewAssignment) IsDmCapable() bool {
	return true
}

func (c *ViewAssignment) Options() []*discordgo.ApplicationCommandOption {
	return []*discordgo.ApplicationCommandOption{
		{
			Type:        discordgo.ApplicationCommandOptionString,
			Name:        "club",
			Description: "Name of the club to view",
			Required:    false,
		},
	}
}

func (c *ViewAssignment) Run(ctx ken.Context) (err error) {
	clubName := ""
	val, ok := ctx.Options().GetByNameOptional("club")

	if ok {
		clubName = val.StringValue()
	} else {
		player, err := sheetDAO.GetPlayerByDiscoId(ctx.User().Username)
		if err == nil {
			snails, err := player.GetSnails()
			if err == nil {
				cr, err := sheetDAO.GetClub(snails[0].Club)
				if err == nil {
					clubName = cr.Name
				}
			}
		}
	}

	if !canListAssignments(ctx, clubName) || clubName == "none" {
		_ = ctx.Respond(helpers.GetDefaultResponse("That's not your club, sorry, can't help you.", true, ctx))
	}

	msg := getKitAssignments(clubName)

	response := helpers.GetDefaultResponse(msg, true, ctx)
	err = ctx.Respond(response)
	if err != nil {
		log.Errorw("error sending disco response", "error", err, "response", response)
	}

	err = ctx.Respond(&discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: msg,
		},
	})
	return
}
