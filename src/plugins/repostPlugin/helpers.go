package repostPlugin

import (
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

func generateButtons() *gotgbot.InlineKeyboardMarkup {
	return _defaultButtons
}

func generateKey() uint64 {
	keyGeneratorMutex.Lock()
	defer keyGeneratorMutex.Unlock()

	lastKey++
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
