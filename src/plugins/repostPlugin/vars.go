package repostPlugin

import (
	"github.com/AnimeKaizoku/ssg/ssg"
	"github.com/PaulSonOfLars/gotgbot/v2"
)

var MoreContentButtons = &gotgbot.InlineKeyboardMarkup{
	InlineKeyboard: [][]gotgbot.InlineKeyboardButton{
		{
			{
				Text: "🔘More Content",
				Url:  "http://t.me/Kaizoku",
			},
		},
	},
}

var (
	repeatCheckerMap      = _getRepeatCheckerMap()
	mediaGroupMessagesMap = _getMediaGroupMessagesMap()
)

var (
	jobsKeyGenerator = ssg.NewNumIdGenerator[uint64]()
)
