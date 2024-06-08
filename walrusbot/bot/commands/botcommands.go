package botcommands

import (
	"walrusbot/bot/assignment"
	"walrusbot/bot/club"
	"walrusbot/bot/snail"
	"walrusbot/commandTemplates/commands"
	"walrusbot/commandTemplates/subcommands"

	"github.com/zekrotja/ken"
)

var Commands = []ken.Command{
	new(commands.TestCommand),
	new(subcommands.SubsCommand),
	new(club.Club),
	new(assignment.ViewAssignment),
	new(assignment.MyAss),
	new(assignment.MyAssignment),
	new(assignment.RefreshAssignment),
	new(assignment.CalculateAssignment),
	new(snail.Snail),
}
