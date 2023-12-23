package helpers

import (
	"errors"

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