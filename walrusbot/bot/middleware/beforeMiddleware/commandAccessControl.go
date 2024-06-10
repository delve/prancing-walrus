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

// An AccessControlledSubcommand has a mixed set of subcommands where only some are access controlled.
//
//	IsControlledSubcommand must return true is the subcommand in ctx is controlled. See helpers.GetSubcommandFromCtx(ctx).
//	Uncontrolled subcommands short circuit out of the access control middleware
type AccessControlledSubcommand interface {
	IsControlledSubcommand(*ken.Ctx) bool
}

// A RequiresAllRoles commande passes the access control middleware only if all required role are found on the user in ctx.
//
//	This interface is processed first and additively to RequiresAllRoles. So you can have a hard requirement
//	for a set of roles plus one more role from a semi-optional set.
type RequiresAllRoles interface {
	RequiredRoles(*ken.Ctx) []string
}

// A RequiresOneRole command passes the access control middleware if any one allowed role is found on the user in ctx.
//
//	This interface is processed second and additively to RequiresAllRoles. So you can have a hard requirement
//	for a set of roles plus one more role from a semi-optional set.
type RequiresOneRole interface {
	AllowedRoles(*ken.Ctx) []string
}

var (
	_ ken.MiddlewareBefore = (*CommandAccessControl)(nil)
)

func (c *CommandAccessControl) Before(ctx *ken.Ctx) (next bool, err error) {
	if cmd, ok := ctx.Command.(AccessControlledSubcommand); ok {
		if !cmd.IsControlledSubcommand(ctx) {
			// if this subcommand is NOT controlled, short circuit and move on
			return true, nil
		}
	}

	if cmd, ok := ctx.Command.(RequiresAllRoles); ok {
		if next, err = hasAllRoles(ctx, cmd); !next || err != nil {
			return
		}
	}

	if cmd, ok := ctx.Command.(RequiresOneRole); ok {
		if next, err = hasSingleRole(ctx, cmd); !next || err != nil {
			return
		}
	}

	// Fallthrough, access restrictionsare satisfied or not required
	return true, nil
}

func hasSingleRole(ctx *ken.Ctx, cmd RequiresOneRole) (next bool, err error) {
	allowedRoles := cmd.AllowedRoles(ctx)
	if len(allowedRoles) < 1 {
		ctx.Respond(helpers.GetDefaultResponse("Sorry, I seem to have a problem. Paging <pingCaretakerRole> to review the log", true, ctx))
		return false, errors.New("allowed role list is empty")
	}

	allowedRoleIds := make([]string, len(allowedRoles))
	for _, roleName := range allowedRoles {
		roleId, err := helpers.GetRoleId(ctx, roleName)
		if roleId == "" || err != nil {
			log.Errorw("error finding role ID in hasSingleRole()", "roleName", roleName, "error", err)
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

func hasAllRoles(ctx *ken.Ctx, cmd RequiresAllRoles) (next bool, err error) {
	requiredRoles := cmd.RequiredRoles(ctx)
	if len(requiredRoles) < 1 {
		ctx.Respond(helpers.GetDefaultResponse("Sorry, I seem to have a problem. Paging <pingCaretakerRole> to review the log", true, ctx))
		return false, errors.New("required role list is empty")
	}
	userRoles := ctx.GetEvent().Member.Roles
	for _, roleName := range requiredRoles {
		roleId, err := helpers.GetRoleId(ctx, roleName)
		if roleId == "" || err != nil {
			log.Errorw("error finding role ID in hasAllRoles()", "roleName", roleName, "error", err)
			ctx.Respond(helpers.GetDefaultResponse("Sorry, I seem to have a problem. Paging <pingCaretakerRole> to review the log", true, ctx))
			return false, errors.New("error finding role ID")
		}
		if !slices.Contains(userRoles, roleId) {
			// missing this role.
			msg := fmt.Sprintf("Sorry, you need all of these roles to do that. %v", requiredRoles)
			ctx.Respond(helpers.GetDefaultResponse(msg, true, ctx))
			return false, nil
		}
	}

	// found all matching roles, proceed
	return true, nil
}
