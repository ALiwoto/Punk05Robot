package wotoValues

import (
	"github.com/AnimeKaizoku/ssg/ssg"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/gotgbot/ratelimiter/ratelimiter"
)

var (
	HelperBot   *gotgbot.Bot
	BotUpdater  *ext.Updater
	RateLimiter *ratelimiter.Limiter
)

var (
	HaltJobs  bool
	PauseJobs bool
)

var PendingJobs = ssg.NewSafeMap[string, PendingJob]()
