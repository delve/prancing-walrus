package snail

import (
	"fmt"
	"regexp"
	"strings"
	"time"
	"walrusbot/sheetDAO"
	"walrusbot/utility/helpers"
	"walrusbot/utility/log"

	"github.com/FedorLap2006/disgolf"
	"github.com/delve/sheetdb"
)

func snailAdd(ctx *disgolf.Ctx) {
	if match, err := regexp.MatchString("^[0-9a-zA-Z_-]+$", ctx.Options["name"].StringValue()); !match || err != nil {
		_ = ctx.Respond(helpers.GetDefaultResponse(fmt.Sprintf("%s doesn't look like a valid name. Just what stunt are you trying to pull here?\nOnly allowing letters, numbers, _, and -. Because I couldn't find a list of characters the game consideres valid.", ctx.Options["name"].StringValue()), true, ctx))
		return
	}
	player, err := sheetDAO.GetPlayerByDiscoId(ctx.Interaction.Member.User.Username)
	if _, isErr := err.(*sheetdb.NotFoundError); isErr {
		// create a new player
		player, err = sheetDAO.AddPlayer(ctx.Interaction.Member.User.Username)
		if err != nil {
			_ = ctx.Respond(helpers.GetDefaultResponse("Sorry, there was a problem adding you to the player table. Paging <pingCaretakerRole> to review the log", true, ctx))
			log.Errorw("Error retrieving player in /snail add", "err", err)
			return
		}
	}
	snail, err := getSnail(ctx, ctx.Options["name"].StringValue())
	if err != nil { // logged and responded in function
		return
	}
	if snail != nil {
		_ = ctx.Respond(helpers.GetDefaultResponse(fmt.Sprintf("Sorry, you already told me about that snail. Try `/snail show %s` to check up on them.\nIf this is a new snail in a different server then you can give me a nickname for it.", ctx.Options["name"].StringValue()), true, ctx))
		return
	}

	snail, err = player.AddSnail(int(time.Now().Unix()), ctx.Options["name"].StringValue(), 0, "", 0, "", 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0)
	if err != nil {
		_ = ctx.Respond(helpers.GetDefaultResponse("Sorry, there was a problem adding your snail. Paging <pingCaretakerRole> to review the log", true, ctx))
		log.Errorw("Error adding snail in /snail add", "err", err)
		return
	}
	_ = ctx.Respond(helpers.GetDefaultResponse(fmt.Sprintf("Thanks for telling me about %s! Please use `/snail update %s` to tell me more about them.", snail.SnailName, snail.SnailName), true, ctx))

}

func snailList(ctx *disgolf.Ctx) {
	snails, err := getSnails(ctx)
	if _, isErr := err.(alreadyResponded); isErr { // response has already been sent
		return
	}
	if snails == nil {
		_ = ctx.Respond(helpers.GetDefaultResponse("Sorry, I don't know any of your snails. Maybe you should `/snail add` one.", true, ctx))
	}

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Hi %s! These are the snails I know about.\n", ctx.Interaction.Member.User.Username))
	for _, snail := range snails {
		sb.WriteString(fmt.Sprintf("%s on %s %d\n", snail.SnailName, snail.Server, snail.ServerNum))
	}
	_ = ctx.Respond(helpers.GetDefaultResponse(sb.String(), true, ctx))
}

func snailShow(ctx *disgolf.Ctx) {
	snail, err := getSnail(ctx, ctx.Options["name"].StringValue())
	if err != nil { // logged and responded in function
		return
	}
	if snail == nil {
		_ = ctx.Respond(helpers.GetDefaultResponse(fmt.Sprintf("Sorry, looks like you haven't told me about %[1]s. Maybe you could `/snail add name:%[1]s` them.", ctx.Options["name"].StringValue()), true, ctx))
	}
	_ = ctx.Respond(helpers.GetDefaultResponse(formatSnailStats(snail), true, ctx))
}

func snailUpdate(ctx *disgolf.Ctx) {
	player, err := sheetDAO.GetPlayerByDiscoId(ctx.Interaction.Member.User.Username)
	if _, isNotFound := err.(*sheetdb.NotFoundError); isNotFound { // Counld't find the player
		msg := fmt.Sprintf("Sorry, looks like we haven't met yet. Please use `/snail add name:%s` first so we can get to know each other.", ctx.Options["name"].StringValue())
		_ = ctx.Respond(helpers.GetDefaultResponse(msg, true, ctx))
		return
	}
	if err != nil {
		thisChan, _ := ctx.Channel(ctx.Interaction.ChannelID)
		log.Errorw("In Handler", "command", "snail", "subcommand", "update", "channel", thisChan.Name, "user", ctx.Interaction.Member.User.Username, "options", ctx.Options, "error", err)
		_ = ctx.Respond(helpers.GetDefaultResponse("Sorry, there was a problem updating your snail. Paging <pingCaretakerRole> to review the log", true, ctx))
		return
	}

	snails, _ := player.GetSnails()
	found := false
	for _, snail := range snails {
		if snail.SnailName == ctx.Options["name"].StringValue() {
			found = true
			break
		}
	}
	if !found {
		_ = ctx.Respond(helpers.GetDefaultResponse("What are you tryna pull? That's not your snail. Go mess with someone that doesn't have meter long tusks and weigh 1000 kilograms. Before I get grumpy.", true, ctx))
		return
	}
	err = updateThisSnail(ctx)
	if _, isErr := err.(alreadyResponded); isErr { // response has already been sent
		return
	}
	if err != nil {
		_ = ctx.Respond(helpers.GetDefaultResponse("Sorry, there was a problem updating your snail. Paging <pingCaretakerRole> to review the log", true, ctx))
		// error already logged in updateSnail()
		return
	}
	_ = ctx.Respond(helpers.GetDefaultResponse(fmt.Sprintf("Thanks! I've updated what I know about %s. You can use `/snail show %s` to confirm it.", ctx.Options["name"].StringValue(), ctx.Options["name"].StringValue()), true, ctx))

}

func getSnails(ctx *disgolf.Ctx) ([]*sheetDAO.Snail, error) {
	var responded alreadyResponded = ""
	// get list of snails from DB by disco username ctx.Interaction.Member.User.Username
	player, err := sheetDAO.GetPlayerByDiscoId(ctx.Interaction.Member.User.Username)
	if err != nil {
		_ = ctx.Respond(helpers.GetDefaultResponse("Sorry, there was a problem looking up your player data. Paging <pingCaretakerRole> to review the log", true, ctx))
		log.Errorw("Error retrieving player in getSnails()", "err", err)
		return nil, responded
	}
	snails, err := player.GetSnails()
	if err != nil {
		_ = ctx.Respond(helpers.GetDefaultResponse("Sorry, there was a problem looking up your snail data. Paging <pingCaretakerRole> to review the log", true, ctx))
		log.Errorw("Error retrieving snails in getSnails()", "err", err)
		return nil, responded
	}
	if len(snails) == 0 {
		return nil, nil
	}

	return snails, nil
}

func getSnail(ctx *disgolf.Ctx, name string) (*sheetDAO.Snail, error) {
	snails, err := getSnails(ctx)
	if err != nil {
		return nil, err
	}
	for _, snail := range snails {
		if strings.EqualFold(snail.SnailName, name) {
			return snail, nil
		}
	}
	return nil, nil
}

func formatSnailStats(snail *sheetDAO.Snail) string {
	var club *sheetDAO.Club
	if snail.Club > 0 {
		club, _ = sheetDAO.GetClub(snail.Club)
	} else {
		club, _ = sheetDAO.GetClub(1)
	}
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("This is everything I know about %s.\n", snail.SnailName))
	sb.WriteString(fmt.Sprintf("Server: %s %d\tClub: %s\n", snail.Server, snail.ServerNum, club.Name))
	sb.WriteString(fmt.Sprintf("Leadership: %d\tHoarded SW Essences: %d\n", snail.Leadership, snail.SpeciesWarEssences))
	sb.WriteString(fmt.Sprintf("Total Power: %d\n", snail.TotalPower))
	sb.WriteString(fmt.Sprintf("__AFFCT__\nArt \t%d\tFaith\t%d\nFame\t%d\tCiv\t%d\nTech \t%d\n", snail.Art, snail.Fth, snail.Fame, snail.Civ, snail.Tech))
	sb.WriteString(fmt.Sprintf("__HARD__\nHP \t%d\tAtk\t%d\nRush\t%d\tDef\t%d\n", snail.Hp, snail.Atk, snail.Rush, snail.Def))
	// custom emoji in the Snailverse server
	sb.WriteString("__Form Tiers__\n")
	sb.WriteString(fmt.Sprintf("<:zombie:1211480757382418443>\t%d\t<:demon:1211480750457888798>\t%d\n", snail.ZombieForm, snail.DemonForm))
	sb.WriteString(fmt.Sprintf("<:angel:1211480749195264010>\t%d\t<:mutant:1211480755935514674>\t%d\n", snail.AngelForm, snail.MutantForm))
	sb.WriteString(fmt.Sprintf("<:mecha:1211480754228305990>\t%d\t<:dragon:1211480752311509042>\t%d", snail.MechaForm, snail.DragonForm))

	return sb.String()
}

func updateThisSnail(ctx *disgolf.Ctx) error {
	var responded alreadyResponded = ""
	snail, err := getSnail(ctx, ctx.Options["name"].StringValue())
	if err != nil {
		return err
	}
	if snail == nil {
		_ = ctx.Respond(helpers.GetDefaultResponse(fmt.Sprintf("Sorry, looks like you haven't told me about %[1]s. Maybe you could `/snail add name:%[1]s` them.", ctx.Options["name"].StringValue()), true, ctx))
	}
	responses := []string{}
	stat := ""
	for key, option := range ctx.Options {
		switch key {
		case "name":
			// do nothing with name
		case "leadership":
			if int(option.IntValue()) < 0 {
				responses = append(responses, fmt.Sprintf("%d doesn't look right. Are you sure that's your Leadership?", option.IntValue()))
				break
			}
			snail.Leadership = int(option.IntValue())
		case "sw_essences":
			stat = "SW essences"
			if int(option.IntValue()) < 0 {
				responses = append(responses, fmt.Sprintf("%d? Are you sure that's how many %s you have?", option.IntValue(), stat))
				break
			}
			snail.SpeciesWarEssences = int(option.IntValue())
		case "totalpower":
			stat = "total power"
			value, err := helpers.DebreviateNumber(option.StringValue())
			if err != nil {
				responses = append(responses, fmt.Sprintf("Your %s (%s) doesn't look like a valid number: %s", stat, option.StringValue(), err))
				break
			}
			if value < 0 {
				responses = append(responses, fmt.Sprintf("%d? Are you sure that's your %s?", int(value), stat))
				break
			}
			// discard any remaining fraction, there shouldn't be any in this context anyway
			snail.TotalPower = int(value)
		case "art":
			stat = "art"
			value, err := helpers.DebreviateNumber(option.StringValue())
			if err != nil {
				responses = append(responses, fmt.Sprintf("Your %s (%s) doesn't look like a valid number: %s", stat, option.StringValue(), err))
				break
			}
			if value < 0 {
				responses = append(responses, fmt.Sprintf("%d? Are you sure that's your %s?", int(value), stat))
				break
			}
			// discard any remaining fraction, there shouldn't be any in this context anyway
			snail.Art = int(value)
		case "fth":
			stat = "faith"
			value, err := helpers.DebreviateNumber(option.StringValue())
			if err != nil {
				responses = append(responses, fmt.Sprintf("Your %s (%s) doesn't look like a valid number: %s", stat, option.StringValue(), err))
				break
			}
			if value < 0 {
				responses = append(responses, fmt.Sprintf("%d? Are you sure that's your %s?", int(value), stat))
				break
			}
			// discard any remaining fraction, there shouldn't be any in this context anyway
			snail.Fth = int(value)
		case "fame":
			stat = "fame"
			value, err := helpers.DebreviateNumber(option.StringValue())
			if err != nil {
				responses = append(responses, fmt.Sprintf("Your %s (%s) doesn't look like a valid number: %s", stat, option.StringValue(), err))
				break
			}
			if value < 0 {
				responses = append(responses, fmt.Sprintf("%d? Are you sure that's your %s?", int(value), stat))
				break
			}
			// discard any remaining fraction, there shouldn't be any in this context anyway
			snail.Fame = int(value)
		case "civ":
			stat = "civ"
			value, err := helpers.DebreviateNumber(option.StringValue())
			if err != nil {
				responses = append(responses, fmt.Sprintf("Your %s (%s) doesn't look like a valid number: %s", stat, option.StringValue(), err))
				break
			}
			if value < 0 {
				responses = append(responses, fmt.Sprintf("%d? Are you sure that's your %s?", int(value), stat))
				break
			}
			// discard any remaining fraction, there shouldn't be any in this context anyway
			snail.Civ = int(value)
		case "tech":
			stat = "tech"
			value, err := helpers.DebreviateNumber(option.StringValue())
			if err != nil {
				responses = append(responses, fmt.Sprintf("Your %s (%s) doesn't look like a valid number: %s", stat, option.StringValue(), err))
				break
			}
			if value < 0 {
				responses = append(responses, fmt.Sprintf("%d? Are you sure that's your %s?", int(value), stat))
				break
			}
			// discard any remaining fraction, there shouldn't be any in this context anyway
			snail.Tech = int(value)
		case "hp":
			stat = "HP"
			value, err := helpers.DebreviateNumber(option.StringValue())
			if err != nil {
				responses = append(responses, fmt.Sprintf("Your %s (%s) doesn't look like a valid number: %s", stat, option.StringValue(), err))
				break
			}
			if value < 0 {
				responses = append(responses, fmt.Sprintf("%d? Are you sure that's your %s?", int(value), stat))
				break
			}
			// discard any remaining fraction, there shouldn't be any in this context anyway
			snail.Hp = int(value)
		case "atk":
			stat = "attack"
			value, err := helpers.DebreviateNumber(option.StringValue())
			if err != nil {
				responses = append(responses, fmt.Sprintf("Your %s (%s) doesn't look like a valid number: %s", stat, option.StringValue(), err))
				break
			}
			if value < 0 {
				responses = append(responses, fmt.Sprintf("%d? Are you sure that's your %s?", int(value), stat))
				break
			}
			// discard any remaining fraction, there shouldn't be any in this context anyway
			snail.Atk = int(value)
		case "rush":
			stat = "rush"
			value, err := helpers.DebreviateNumber(option.StringValue())
			if err != nil {
				responses = append(responses, fmt.Sprintf("Your %s (%s) doesn't look like a valid number: %s", stat, option.StringValue(), err))
				break
			}
			if value < 0 {
				responses = append(responses, fmt.Sprintf("%d? Are you sure that's your %s?", int(value), stat))
				break
			}
			// discard any remaining fraction, there shouldn't be any in this context anyway
			snail.Rush = int(value)
		case "def":
			stat = "defense"
			value, err := helpers.DebreviateNumber(option.StringValue())
			if err != nil {
				responses = append(responses, fmt.Sprintf("Your %s (%s) doesn't look like a valid number: %s", stat, option.StringValue(), err))
				break
			}
			if value < 0 {
				responses = append(responses, fmt.Sprintf("%d? Are you sure that's your %s?", int(value), stat))
				break
			}
			// discard any remaining fraction, there shouldn't be any in this context anyway
			snail.Def = int(value)
		case "zombie":
			stat = "zombie form tier"
			if int(option.IntValue()) < 0 {
				responses = append(responses, fmt.Sprintf("%d? Are you sure that's your %s?", option.IntValue(), stat))
				break
			}
			snail.ZombieForm = int(option.IntValue())
		case "demon":
			stat = "demon form tier"
			if int(option.IntValue()) < 0 {
				responses = append(responses, fmt.Sprintf("%d? Are you sure that's your %s?", option.IntValue(), stat))
				break
			}
			snail.DemonForm = int(option.IntValue())
		case "angel":
			stat = "angel form tier"
			if int(option.IntValue()) < 0 {
				responses = append(responses, fmt.Sprintf("%d? Are you sure that's your %s?", option.IntValue(), stat))
				break
			}
			snail.AngelForm = int(option.IntValue())
		case "mutant":
			stat = "mutant form tier"
			if int(option.IntValue()) < 0 {
				responses = append(responses, fmt.Sprintf("%d? Are you sure that's your %s?", option.IntValue(), stat))
				break
			}
			snail.MutantForm = int(option.IntValue())
		case "mecha":
			stat = "mecha form tier"
			if int(option.IntValue()) < 0 {
				responses = append(responses, fmt.Sprintf("%d? Are you sure that's your %s?", option.IntValue(), stat))
				break
			}
			snail.MechaForm = int(option.IntValue())
		case "dragon":
			stat = "dragon form tier"
			if int(option.IntValue()) < 0 {
				responses = append(responses, fmt.Sprintf("%d? Are you sure that's your %s?", option.IntValue(), stat))
				break
			}
			snail.DragonForm = int(option.IntValue())

		case "newname":
			//lint:ignore SA6000 looping over a map, this only triggers once
			if match, err := regexp.MatchString("^[0-9a-zA-Z_-]+$", option.StringValue()); !match || err != nil {
				responses = append(responses, fmt.Sprintf("%s doesn't look like a valid name. Just what stunt are you trying to pull here?\nI'm allowing letters, numbers, _, and -. Because I couldn't find a list of characters the game considers valid.", option.StringValue()))
				break
			}
			sn, _ := getSnail(ctx, option.StringValue())
			if sn != nil {
				responses = append(responses, fmt.Sprintf("Sorry, you already have a snail named %s. I can't use that name again. You can use `/snail list` to see all your snails.", option.StringValue()))
				break
			}
			snail.SnailName = option.StringValue()
		}
	}
	if len(responses) > 0 {
		_ = ctx.Respond(helpers.GetDefaultResponse(strings.Join(responses, "\n"), true, ctx))
		return responded
	}

	err = snail.UpdateThisSnail()
	if err != nil {
		log.Errorw("error updating snail", "err", err, "snail", snail, "changes", ctx.Options)
		return err
	}
	return nil
}
