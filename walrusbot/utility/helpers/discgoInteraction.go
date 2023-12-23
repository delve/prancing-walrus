package helpers

import (
	"strings"
	"walrusbot/utility/log"

	"github.com/FedorLap2006/disgolf"
	"github.com/bwmarrin/discordgo"
)

func GetDefaultResponse(message string, ephemeral bool, ctx *disgolf.Ctx) *discordgo.InteractionResponse {
	msg := AddPings(ctx, message)
	// <@& indicates an attempt to ping a role
	if ephemeral && strings.Contains(msg, "<@&") {
		log.Warnw("overriding ephemeral setting due to ping", "interactionMessage", msg)
		ephemeral = false
	}
	if ephemeral {
		return &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: msg,
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		}
	} else {
		return &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: msg,
			},
		}
	}

}
