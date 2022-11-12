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

// ChannelSettings struct contains information about settings of a registered channel.
type ChannelSettings struct {
	// ChannelId is the channel id.
	ChannelId int64 `json:"unique_id" gorm:"primaryKey"`

	// AddedBy is the id of the person who added the bot and has the
	// rights to edit stuff.
	AddedBy int64 `json:"added_by"`

	// IgnoreMediaGroups determines whether bot should ignores media groups in
	// this channel or not.
	IgnoreMediaGroups bool `json:"ignore_media_groups"`

	// IgnoreRepeatChecker field is set to true only if the bot is not
	// supposed to check for repeating post and remove them.
	IgnoreRepeatChecker bool `json:"ignore_repeat_checker"`

	// IsTmpIgnoring determines whether the bot is currently temporarily ignoring
	// all errors received from this channel or not.
	// WARNING: as the name suggests, this field is *temporary*, means it
	// won't get inserted into the db at all, by restarting the bot, this field
	// will go back to its default value (`false``).
	IsTmpIgnoring bool `json:"is_tmp_ignoring" sql:"-" gorm:"-"`

	// AllowUploadFromUrl determines whether the bot should fetch the content from
	// the given url and upload it to the channel or not.
	AllowUploadFromUrl bool `json:"allow_upload_from_url"`

	// AllowCaption determines whether the bot should post any caption on posts
	// or not.
	AllowCaption bool `json:"allow_caption"`

	// AllowFooterText determines whether the bot should use any footer text on posts
	// or not.
	AllowFooterText bool `json:"allow_footer_text" default:"true"`

	// AllowButtons determines whether the bot should parse and use buttons for posts
	// or not.
	AllowButtons bool `json:"allow_buttons" default:"true"`

	// RepostingMode is the mode applied to a channel by bot.
	RepostingMode ChannelRepostingMode `json:"channel_reposting_mode"`

	// FooterText is the text that should be put on every post as the footer, if and
	// only if `AllowFooterText` and `AllowCaption` fields are both set to  `true`.
	FooterText string `json:"footer_text"`

	// ButtonsUniqueId is the unique id of the buttons which will be added to the posts
	// on reposting, if and only if the `AllowButtons` field is set to `true`.
	ButtonsUniqueId ButtonsUniqueId `json:"buttons_unique_id"`

	// AccessMap is a safe map of `ChannelAccessElement`, determining the users that
	// have access to this channel's settings. This feature is not fully implemented
	// at this moment, this field is just here as a reminder to implement this feature in
	// future versions of the bot.
	// WARNING: This field doesn't get saved to the db, it has to be loaded from the db
	// whenever someone tries to access the settings of this channel.
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
