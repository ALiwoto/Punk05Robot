package channelsPlugin

import (
	"strconv"

	"github.com/ALiwoto/mdparser/mdparser"
	"github.com/AnimeKaizoku/Punk05Robot/src/core/wotoConfig"
	wv "github.com/AnimeKaizoku/Punk05Robot/src/core/wotoValues"
	"github.com/AnimeKaizoku/Punk05Robot/src/database"
	"github.com/AnimeKaizoku/ssg/ssg"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func registerCommandResponse(b *gotgbot.Bot, ctx *ext.Context) error {
	user := ctx.EffectiveUser
	msg := ctx.EffectiveMessage
	if !wotoConfig.IsSudoUser(user.Id) {
		return ext.ContinueGroups
	}

	myStrs := ssg.Split(msg.Text, " ", "\n")
	if len(myStrs) < 2 {
		md := mdparser.GetBold("Usage: ").TabThis()
		md.Normal("/register ").Mono("channel_id")
		_, _ = msg.Reply(b, md.ToString(), &gotgbot.SendMessageOpts{
			ParseMode: wv.MarkdownV2,
		})

		return ext.EndGroups
	}

	channelId := ssg.ToInt64(myStrs[1])
	if channelId == 0 {
		md := mdparser.GetNormal("Invalid chat-id provided.")
		_, _ = msg.Reply(b, md.ToString(), &gotgbot.SendMessageOpts{
			ParseMode: wv.MarkdownV2,
		})
		return ext.EndGroups
	}

	settings := database.GetChannelSettings(channelId)
	if settings != nil {
		md := mdparser.GetNormal("Channel with id ")
		md.Mono(strconv.FormatInt(channelId, 10))
		md.Normal(" is already registered.")
		_, _ = msg.Reply(b, md.ToString(), &gotgbot.SendMessageOpts{
			ParseMode: wv.MarkdownV2,
		})
		return ext.EndGroups
	}

	database.SaveChannelSettings(&wv.ChannelSettings{
		ChannelId: channelId,
		AddedBy:   user.Id,
	}, true)

	md := mdparser.GetNormal("Successfully registered channel with id ")
	md.Mono(ssg.ToBase10(channelId)).Normal(".")
	_, _ = msg.Reply(b, md.ToString(), &gotgbot.SendMessageOpts{
		ParseMode: wv.MarkdownV2,
	})

	return ext.EndGroups
}

func tmpIgnoreResponse(b *gotgbot.Bot, ctx *ext.Context) error {
	user := ctx.EffectiveUser
	msg := ctx.EffectiveMessage
	if !wotoConfig.IsSudoUser(user.Id) {
		return ext.ContinueGroups
	}

	myStrs := ssg.Split(msg.Text, " ", "\n")
	if len(myStrs) < 2 {
		md := mdparser.GetBold("Usage: ").TabThis()
		md.Normal("/tmpIgnore ").Mono("channel_id")
		_, _ = msg.Reply(b, md.ToString(), &gotgbot.SendMessageOpts{
			ParseMode: wv.MarkdownV2,
		})

		return ext.EndGroups
	}

	channelId := ssg.ToInt64(myStrs[1])
	if channelId == 0 {
		md := mdparser.GetNormal("Invalid chat-id provided.")
		_, _ = msg.Reply(b, md.ToString(), &gotgbot.SendMessageOpts{
			ParseMode: wv.MarkdownV2,
		})
		return ext.EndGroups
	}

	settings := database.GetChannelSettings(channelId)
	if settings == nil {
		md := mdparser.GetNormal("Channel is not registered.")
		_, _ = msg.Reply(b, md.ToString(), &gotgbot.SendMessageOpts{
			ParseMode: wv.MarkdownV2,
		})
		return ext.EndGroups
	}

	if settings.IsTmpIgnoring {
		settings.IsTmpIgnoring = false
		md := mdparser.GetNormal("I won't ignore this channel anymore.")
		_, _ = msg.Reply(b, md.ToString(), &gotgbot.SendMessageOpts{
			ParseMode: wv.MarkdownV2,
		})
		return ext.EndGroups
	} else {
		settings.IsTmpIgnoring = true
		md := mdparser.GetNormal("I will temporary ignore this channel from now on.")
		_, _ = msg.Reply(b, md.ToString(), &gotgbot.SendMessageOpts{
			ParseMode: wv.MarkdownV2,
		})
		return ext.EndGroups
	}
}

func addUserResponse(b *gotgbot.Bot, ctx *ext.Context) error {
	//TODO
	return nil
}

func setFooterResponse(b *gotgbot.Bot, ctx *ext.Context) error {
	//TODO
	return nil
}

func setButtonsResponse(b *gotgbot.Bot, ctx *ext.Context) error {
	//TODO
	return nil
}

func addButtonsResponse(b *gotgbot.Bot, ctx *ext.Context) error {
	//TODO
	return nil
}

func removeButtonsResponse(b *gotgbot.Bot, ctx *ext.Context) error {
	//TODO
	return nil
}
