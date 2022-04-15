package utils

import (
	"log"
	"strconv"
	"strings"

	"github.com/ALiwoto/mdparser/mdparser"
	"github.com/AnimeKaizoku/RepostingRobot/src/core/logging"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func GetIdFromToken(token string) int64 {
	if !strings.Contains(token, ":") {
		return 0
	}

	i, err := strconv.ParseInt(strings.Split(token, ":")[0], 10, 64)
	if err != nil {
		return 0
	}

	return i
}

func SendAlert(b *gotgbot.Bot, m *gotgbot.Message, md mdparser.WMarkDown) error {
	str := md.ToString()
	str = strings.ReplaceAll(str, b.Token, "")
	_, err := m.Reply(b, str, &gotgbot.SendMessageOpts{ParseMode: MarkDownV2})
	if err != nil {
		log.Println(err)
	}

	return nil
}

func SafeReply(b *gotgbot.Bot, ctx *ext.Context, output string) error {
	msg := ctx.EffectiveMessage
	if len(output) < 4096 {
		_, err := msg.Reply(b, output,
			&gotgbot.SendMessageOpts{ParseMode: MarkDownV2})
		if err != nil {
			logging.Error("got an error when trying to send results: ", err)
			return err
		}
	} else {
		_, err := b.SendDocument(ctx.EffectiveChat.Id, []byte(output), &gotgbot.SendDocumentOpts{
			ReplyToMessageId: msg.MessageId,
		})
		if err != nil {
			logging.Error("got an error when trying to send document: ", err)
			return err
		}
	}

	return nil
}
