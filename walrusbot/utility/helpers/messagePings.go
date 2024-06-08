package helpers

import (
	"regexp"
	"strings"
	"walrusbot/utility/config"
	"walrusbot/utility/log"

	"github.com/zekrotja/ken"
)

func addPings(ctx ken.Context, message string) string {
	ret := message
	r := regexp.MustCompile("<ping([^>]*)>")

	matches := r.FindAllStringSubmatch(message, -1)
	for _, match := range matches {
		roleId, err := GetRoleId(ctx, config.Values.Roles[match[1]])
		if roleId == "" || err != nil {
			log.Errorw("error finding role ID", "string", message, "roleTag", match[1], "configuredRole", config.Values.Roles[match[1]], "error", err)
			continue
		}
		ret = strings.Replace(ret, match[0], "<@&"+roleId+">", 1)
	}
	return ret
}
