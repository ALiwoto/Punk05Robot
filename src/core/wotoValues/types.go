package wotoValues

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

type PendingJob struct {
	Bot            *gotgbot.Bot
	Ctx            *ext.Context
	Handler        func(job *PendingJob) error
	RegisteredTime time.Time
	TimeDistance   time.Duration
}

type ChannelSettings struct {
	// ChannelId is the channel id.
	ChannelId int64 `json:"unique_id" gorm:"primaryKey"`
	// AddedBy is the id of the person who added the bot and has the
	// rights to edit stuff.
	AddedBy           int64 `json:"added_by"`
	IgnoreMediaGroups bool  `json:"ignore_media_groups"`
	// IgnoreRepeatChecker field is set to true only if the bot is not
	// supposed to check for repeating post and remove them.
	IgnoreRepeatChecker bool `json:"ignore_repeat_checker"`
	IsTmpIgnoring       bool `json:"is_tmp_ignoring" sql:"-" gorm:"-"`
}
