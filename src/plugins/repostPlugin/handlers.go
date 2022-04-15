package repostPlugin

import (
	wv "github.com/AnimeKaizoku/RepostingRobot/src/core/wotoValues"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func repostMessageFilter(msg *gotgbot.Message) bool {
	return isMediaMessage(msg) && msg.Chat.Type == "channel"
}

func repostMessageHandler(b *gotgbot.Bot, ctx *ext.Context) error {
	wv.PendingJobs.Add(generateKey(ctx.EffectiveMessage), &wv.PendingJob{
		Bot:     b,
		Ctx:     ctx,
		Handler: handleRepost,
	})

	return nil
}

func handleRepost(job *wv.PendingJob) error {
	msg := job.Ctx.EffectiveMessage
	chat := msg.Chat
	bot := job.Bot
	theCaption := "@" + chat.Username
	if len(theCaption) < 5 {
		theCaption = ""
	}

	_, _ = msg.Delete(bot)

	var err error

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
		_, err = bot.SendPhoto(chat.Id, msg.Photo[0].FileId, &gotgbot.SendPhotoOpts{
			Caption:     theCaption,
			ReplyMarkup: generateButtons(),
		})
	case msg.Video != nil:
		_, err = bot.SendVideo(chat.Id, msg.Video.FileId, &gotgbot.SendVideoOpts{
			Caption:     theCaption,
			ReplyMarkup: generateButtons(),
		})
	case msg.Audio != nil:
		_, err = bot.SendAudio(chat.Id, msg.Audio.FileId, &gotgbot.SendAudioOpts{
			Caption:     theCaption,
			ReplyMarkup: generateButtons(),
		})
	case msg.Voice != nil:
		_, err = bot.SendVoice(chat.Id, msg.Voice.FileId, &gotgbot.SendVoiceOpts{
			Caption:     theCaption,
			ReplyMarkup: generateButtons(),
		})
	case msg.Sticker != nil:
		_, err = bot.SendSticker(chat.Id, msg.Sticker.FileId, &gotgbot.SendStickerOpts{
			//Caption: w.Text,
			ReplyMarkup: generateButtons(),
		})
	case msg.Document != nil:
		_, err = bot.SendDocument(chat.Id, msg.Document.FileId, &gotgbot.SendDocumentOpts{
			Caption:     theCaption,
			ReplyMarkup: generateButtons(),
		})
	case msg.VideoNote != nil:
		_, err = bot.SendVideoNote(chat.Id, msg.VideoNote.FileId, &gotgbot.SendVideoNoteOpts{
			//Caption: w.Text,
			ReplyMarkup: generateButtons(),
		})
	case msg.Animation != nil:
		_, err = bot.SendAnimation(chat.Id, msg.Animation.FileId, &gotgbot.SendAnimationOpts{
			Caption:     theCaption,
			ReplyMarkup: generateButtons(),
		})
	}

	return err
}
