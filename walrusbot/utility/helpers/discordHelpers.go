package helpers

import (
	"errors"
	"regexp"
	"strings"
	"walrusbot/sheetDAO"
	"walrusbot/utility/log"

	"github.com/FedorLap2006/disgolf"
	"github.com/bwmarrin/discordgo"
)

var IntegerOptionZeroValue = 0.0

func GetRoleId(ctx *disgolf.Ctx, roleName string) (string, error) {
	for _, guildRole := range ctx.Session.State.Guilds[0].Roles {
		if guildRole.Name == roleName {
			return guildRole.ID, nil
		}
	}
	return "", errors.New("role not found")
}

func CheckRoleMembership(ctx *disgolf.Ctx, roleID string) bool {
	for _, v := range ctx.Interaction.Member.Roles {
		if v == roleID {
			return true
		}
	}
	return false
}

func getOfficerRoles(ctx *disgolf.Ctx) (roles []*discordgo.Role) {
	roles = []*discordgo.Role{}
	r := regexp.MustCompile(".* Officers$")
	for _, guildRole := range ctx.Session.State.Guilds[0].Roles {
		if r.Match([]byte(guildRole.Name)) {
			roles = append(roles, guildRole)
		}
	}
	return
}

func GetOfficerRoleMemberships(ctx *disgolf.Ctx) (memberships []*discordgo.Role) {
	memberships = []*discordgo.Role{}
	officerRoles := getOfficerRoles(ctx)
	for _, role := range ctx.Interaction.Member.Roles {
		for _, officerRole := range officerRoles {
			if role == officerRole.ID {
				memberships = append(memberships, officerRole)
			}
		}
	}
	return memberships
}

func IsDiscoUserInClub(username, club string) bool {
	player, err := sheetDAO.GetPlayerByDiscoId(username)
	if err != nil {
		log.Warnw("unable to find player by Discord username in IsDiscoUserInClub", "username", username, "club", club)
		return false
	}

	clubRec, err := sheetDAO.GetClubByName(club)
	if err != nil {
		log.Warnw("unable to find club in IsDiscoUserInClub", "username", username, "club", club)
		return false
	}

	snails, err := player.GetSnails(sheetDAO.SnailFilter(func(snail *sheetDAO.Snail) bool { return snail.Club == clubRec.ClubID }))
	if err != nil {
		log.Warnw("unable to find snails for player in IsDiscoUserInClub", "username", username, "club", club)
		return false
	}

	if len(snails) > 0 {
		return true
	}

	return false
}

func IsClubOfficer(ctx *disgolf.Ctx, club string) bool {
	officerRole := strings.ReplaceAll(club, " ", "") + " Officers"
	officerRoleId, err := GetRoleId(ctx, officerRole)
	if err != nil {
		log.Errorw("error finding role ID from context", "role", officerRole, "error", err)
		return false
	}
	if CheckRoleMembership(ctx, officerRoleId) {
		return true
	}

	return false
}
