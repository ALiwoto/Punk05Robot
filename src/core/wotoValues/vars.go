package wotoValues

import (
	"github.com/ALiwoto/mdparser/mdparser"
	"github.com/AnimeKaizoku/ssg/ssg"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/gotgbot/ratelimiter/ratelimiter"
	"gorm.io/gorm"
)

var (
	HelperBot       *gotgbot.Bot
	BotUpdater      *ext.Updater
	RateLimiter     *ratelimiter.Limiter
	DatabaseSession *gorm.DB
)

var (
	HaltJobs  bool
	PauseJobs bool
)

var PendingJobs = ssg.NewSafeMap[uint64, PendingJob]()

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

var DefaultImplementedButtons = map[ButtonsUniqueId]*gotgbot.InlineKeyboardMarkup{
	ButtonsMoreContent: {
		InlineKeyboard: [][]gotgbot.InlineKeyboardButton{
			{
				{
					Text: "🔘More Content",
					Url:  "http://t.me/Kaizoku",
				},
			},
		},
	},
}

var MoreContentsMd = mdparser.GetHyperLink("🔘More Content", "http://t.me/Kaizoku")
