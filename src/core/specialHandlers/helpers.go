package specialHandlers

import (
	"github.com/AnimeKaizoku/Punk05Robot/src/core/wotoConfig"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

func SudoOnlyCommand(c string, r handlers.Response) handlers.Command {
	return handlers.NewCommand(c, func(b *gotgbot.Bot, ctx *ext.Context) error {
		if !wotoConfig.IsSudoUser(ctx.EffectiveUser.Id) {
			return ext.ContinueGroups
		}

		return r(b, ctx)
	})
}
