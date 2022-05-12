package repostPlugin

import (
	"time"

	"github.com/AnimeKaizoku/RepostingRobot/src/core/logging"
	wv "github.com/AnimeKaizoku/RepostingRobot/src/core/wotoValues"
	"github.com/AnimeKaizoku/RepostingRobot/src/database"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func chatMemberFilter(u *gotgbot.ChatMemberUpdated) bool {
	return u.NewChatMember.GetUser().Id == wv.HelperBot.Id
}

func chatMemberResponse(b *gotgbot.Bot, ctx *ext.Context) error {
	chatMember := ctx.MyChatMember.NewChatMember
	if chatMember == nil {
		return nil
	}

	if chatMember.GetStatus() == "left" {
		// bot is leaving, don't handle the update...
		return nil
	}

	if !database.IsChannelRegistered(ctx.MyChatMember.Chat.Id) {
		_, _ = b.LeaveChat(ctx.MyChatMember.Chat.Id)
	}
	return nil
}

func repostMessageFilter(msg *gotgbot.Message) bool {
	if !database.IsChannelRegistered(msg.Chat.Id) {
		_, _ = wv.HelperBot.LeaveChat(msg.Chat.Id)
		return false
	}

	return isMediaMessage(msg) && msg.Chat.Type == "channel"
}

func repostMessageResponse(b *gotgbot.Bot, ctx *ext.Context) error {
	if ctx.EditedChannelPost != nil {
		return ext.ContinueGroups
	}

	msg := ctx.EffectiveMessage
	chat := ctx.EffectiveChat
	var distance time.Duration
	shouldDeleteMessage := true

	settings := database.GetChannelSettings(chat.Id)
	if settings.IsTmpIgnoring {
		return ext.ContinueGroups
	}

	if msg.MediaGroupId != "" {
		if settings.IgnoreMediaGroups {
			return ext.ContinueGroups
		}

		if !settings.IgnoreRepeatChecker {
			fId := getFilesId(msg)
			if repeatCheckerMap.GetValue(fId) {
				// this post is repeated
				return ext.ContinueGroups
			}

			repeatCheckerMap.Set(fId, true)
		}

		shouldDeleteMessage = false
		distance = mediaGroupDistance
	}

	if shouldDeleteMessage {
		_, err := ctx.EffectiveMessage.Delete(b)
		if err != nil {
			logging.Error("while deleteing: ", err)
		}
		shouldDeleteMessage = false
	} else {
		shouldDeleteMessage = true
	}

	wv.PendingJobs.Add(generateKey(), &wv.PendingJob{
		Bot:                 b,
		Ctx:                 ctx,
		ShouldDeleteMessage: shouldDeleteMessage,
		Handler:             handleRepost,
		RegisteredTime:      time.Now(),
		TimeDistance:        distance,
	})

	return nil
}

func handleRepost(job *wv.PendingJob) error {
	msg := job.Ctx.EffectiveMessage
	if job.ShouldDeleteMessage {
		_, err := msg.Delete(job.Bot)
		if err != nil {
			logging.Error("while deleteing: ", err)
		}
		job.ShouldDeleteMessage = false
	}

	chat := msg.Chat
	bot := job.Bot
	theCaption := "@" + chat.Username
	if len(theCaption) < 5 {
		theCaption = ""
	}

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
