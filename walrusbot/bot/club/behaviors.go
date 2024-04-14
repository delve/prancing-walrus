package club

import (
	"fmt"
	"slices"
	"strings"
	"walrusbot/sheetDAO"
	"walrusbot/utility/helpers"
	"walrusbot/utility/log"

	"github.com/FedorLap2006/disgolf"
	"github.com/bwmarrin/discordgo"
	"github.com/delve/sheetdb"
)

func memberList(ctx *disgolf.Ctx) {

	clubs, err := sheetDAO.GetPlayerClubMemberships(ctx.Interaction.Member.User.Username)
	if err != nil {
		log.Errorw("Error retrieving player club memberships in /club members", "err", err)
		_ = ctx.Respond(helpers.GetDefaultResponse("Sorry, there was an issue retrieving your club info. Paging <pingCaretakerRole> to review the log", true, ctx))
		return
	}
	if len(clubs) < 1 {
		_ = ctx.Respond(helpers.GetDefaultResponse("Sorry, looks like you don't have any snails in a club. Join one, they're fun! Or maybe you need to ping your club officers to induct you on Discord.", true, ctx))
		return
	}

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Hi %s! You have snails in %d clubs. Here's a list of all of them in no particular order.\n", ctx.Interaction.Member.User.Username, len(clubs)))
	for _, club := range clubs {
		sb.WriteString(fmt.Sprintf("\n--- __%s__ ---\n", club.Name))

		clubFilter := func(snail *sheetDAO.Snail) bool { return snail.Club == club.ClubID }
		members, _ := sheetDAO.GetAllSnails(sheetDAO.SnailFilter(clubFilter))
		for _, member := range members {
			sb.WriteString(fmt.Sprintf("%s\n", member.SnailName))
		}
	}
	_ = ctx.Respond(helpers.GetDefaultResponse(sb.String(), true, ctx))
}

func snailInduct(ctx *disgolf.Ctx) {
	var sb strings.Builder
	var targetClub *sheetDAO.Club

	type officerRole struct {
		Name   string
		role   *discordgo.Role
		record *sheetDAO.Club
	}
	officerRoles := []officerRole{}
	for _, rle := range helpers.GetOfficerRoleMemberships(ctx) {
		// get the club from the DB using the abbreviation at the front of the role name
		clubRec, err := sheetDAO.GetClubByAbrv(strings.Split(rle.Name, " ")[0])
		// just ignoring errors, skipping those :(
		if err != nil {
			continue
		}
		officerRoles = append(officerRoles, officerRole{Name: clubRec.Name, role: rle, record: clubRec})
	}

	playerClubList := ""
	if len(officerRoles) > 0 {
		sb.WriteString(officerRoles[0].Name)
		for idx, cb := range officerRoles {
			if idx > 0 {
				sb.WriteString(fmt.Sprintf(", %s", cb.Name))
			}
		}
		playerClubList = sb.String()
	}
	sb.Reset()

	// if club param given
	clubParam, ok := ctx.Options["club"]
	if ok {
		// validate club name
		club, err := sheetDAO.GetClubByName(clubParam.StringValue())
		if _, isNotFound := err.(*sheetdb.NotFoundError); isNotFound {
			msg := fmt.Sprintf("I've never heard of the club %s.", ctx.Options["club"].StringValue())
			if len(playerClubList) > 0 {
				msg = fmt.Sprintf("%s Did you mean one of these?\n%s", msg, playerClubList)
			}
			_ = ctx.Respond(helpers.GetDefaultResponse(msg, true, ctx))
			return
		}
		if err != nil {
			_ = ctx.Respond(helpers.GetDefaultResponse("Sorry, there was a problem getting your club info. Paging <pingCaretakerRole> to review the log", true, ctx))
			log.Errorw("Error retrieving player clubs in /club induct", "err", err)
			return
		}

		targetClub = club
	} else {
		// if club memberships > 1 fail due to ambiguity, else assume club
		if len(officerRoles) > 1 {
			_ = ctx.Respond(helpers.GetDefaultResponse(fmt.Sprintf("Sorry, you're a Discord officer of more than 1 club. You're going to have to tell me which of these you want.\n%s", playerClubList), true, ctx))
			return
		}
		targetClub = officerRoles[0].record
	}

	// check officer role membership
	if !helpers.IsClubOfficer(ctx, targetClub.Name) {
		_ = ctx.Respond(helpers.GetDefaultResponse(fmt.Sprintf("Looks like you're not an officer of %s. You Need to be a member of the club's Officer discord role to use this command.", targetClub.Name), true, ctx))
		return
	}

	// check target snail club membership, if already in a club abort
	snailFilter := func(snail *sheetDAO.Snail) bool { return snail.SnailName == ctx.Options["snail"].StringValue() }
	snail, err := sheetDAO.GetAllSnails(sheetDAO.SnailFilter(snailFilter))
	if _, isNotFound := err.(*sheetdb.NotFoundError); isNotFound {
		_ = ctx.Respond(helpers.GetDefaultResponse(fmt.Sprintf("I don't know any snail named %s.", ctx.Options["snail"].StringValue()), true, ctx))
		return
	}
	if err != nil || len(snail) != 1 { // if there's an error or we don't get exactly 1 snail, abort
		_ = ctx.Respond(helpers.GetDefaultResponse("Sorry, there was a problem. Paging <pingCaretakerRole> to review the log", true, ctx))
		log.Errorw("Error retrieving target snail in /club induct", "err", err)
		return
	}

	if snail[0].Club != 1 { // they're in a club already, they need to leave the current one first.
		_ = ctx.Respond(helpers.GetDefaultResponse("Sorry, that snail is in a club already. They need to be kicked from their current club before you can induct them", true, ctx))
		return
	}

	// update target snail
	snail[0].Club = targetClub.ClubID
	err = snail[0].UpdateThisSnail()
	if err != nil {
		_ = ctx.Respond(helpers.GetDefaultResponse("Sorry, there was a problem. Paging <pingCaretakerRole> to review the log", true, ctx))
		log.Errorw("Error updating target snail in /club induct", "err", err)
		return
	}
	_ = ctx.Respond(helpers.GetDefaultResponse(fmt.Sprintf("%s has been inducted to %s! Send them my congratulations and a big hug <:lucylove:1181734603401220177>", ctx.Options["snail"].StringValue(), targetClub.Name), true, ctx))
}

func snailKick(ctx *disgolf.Ctx) {
	// check target snail club membership
	snailFilter := func(snail *sheetDAO.Snail) bool { return snail.SnailName == ctx.Options["snail"].StringValue() }
	snails, err := sheetDAO.GetAllSnails(sheetDAO.SnailFilter(snailFilter))
	if _, isNotFound := err.(*sheetdb.NotFoundError); isNotFound {
		_ = ctx.Respond(helpers.GetDefaultResponse(fmt.Sprintf("I don't know any snail named %s.", ctx.Options["snail"].StringValue()), true, ctx))
		return
	}
	if err != nil || len(snails) != 1 { // if there's an error or we don't get exactly 1 snail, abort
		_ = ctx.Respond(helpers.GetDefaultResponse("Sorry, there was a problem. Paging <pingCaretakerRole> to review the log", true, ctx))
		log.Errorw("Error retrieving target snail in /club kick", "err", err)
		return
	}
	snail := snails[0]

	if snail.Club == 1 { // they're in a club already, they need to leave the current one first.
		_ = ctx.Respond(helpers.GetDefaultResponse("Sorry, that snail isn't in a club. Don't kick a snail while they're down.", true, ctx))
		return
	}

	// verify user's officer permissions
	officerships := []int{}
	for _, role := range helpers.GetOfficerRoleMemberships(ctx) {
		// get the club from the DB using the abbreviation at the front of the role name
		clubRec, err := sheetDAO.GetClubByAbrv(strings.Split(role.Name, " ")[0])
		// just ignoring errors, skipping those :(
		if err != nil {
			continue
		}
		officerships = append(officerships, clubRec.ClubID)
	}
	if !slices.Contains(officerships, snail.Club) {
		_ = ctx.Respond(helpers.GetDefaultResponse("You're not an officer for that club! Mind your own club's business.", true, ctx))
		return
	}

	// update target snail
	snail.Club = 1
	err = snails[0].UpdateThisSnail()
	if err != nil {
		_ = ctx.Respond(helpers.GetDefaultResponse("Sorry, there was a problem. Paging <pingCaretakerRole> to review the log", true, ctx))
		log.Errorw("Error updating target snail in /club induct", "err", err)
		return
	}
	_ = ctx.Respond(helpers.GetDefaultResponse(fmt.Sprintf("%s has been kicked from your club. I hope everyone is still on good terms.\nRemember, I'm just a walrus. I can't update the game so you'll have to log into Snails to remove them properly.", ctx.Options["snail"].StringValue()), true, ctx))
}
