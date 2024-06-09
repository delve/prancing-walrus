package beforeMiddleware

import (
	"errors"
	"fmt"
	"slices"
	"walrusbot/utility/helpers"
	"walrusbot/utility/log"

	"github.com/zekrotja/ken"
)

type CommandAccessControl struct{}

type RequiresOneRole interface {
	AllowedRoles(*ken.Ctx) []string
}

var (
	_ ken.MiddlewareBefore = (*CommandAccessControl)(nil)
)

func (c *CommandAccessControl) Before(ctx *ken.Ctx) (next bool, err error) {
	cmd, ok := ctx.Command.(RequiresOneRole)
	if ok {
		return requiresRole(ctx, cmd)
	}

	// Fallthrough, no access restrictions
	return true, nil
}

func requiresRole(ctx *ken.Ctx, cmd RequiresOneRole) (next bool, err error) {
	allowedRoles := cmd.AllowedRoles(ctx)
	if len(allowedRoles) < 1 {
		ctx.Respond(helpers.GetDefaultResponse("Sorry, I seem to have a problem. Paging <pingCaretakerRole> to review the log", true, ctx))
		return false, errors.New("allowed role list is empty")
	}

	allowedRoleIds := make([]string, len(allowedRoles))
	for _, roleName := range allowedRoles {
		roleId, err := helpers.GetRoleId(ctx, roleName)
		if roleId == "" || err != nil {
			log.Errorw("error finding role ID in requiresRole()", "roleName", roleName, "error", err)
			ctx.Respond(helpers.GetDefaultResponse("Sorry, I seem to have a problem. Paging <pingCaretakerRole> to review the log", true, ctx))
			return false, errors.New("error finding role ID")
		}
		allowedRoleIds = append(allowedRoleIds, roleId)
	}

	for _, v := range ctx.GetEvent().Member.Roles {
		if slices.Contains(allowedRoleIds, v) {
			return true, nil
		}
	}

	// found no matching roles, reject
	msg := fmt.Sprintf("Sorry, you need one of these roles to do that. %v", allowedRoles)
	ctx.Respond(helpers.GetDefaultResponse(msg, true, ctx))
	return false, nil
}
