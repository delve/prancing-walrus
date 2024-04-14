package botcommands

import (
	"walrusbot/bot/assignment"
	"walrusbot/bot/club"
	"walrusbot/bot/snail"

	"github.com/FedorLap2006/disgolf"
)

// TODO: make a helper function called from main that adds the commands. take a pointer to the bot.

var Commands = []*disgolf.Command{
	assignment.MyAssignment,
	assignment.MyAss,
	assignment.RefreshAssignment,
	assignment.CalculateAssignment,
	assignment.ViewAssignments,
	snail.Snail,
	club.Club,
}

func Load(bot *disgolf.Bot) {
	for _, command := range Commands {
		bot.Router.Register(command)
	}
	return
}
