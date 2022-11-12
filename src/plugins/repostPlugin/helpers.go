package repostPlugin

import (
	"time"

	"github.com/ALiwoto/mdparser/mdparser"
	"github.com/AnimeKaizoku/Punk05Robot/src/core/utils"
	wv "github.com/AnimeKaizoku/Punk05Robot/src/core/wotoValues"
	"github.com/AnimeKaizoku/ssg/ssg"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

func getCaption(job *wv.PendingJob) string {
	md := mdparser.GetNormal(job.Settings.FooterText)
	md.ReplaceMdThis(
		mdparser.GetNormal("CHANNEL_USERNAME"),
		mdparser.GetNormal("@"+job.Ctx.EffectiveChat.Username),
	)
	md.ReplaceMdThis(
		mdparser.GetNormal("CHANNEL_TITLE"),
		mdparser.GetNormal(job.Ctx.EffectiveChat.Title),
	)
	md.ReplaceMdThis(
		mdparser.GetNormal("CHANNEL_ID"),
		mdparser.GetNormal(ssg.ToBase10(job.Ctx.EffectiveChat.Id)),
	)
	md.ReplaceMdThis(
		mdparser.GetNormal("MORE_CONTENTS"),
		wv.MoreContentsMd,
	)
	md.ReplaceMdThis(
		mdparser.GetNormal("MORE_CONTENT"),
		wv.MoreContentsMd,
	)
	return md.ToString()
}

func getButtons(job *wv.PendingJob) *gotgbot.InlineKeyboardMarkup {
	// TODO: Add support for getting buttons from db using their unique id.
	return wv.DefaultImplementedButtons[job.Settings.ButtonsUniqueId]
}

func isMediaMessage(msg *gotgbot.Message) bool {
	switch {
	case len(msg.Photo) > 0:
		return true
	case msg.Video != nil:
		return true
	case msg.Audio != nil:
		return true
	case msg.Voice != nil:
		return true
	case msg.Sticker != nil:
		return true
	case msg.Document != nil:
		return true
	case msg.VideoNote != nil:
		return true
	case msg.Animation != nil:
		return true
	}

	return msg.Text != "" && utils.IsSupportedUploadingUrl(msg.Text)
}

func getFilesId(msg *gotgbot.Message) string {
	switch {
	//WhisperTypePhoto
	//WhisperTypeVideo
	//WhisperTypeAudio
	//WhisperTypeVoice
	//WhisperTypeSticker
	//WhisperTypeDocument
	//WhisperTypeVideoNote
	//WhisperTypeAnimation
	//WhisperTypeDice
	case len(msg.Photo) != 0:
		return msg.Photo[0].FileId + "/" + msg.Photo[0].FileUniqueId
	case msg.Video != nil:
		return msg.Video.FileId + "/" + msg.Video.FileUniqueId
	case msg.Audio != nil:
		return msg.Audio.FileId + "/" + msg.Audio.FileUniqueId
	case msg.Voice != nil:
		return msg.Voice.FileId + "/" + msg.Voice.FileUniqueId
	case msg.Sticker != nil:
		return msg.Sticker.FileId + "/" + msg.Sticker.FileUniqueId
	case msg.Document != nil:
		return msg.Document.FileId + "/" + msg.Document.FileUniqueId
	case msg.VideoNote != nil:
		return msg.VideoNote.FileId + "/" + msg.VideoNote.FileUniqueId
	case msg.Animation != nil:
		return msg.Animation.FileId + "/" + msg.Animation.FileUniqueId
	}

	return ""
}

func generateKey() uint64 {
	return jobsKeyGenerator.Next()
}

func LoadAllHandlers(d *ext.Dispatcher, t []rune) {
	repostMessageHandler := &channelPost{
		Filter:   repostMessageFilter,
		Response: repostMessageResponse,
	}
	myChatAddedHandler := handlers.NewMyChatMember(chatMemberFilter, chatMemberResponse)

	d.AddHandler(repostMessageHandler)
	d.AddHandler(myChatAddedHandler)
}

func _getRepeatCheckerMap() *ssg.SafeEMap[string, bool] {
	m := ssg.NewSafeEMap[string, bool]()
	m.SetInterval(45 * time.Second)
	m.SetExpiration(5 * mediaGroupDistance)
	m.EnableChecking()

	return m
}

func _getMediaGroupMessagesMap() *ssg.SafeEMap[int64, []*wv.PendingJob] {
	m := ssg.NewSafeEMap[int64, []*wv.PendingJob]()
	m.SetInterval(time.Minute)
	m.SetExpiration(15 * mediaGroupDistance)
	m.EnableChecking()

	return m
}
