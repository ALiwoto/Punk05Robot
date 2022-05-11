package wotoValues

import (
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
