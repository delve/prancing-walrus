package assignment

import (
	"fmt"
	"strings"
	"walrusbot/sheetDAO"
	"walrusbot/utility/config"
	"walrusbot/utility/helpers"
	"walrusbot/utility/log"

	"github.com/FedorLap2006/disgolf"
	"golang.org/x/exp/slices"
)

func assignmentRefresh(ctx *disgolf.Ctx) {
	refreshRoleId, err := helpers.GetRoleId(ctx, config.Values.Roles["CanRefresh"])
	if refreshRoleId == "" || err != nil {
		log.Errorw("error finding role ID", "roleTag", "CanRefresh", "configuredRole", config.Values.Roles["CanRefresh"], "error", err)
		return
	}

	if helpers.CheckRoleMembership(ctx, refreshRoleId) {
		_ = ctx.Respond(helpers.GetDefaultResponse("Refreshing assignment cache.", false, ctx))
		CacheAssignments()
		// TODO: figure out mutlipart interaction responses
	} else {
		_ = ctx.Respond(helpers.GetDefaultResponse("This command is not available to this user in this context.", true, ctx))
	}

}

func requestAssignment(ctx *disgolf.Ctx) {
	sass := (ctx.Caller.Name == "myass")
	_ = ctx.Respond(helpers.GetDefaultResponse(getAssignmentMessage(ctx.Interaction.Member.User.Username, sass), true, ctx))
}

func assignmentCalculate(ctx *disgolf.Ctx) {
	refreshRoleId, err := helpers.GetRoleId(ctx, config.Values.Roles["CanRefresh"])
	if refreshRoleId == "" || err != nil {
		log.Errorw("error finding role ID", "roleTag", "CanRefresh", "configuredRole", config.Values.Roles["CanRefresh"], "error", err)
		return
	}

	if helpers.CheckRoleMembership(ctx, refreshRoleId) {
		_ = ctx.Respond(helpers.GetDefaultResponse("Calculating kit assignments.", false, ctx))
		calculateAssignments()
		// TODO: figure out mutlipart interaction responses
	} else {
		_ = ctx.Respond(helpers.GetDefaultResponse("This command is not available to this user in this context.", true, ctx))
	}
}

func assignmentView(ctx *disgolf.Ctx) {
	clubName := ""
	val, ok := ctx.Options["club"]

	if ok {
		clubName = val.StringValue()
	} else {
		player, err := sheetDAO.GetPlayerByDiscoId(ctx.Interaction.Member.User.Username)
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
	_ = ctx.Respond(helpers.GetDefaultResponse(msg, true, ctx))
}

func getAssignmentMessage(name string, sass bool) (assignMsg string) {
	tone := 0
	if sass {
		tone = 1
	}
	assignMsg = noAssignmentTxt[tone]
	assign, found := assignments[name]
	if found {
		var sb strings.Builder
		sb.WriteString(fmt.Sprintf("Hi %s! ", assign.gameName))
		if assign.role == "" {
			sb.WriteString(noKitAssignmentTxt[tone])
		} else {
			sb.WriteString(fmt.Sprintf(kitAssignmentTxt[tone], assign.role))
		}

		if assign.gather == "" {
			sb.WriteString(noFossilAssignmentTxt[tone])
		} else {
			sb.WriteString(fmt.Sprintf(fossilAssignmentTxt[tone], assign.gather))
		}

		if assign.canUseClamMagic {
			sb.WriteString(canUseClamMagicTxt[tone])
		} else {
			sb.WriteString(canNotUseClamMagicTxt[tone])
		}
		sb.WriteString(closing[tone])
		assignMsg = sb.String()
	}
	return
}

func canListAssignments(ctx *disgolf.Ctx, club string) bool {
	if helpers.IsDiscoUserInClub(ctx.Interaction.Member.User.Username, club) {
		return true
	}
	return helpers.IsClubOfficer(ctx, club)
}

func getKitAssignments(clubName string) string {
	var msg strings.Builder

	clubRec, err := sheetDAO.GetClubByName(clubName)
	if err != nil {
		log.Warnw("could not find club in getKitAssignments", "club", clubName)
		return "Oops. Had a problem finding your club :confounded:"
	}

	clubFilter := func(snail *sheetDAO.Snail) bool { return snail.Club == clubRec.ClubID }
	leaderSort := func(snails []*sheetDAO.Snail) {
		slices.SortStableFunc(snails,
			func(a, b *sheetDAO.Snail) int { return b.Leadership - a.Leadership })
	}

	snails, err := sheetDAO.GetAllSnails(sheetDAO.SnailFilter(clubFilter), sheetDAO.SnailSort(leaderSort))
	if err != nil {
		log.Warnw("could not get snail list in getKitAssignments", "club", clubName)
		return "Oops. Had a problem retrieving the list :confounded:"
	}
	if len(snails) < 1 {
		log.Warnw("found zero snails in getKitAssignments", "club", clubName)
		return "Oops. Had a problem retrieving the list :confounded:"
	}

	msg.WriteString(fmt.Sprintf("__Species War Kit Assignments for %s__\n", clubName))
	for _, snail := range snails {
		msg.WriteString(fmt.Sprintf("%d\t%s\t%s (Leadership: %d)\n", snail.SWKitRank, snail.SWKit, snail.SnailName, snail.Leadership))
	}

	return msg.String()
}
