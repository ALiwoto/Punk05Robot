package repostPlugin

import (
	"time"

	"github.com/AnimeKaizoku/Punk05Robot/src/core/downloadUtils"
	"github.com/AnimeKaizoku/Punk05Robot/src/core/logging"
	wv "github.com/AnimeKaizoku/Punk05Robot/src/core/wotoValues"
	"github.com/AnimeKaizoku/Punk05Robot/src/database"
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
		_, _ = b.LeaveChat(ctx.MyChatMember.Chat.Id, nil)
	}
	return nil
}

func repostMessageFilter(msg *gotgbot.Message) bool {
	if !database.IsChannelRegistered(msg.Chat.Id) {
		_, _ = wv.HelperBot.LeaveChat(msg.Chat.Id, nil)
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
	} else if msg.Text != "" && !settings.AllowUploadFromUrl {
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
		_, err := ctx.EffectiveMessage.Delete(b, nil)
		if err != nil {
			logging.Error("while deleting: ", err)
		}
		shouldDeleteMessage = false
	} else {
		shouldDeleteMessage = true
	}

	job := &wv.PendingJob{
		Bot:                 b,
		Ctx:                 ctx,
		Settings:            settings,
		ShouldDeleteMessage: shouldDeleteMessage,
		UrlUploadHandler:    downloadUtils.GetUrlUploaderHandler(msg.Text),
		Handler:             handleRepost,
		RegisteredTime:      time.Now(),
		TimeDistance:        distance,
	}

	if settings.AllowCaption && settings.AllowFooterText {
		if settings.FooterText != "" || settings.RepostingMode == wv.RepostingModeWithOriginalContext {
			job.CaptionGetter = getCaption
		}
	}

	if settings.AllowButtons {
		if !settings.ButtonsUniqueId.IsEmpty() {
			job.ButtonGenerator = getButtons
		}
	}

	if msg.MediaGroupId != "" {
		addJobToMediaGroupMessagesMap(chat.Id, job)
		updateAllMediaGroupsRegisteredTime(chat.Id, job.RegisteredTime)
	}

	wv.PendingJobs.Add(generateKey(), job)

	return nil
}

func updateAllMediaGroupsRegisteredTime(chatId int64, t time.Time) {
	jobs := mediaGroupMessagesMap.GetValue(chatId)
	if len(jobs) == 0 {
		return
	}

	for _, current := range jobs {
		current.RegisteredTime = t
	}
	mediaGroupMessagesMap.Set(chatId, jobs)
}

func addJobToMediaGroupMessagesMap(id int64, j *wv.PendingJob) {
	jobs := mediaGroupMessagesMap.GetValue(id)
	jobs = append(jobs, j)
	mediaGroupMessagesMap.Set(id, jobs)
}

func handleRepost(job *wv.PendingJob) error {
	msg := job.Ctx.EffectiveMessage
	if job.ShouldDeleteMessage {
		_, err := msg.Delete(job.Bot, nil)
		if err != nil {
			logging.Error("while deleting: ", err)
		}
		job.ShouldDeleteMessage = false
	}

	chat := msg.Chat
	bot := job.Bot
	var err error

	if job.UrlUploadHandler != nil {
		// TODO: Support other kinds of url later, for now, we only support
		// Twitter. Pixiv is planned to be supported in future version.
		media, err := job.UrlUploadHandler(msg.Text)
		if err != nil {
			// TODO: add a new field to channel settings called `log_chat`, and send the error
			// to that chat as well.
			logging.Error(err)
			return err
		}

		job.MediaOnCaptionName = media.Owner
		job.MediaOnCaptionUrl = msg.Text

		// TODO: Add support for sending a media group (photo album), here, if urls are multiple.
		_, err = bot.SendPhoto(chat.Id, media.Urls[0], &gotgbot.SendPhotoOpts{
			Caption:     job.GetPostCaption(),
			ReplyMarkup: job.GenerateButtons(),
			ParseMode:   gotgbot.ParseModeMarkdownV2,
		})
		return err
	}

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
			Caption:     job.GetPostCaption(),
			ReplyMarkup: job.GenerateButtons(),
			ParseMode:   gotgbot.ParseModeMarkdownV2,
		})
	case msg.Video != nil:
		_, err = bot.SendVideo(chat.Id, msg.Video.FileId, &gotgbot.SendVideoOpts{
			Caption:     job.GetPostCaption(),
			ReplyMarkup: job.GenerateButtons(),
			ParseMode:   gotgbot.ParseModeMarkdownV2,
		})
	case msg.Audio != nil:
		_, err = bot.SendAudio(chat.Id, msg.Audio.FileId, &gotgbot.SendAudioOpts{
			Caption:     job.GetPostCaption(),
			ReplyMarkup: job.GenerateButtons(),
			ParseMode:   gotgbot.ParseModeMarkdownV2,
		})
	case msg.Voice != nil:
		_, err = bot.SendVoice(chat.Id, msg.Voice.FileId, &gotgbot.SendVoiceOpts{
			Caption:     job.GetPostCaption(),
			ReplyMarkup: job.GenerateButtons(),
			ParseMode:   gotgbot.ParseModeMarkdownV2,
		})
	case msg.Sticker != nil:
		_, err = bot.SendSticker(chat.Id, msg.Sticker.FileId, &gotgbot.SendStickerOpts{
			//Caption: w.Text,
			ReplyMarkup: job.GenerateButtons(),
		})
	case msg.Document != nil:
		_, err = bot.SendDocument(chat.Id, msg.Document.FileId, &gotgbot.SendDocumentOpts{
			Caption:     job.GetPostCaption(),
			ReplyMarkup: job.GenerateButtons(),
			ParseMode:   gotgbot.ParseModeMarkdownV2,
		})
	case msg.VideoNote != nil:
		_, err = bot.SendVideoNote(chat.Id, msg.VideoNote.FileId, &gotgbot.SendVideoNoteOpts{
			//Caption: w.Text,
			ReplyMarkup: job.GenerateButtons(),
		})
	case msg.Animation != nil:
		_, err = bot.SendAnimation(chat.Id, msg.Animation.FileId, &gotgbot.SendAnimationOpts{
			Caption:     job.GetPostCaption(),
			ReplyMarkup: job.GenerateButtons(),
			ParseMode:   gotgbot.ParseModeMarkdownV2,
		})
	}

	return err
}
