package plugins

import (
	"github.com/AnimeKaizoku/kaizokuReposterRobot/src/plugins/repostPlugin"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func LoadAllHandlers(d *ext.Dispatcher, triggers []rune) {
	repostPlugin.LoadAllHandlers(d, triggers)
}
