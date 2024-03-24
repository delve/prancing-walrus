package helpers

import (
	"strings"
	"walrusbot/utility/log"

	"github.com/FedorLap2006/disgolf"
	"github.com/bwmarrin/discordgo"
)

// TODO: create addCustomEmoji to parse custom emoji easier
/*
<emoji:snailzombie> -> <:zombie:1211480757382418443>
&etc for below
<:demon:1211480750457888798>
<:angel:1211480749195264010>
<:mutant:1211480755935514674>
<:mecha:1211480754228305990>
<:dragon:1211480752311509042>
*/

func GetDefaultResponse(message string, ephemeral bool, ctx *disgolf.Ctx) *discordgo.InteractionResponse {
	msg := addPings(ctx, message)
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
