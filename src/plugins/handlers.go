package plugins

import (
	"github.com/AnimeKaizoku/Punk05Robot/src/plugins/channelsPlugin"
	"github.com/AnimeKaizoku/Punk05Robot/src/plugins/repostPlugin"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func LoadAllHandlers(d *ext.Dispatcher, triggers []rune) {
	channelsPlugin.LoadAllHandlers(d, triggers)
	repostPlugin.LoadAllHandlers(d, triggers)
}
