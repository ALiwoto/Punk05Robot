package repostPlugin

import (
	"time"

	"github.com/AnimeKaizoku/ssg/ssg"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

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

	return false
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

func generateButtons() *gotgbot.InlineKeyboardMarkup {
	return _defaultButtons
}

func generateKey() uint64 {
	keyGeneratorMutex.Lock()
	lastKey++
	keyGeneratorMutex.Unlock()

	return lastKey
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
	m.SetInterval(5 * time.Second)
	m.SetExpiration(2 * mediaGroupDistance)
	m.EnableChecking()

	return m
}
