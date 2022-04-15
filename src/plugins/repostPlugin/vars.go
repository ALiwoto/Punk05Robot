package repostPlugin

import "github.com/PaulSonOfLars/gotgbot/v2"

var _defaultButtons = &gotgbot.InlineKeyboardMarkup{
	InlineKeyboard: [][]gotgbot.InlineKeyboardButton{
		{
			{
				Text: "🔘More Content",
				Url:  "http://t.me/Kaizoku",
			},
		},
	},
}
