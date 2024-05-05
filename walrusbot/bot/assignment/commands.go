package assignment

import (
	"walrusbot/utility/helpers"

	"github.com/FedorLap2006/disgolf"
	"github.com/bwmarrin/discordgo"
)

var myAssignment = func(ctx *disgolf.Ctx) {
	helpers.HandlerWrapper("", ctx, requestAssignment)
}

var myAss = func(ctx *disgolf.Ctx) {
	helpers.HandlerWrapper("", ctx, requestAssignment)
}

var refreshAssignments = func(ctx *disgolf.Ctx) {
	helpers.HandlerWrapper("", ctx, assignmentRefresh)
}

var calculateAssignment = func(ctx *disgolf.Ctx) {
	helpers.HandlerWrapper("", ctx, assignmentCalculate)
}

var viewAssignments = func(ctx *disgolf.Ctx) {
	helpers.HandlerWrapper("", ctx, assignmentView)
}

var MyAssignment = &disgolf.Command{
	Name:        "myassignment",
	Description: "Get your species war assignments",
	Type:        discordgo.ChatApplicationCommand,
	Handler:     disgolf.HandlerFunc(myAssignment),
}

var MyAss = &disgolf.Command{
	Name:        "myass",
	Description: "Get your species war assignments",
	Type:        discordgo.ChatApplicationCommand,
	Handler:     disgolf.HandlerFunc(myAss),
}

var RefreshAssignment = &disgolf.Command{
	Name:        "refreshassignments",
	Description: "Refresh the species war assignment cache from the data source",
	Type:        discordgo.ChatApplicationCommand,
	Handler:     disgolf.HandlerFunc(refreshAssignments),
}

var CalculateAssignment = &disgolf.Command{
	Name:        "calculateassignments",
	Description: "Calculate the species war assignments based on snail stats in the database",
	Type:        discordgo.ChatApplicationCommand,
	Handler:     disgolf.HandlerFunc(calculateAssignment),
}

var ViewAssignments = &disgolf.Command{
	Name:        "viewassignments",
	Description: "List the assignments for a club",
	Type:        discordgo.ChatApplicationCommand,
	Options: []*discordgo.ApplicationCommandOption{{
		Type:        discordgo.ApplicationCommandOptionString,
		Name:        "club",
		Description: "Name of the club to view",
		Required:    false,
	}},
	Handler: disgolf.HandlerFunc(viewAssignments),
}
