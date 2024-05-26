package assignment

import (
	"fmt"
	"strconv"
	"strings"
	"walrusbot/sheetDAO"
	"walrusbot/utility/helpers"
	"walrusbot/utility/log"

	"github.com/olekukonko/tablewriter"
	"github.com/zekrotja/ken"
	"golang.org/x/exp/slices"
)

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

func canListAssignments(ctx ken.Context, club string) bool {
	if helpers.IsDiscoUserInClub(ctx.User().Username, club) {
		return true
	}
	return helpers.IsClubOfficer(ctx, club)
}

func getKitAssignments(clubName string) string {
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

	msg := &strings.Builder{}
	msg.WriteString(fmt.Sprintf("__Species War Kit Assignments for %s__\n", clubName))

	data := [][]string{}
	for _, snail := range snails {
		data = append(data, []string{strconv.Itoa(snail.SWKitRank), snail.SWKit, snail.SnailName, strconv.Itoa(snail.Leadership), strconv.Itoa(snail.MinionSimPower)})
		// msg.WriteString(fmt.Sprintf("%d\t%s\t%s (Leadership: %d)\n", snail.SWKitRank, snail.SWKit, snail.SnailName, snail.Leadership))
	}

	table := tablewriter.NewWriter(msg)
	table.SetHeader([]string{"Rank", "Kit", "Snail", "Leadership", "SimPower"})
	// table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.AppendBulk(data)
	table.Render()

	return msg.String()
}
