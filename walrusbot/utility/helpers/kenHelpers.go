package helpers

import "github.com/zekrotja/ken"

func GetSubcommandFromCtx(ctx *ken.Ctx) string {
	if len(ctx.GetEvent().Interaction.ApplicationCommandData().Options) > 0 {
		return ctx.GetEvent().Interaction.ApplicationCommandData().Options[0].Name
	}

	return ""
}
