package helpers

import (
	"errors"
	"walrusbot/sheetDAO"
	"walrusbot/utility/log"

	"github.com/FedorLap2006/disgolf"
)

func GetRoleId(ctx *disgolf.Ctx, roleName string) (string, error) {
	for _, guildRole := range ctx.Session.State.Guilds[0].Roles {
		if guildRole.Name == roleName {
			return guildRole.ID, nil
		}
	}
	return "", errors.New("role not found")
}

func CheckroleMembership(ctx *disgolf.Ctx, roleID string) bool {
	for _, v := range ctx.Interaction.Member.Roles {
		if v == roleID {
			return true
		}
	}
	return false
}

func IsDiscoUserInClub(username, club string) bool{
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