package beforeMiddleware

import (
	"walrusbot/utility/helpers"
	"walrusbot/utility/log"

	"github.com/zekrotja/ken"
)

type CommandLogging struct{}

var (
	_ ken.MiddlewareBefore = (*CommandLogging)(nil)
)

func (c *CommandLogging) Before(ctx *ken.Ctx) (next bool, err error) {
	chanName := ""
	thisChan, err := ctx.Channel()
	if err != nil {
		log.Errorw("Error retrieving channel from context in command logging", "ctx", ctx, "channelId", ctx.GetEvent().ChannelID)
		chanName = ""
	} else {
		chanName = thisChan.Name
	}
	command := ctx.GetCommand().Name()

	log.Infow("Command logging", "command", command, "subcommand", helpers.GetSubcommandFromCtx(ctx), "channel", chanName, "user", ctx.User().Username)

	return true, nil
}
