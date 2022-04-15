package wotoValues

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

type PendingJob struct {
	Bot     *gotgbot.Bot
	Ctx     *ext.Context
	Handler func(job *PendingJob) error
}
