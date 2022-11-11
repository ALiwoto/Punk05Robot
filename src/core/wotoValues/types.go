package wotoValues

import (
	"time"

	"github.com/AnimeKaizoku/ssg/ssg"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

type PendingJob struct {
	Bot                 *gotgbot.Bot
	Ctx                 *ext.Context
	Settings            *ChannelSettings
	ShouldDeleteMessage bool

	Handler         func(job *PendingJob) error
	ButtonGenerator func(job *PendingJob) *gotgbot.InlineKeyboardMarkup
	CaptionGetter   func(job *PendingJob) string

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
	AllowUploadFromUrl  bool `json:"allow_upload_from_url"`
	// RepostingMode is the mode applied to a channel by bot.
	RepostingMode   ChannelRepostingMode `json:"channel_reposting_mode"`
	FooterText      string               `json:"footer_text"`
	ButtonsUniqueId ButtonsUniqueId      `json:"buttons_unique_id"`

	AccessMap *ssg.SafeMap[int64, ChannelAccessElement] `json:"-" sql:"-" gorm:"-"`
}

type ChannelAccessElement struct {
	AccessUniqueId string `gorm:"primaryKey"`
	UserId         int64
	ChannelId      int64
	AddedBy        int64
}

type ChannelRepostingMode int
type ButtonsUniqueId string
