package repostPlugin

import (
	"strconv"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

func isMediaMessage(msg *gotgbot.Message) bool {
	switch {
	case len(msg.Photo) > 0:
	case msg.Video != nil:
	case msg.Audio != nil:
	case msg.Voice != nil:
	case msg.Sticker != nil:
	case msg.Document != nil:
	case msg.VideoNote != nil:
	case msg.Animation != nil:
		return true
	}

	return false
}

func generateButtons() *gotgbot.InlineKeyboardMarkup {
	return nil
}

func generateKey(msg *gotgbot.Message) string {
	return strconv.FormatInt(msg.Chat.Id, 10) + "_" + strconv.FormatInt(msg.MessageId, 10)
}

func LoadAllHandlers(d *ext.Dispatcher, t []rune) {
	repostMessageHandler := handlers.NewMessage(repostMessageFilter, repostMessageHandler)
	repostMessageHandler.AllowChannel = true

	d.AddHandler(repostMessageHandler)
}
