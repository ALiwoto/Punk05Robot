package channelsPlugin

import (
	"strconv"

	"github.com/ALiwoto/mdparser/mdparser"
	wv "github.com/AnimeKaizoku/Punk05Robot/src/core/wotoValues"
	"github.com/AnimeKaizoku/Punk05Robot/src/database"
	"github.com/AnimeKaizoku/ssg/ssg"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func registerCommandResponse(b *gotgbot.Bot, ctx *ext.Context) error {
	user := ctx.EffectiveUser
	msg := ctx.EffectiveMessage

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
	msg := ctx.EffectiveMessage

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
	message := ctx.EffectiveMessage
	targetString := ""
	idStr := ""
	var channelId int64

	if message.ReplyToMessage != nil {
		targetString = message.ReplyToMessage.Text
		myStrs := ctx.Args()
		if len(myStrs) < 2 {
			txt := mdparser.GetBold("Usage: \n")
			txt.Mono("\t\t/setFooter -100123456 TEXT HERE (or reply)")
			_, _ = message.Reply(b, txt.ToString(), &gotgbot.SendMessageOpts{
				ParseMode: gotgbot.ParseModeMarkdownV2,
			})
			return ext.EndGroups
		}

		idStr = myStrs[1]
	} else {
		// /setFooter ID TEXT
		// 0: /setFooter
		// 1: The channel ID
		// 2: rest (which is text)
		myStrs := ssg.SplitN(message.Text, 3, " ")
		if len(myStrs) < 3 {
			txt := mdparser.GetNormal("Usage: ")
			txt.Mono("/setFooter -100123456 TEXT HERE (or reply)")
			_, _ = message.Reply(b, txt.ToString(), &gotgbot.SendMessageOpts{
				ParseMode: gotgbot.ParseModeMarkdownV2,
			})
			return ext.EndGroups
		}
		idStr = myStrs[1]
		targetString = myStrs[2]
	}

	channelId = ssg.ToInt64(idStr)
	if channelId >= 0 {
		txt := mdparser.GetNormal("Usage: ")
		txt.Mono("\t\t/setFooter -100123456 TEXT HERE\n")
		txt.Bold("Please make sure you have entered a correct channel ID.\n")
		txt.Normal("Channel IDs should always start with -100.")
		_, _ = message.Reply(b, txt.ToString(), &gotgbot.SendMessageOpts{
			ParseMode: gotgbot.ParseModeMarkdownV2,
		})
		return ext.EndGroups
	}

	settings := database.GetChannelSettings(channelId)
	if settings == nil {
		txt := mdparser.GetBold("Looks like this channel's settings doesn't exist in my database.\n")
		txt.Normal("You have to register the channel using:\n")
		txt.Mono("\t\t/register CHANNEL_ID (e.g. -10012345678)")
		_, _ = message.Reply(b, txt.ToString(), &gotgbot.SendMessageOpts{
			ParseMode: gotgbot.ParseModeMarkdownV2,
		})
		return ext.EndGroups
	}

	settings.FooterText = targetString
	database.SaveChannelSettings(settings, false)

	return ext.EndGroups
}

func setButtonsResponse(b *gotgbot.Bot, ctx *ext.Context) error {
	message := ctx.EffectiveMessage
	targetString := ""
	idStr := ""
	var channelId int64

	if message.ReplyToMessage != nil {
		targetString = message.ReplyToMessage.Text
		myStrs := ctx.Args()
		if len(myStrs) < 2 {
			txt := mdparser.GetNormal("Usage: ")
			txt.Mono("/setButtons -100123456 TEXT HERE (or reply)")
			_, _ = message.Reply(b, txt.ToString(), &gotgbot.SendMessageOpts{
				ParseMode: gotgbot.ParseModeMarkdownV2,
			})
			return ext.EndGroups
		}

		idStr = myStrs[1]
	} else {
		// /setFooter ID TEXT
		// 0: /setFooter
		// 1: The channel ID
		// 2: rest (which is text)
		myStrs := ssg.SplitN(message.Text, 3, " ")
		if len(myStrs) < 3 {
			txt := mdparser.GetNormal("Usage: ")
			txt.Mono("/setButtons -100123456 TEXT HERE (or reply)")
			_, _ = message.Reply(b, txt.ToString(), &gotgbot.SendMessageOpts{
				ParseMode: gotgbot.ParseModeMarkdownV2,
			})
			return ext.EndGroups
		}
		idStr = myStrs[1]
		targetString = myStrs[2]
	}

	channelId = ssg.ToInt64(idStr)
	if channelId >= 0 {
		txt := mdparser.GetNormal("Usage: ")
		txt.Mono("\t\t/setButtons -100123456 TEXT HERE\n")
		txt.Bold("Please make sure you have entered a correct channel ID.\n")
		txt.Normal("Channel IDs should always start with -100.")
		_, _ = message.Reply(b, txt.ToString(), &gotgbot.SendMessageOpts{
			ParseMode: gotgbot.ParseModeMarkdownV2,
		})
		return ext.EndGroups
	}

	settings := database.GetChannelSettings(channelId)
	if settings == nil {
		txt := mdparser.GetBold("Looks like this channel's settings doesn't exist in my database.\n")
		txt.Normal("You have to register the channel using:\n")
		txt.Mono("\t\t/register CHANNEL_ID (e.g. -10012345678)")
		_, _ = message.Reply(b, txt.ToString(), &gotgbot.SendMessageOpts{
			ParseMode: gotgbot.ParseModeMarkdownV2,
		})
		return ext.EndGroups
	}

	if !wv.IsValidButtonsUniqueId(targetString) {
		// TODO: parse buttons from user input here and allow
		// using custom buttons to the user.
		txt := mdparser.GetBold("Using custom buttons are not supported yet.")
		_, _ = message.Reply(b, txt.ToString(), &gotgbot.SendMessageOpts{
			ParseMode: gotgbot.ParseModeMarkdownV2,
		})
		return ext.EndGroups
	}

	settings.ButtonsUniqueId = wv.ButtonsUniqueId(targetString)
	database.SaveChannelSettings(settings, false)

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

func allowFooterResponse(b *gotgbot.Bot, ctx *ext.Context) error {
	message := ctx.EffectiveMessage
	idStr := ""
	var channelId int64

	// /allowFooter ID TEXT
	// 0: /setFooter
	// 1: The channel ID
	myStrs := ctx.Args()
	if len(myStrs) < 2 {
		txt := mdparser.GetNormal("Usage: ")
		txt.Mono("/allowFooter -100123456")
		_, _ = message.Reply(b, txt.ToString(), &gotgbot.SendMessageOpts{
			ParseMode: gotgbot.ParseModeMarkdownV2,
		})
		return ext.EndGroups
	}
	idStr = myStrs[1]

	channelId = ssg.ToInt64(idStr)
	if channelId >= 0 {
		txt := mdparser.GetNormal("Usage: ")
		txt.Mono("\t\t/allowFooter -100123456\n")
		txt.Bold("Please make sure you have entered a correct channel ID.\n")
		txt.Normal("Channel IDs should always start with -100.")
		_, _ = message.Reply(b, txt.ToString(), &gotgbot.SendMessageOpts{
			ParseMode: gotgbot.ParseModeMarkdownV2,
		})
		return ext.EndGroups
	}

	settings := database.GetChannelSettings(channelId)
	if settings == nil {
		txt := mdparser.GetBold("Looks like this channel's settings doesn't exist in my database.\n")
		txt.Normal("You have to register the channel using:\n")
		txt.Mono("\t\t/register CHANNEL_ID (e.g. -10012345678)")
		_, _ = message.Reply(b, txt.ToString(), &gotgbot.SendMessageOpts{
			ParseMode: gotgbot.ParseModeMarkdownV2,
		})
		return ext.EndGroups
	}

	if settings.AllowFooterText {
		_, _ = message.Reply(b, mdparser.GetNormal("No new settings to be updated!").ToString(), &gotgbot.SendMessageOpts{
			ParseMode: gotgbot.ParseModeMarkdownV2,
		})
		return ext.EndGroups
	}
	settings.AllowFooterText = true
	database.SaveChannelSettings(settings, false)

	_, _ = message.Reply(b, mdparser.GetBold("Channel settings updated!").ToString(), &gotgbot.SendMessageOpts{
		ParseMode: gotgbot.ParseModeMarkdownV2,
	})

	return ext.EndGroups
}

func disallowFooterResponse(b *gotgbot.Bot, ctx *ext.Context) error {
	message := ctx.EffectiveMessage
	idStr := ""
	var channelId int64

	// /allowFooter ID TEXT
	// 0: /setFooter
	// 1: The channel ID
	myStrs := ctx.Args()
	if len(myStrs) < 2 {
		txt := mdparser.GetNormal("Usage: ")
		txt.Mono("/disallowFooter -100123456")
		_, _ = message.Reply(b, txt.ToString(), &gotgbot.SendMessageOpts{
			ParseMode: gotgbot.ParseModeMarkdownV2,
		})
		return ext.EndGroups
	}
	idStr = myStrs[1]

	channelId = ssg.ToInt64(idStr)
	if channelId >= 0 {
		txt := mdparser.GetNormal("Usage: ")
		txt.Mono("\t\t/disallowFooter -100123456\n")
		txt.Bold("Please make sure you have entered a correct channel ID.\n")
		txt.Normal("Channel IDs should always start with -100.")
		_, _ = message.Reply(b, txt.ToString(), &gotgbot.SendMessageOpts{
			ParseMode: gotgbot.ParseModeMarkdownV2,
		})
		return ext.EndGroups
	}

	settings := database.GetChannelSettings(channelId)
	if settings == nil {
		txt := mdparser.GetBold("Looks like this channel's settings doesn't exist in my database.\n")
		txt.Normal("You have to register the channel using:\n")
		txt.Mono("\t\t/register CHANNEL_ID (e.g. -10012345678)")
		_, _ = message.Reply(b, txt.ToString(), &gotgbot.SendMessageOpts{
			ParseMode: gotgbot.ParseModeMarkdownV2,
		})
		return ext.EndGroups
	}

	if !settings.AllowFooterText {
		_, _ = message.Reply(b, mdparser.GetNormal("No new settings to be updated!").ToString(), &gotgbot.SendMessageOpts{
			ParseMode: gotgbot.ParseModeMarkdownV2,
		})
		return ext.EndGroups
	}
	settings.AllowFooterText = false
	database.SaveChannelSettings(settings, false)

	_, _ = message.Reply(b, mdparser.GetBold("Channel settings updated!").ToString(), &gotgbot.SendMessageOpts{
		ParseMode: gotgbot.ParseModeMarkdownV2,
	})

	return ext.EndGroups
}

func allowCaptionResponse(b *gotgbot.Bot, ctx *ext.Context) error {
	message := ctx.EffectiveMessage
	idStr := ""
	var channelId int64

	// /allowCaption ID
	// 0: /allowCaption
	// 1: The channel ID
	myStrs := ctx.Args()
	if len(myStrs) < 2 {
		txt := mdparser.GetNormal("Usage: ")
		txt.Mono("/allowCaption -100123456")
		_, _ = message.Reply(b, txt.ToString(), &gotgbot.SendMessageOpts{
			ParseMode: gotgbot.ParseModeMarkdownV2,
		})
		return ext.EndGroups
	}
	idStr = myStrs[1]

	channelId = ssg.ToInt64(idStr)
	if channelId >= 0 {
		txt := mdparser.GetNormal("Usage: ")
		txt.Mono("\t\t/allowCaption -100123456\n")
		txt.Bold("Please make sure you have entered a correct channel ID.\n")
		txt.Normal("Channel IDs should always start with -100.")
		_, _ = message.Reply(b, txt.ToString(), &gotgbot.SendMessageOpts{
			ParseMode: gotgbot.ParseModeMarkdownV2,
		})
		return ext.EndGroups
	}

	settings := database.GetChannelSettings(channelId)
	if settings == nil {
		txt := mdparser.GetBold("Looks like this channel's settings doesn't exist in my database.\n")
		txt.Normal("You have to register the channel using:\n")
		txt.Mono("\t\t/register CHANNEL_ID (e.g. -10012345678)")
		_, _ = message.Reply(b, txt.ToString(), &gotgbot.SendMessageOpts{
			ParseMode: gotgbot.ParseModeMarkdownV2,
		})
		return ext.EndGroups
	}

	if settings.AllowCaption {
		_, _ = message.Reply(b, mdparser.GetNormal("No new settings to be updated!").ToString(), &gotgbot.SendMessageOpts{
			ParseMode: gotgbot.ParseModeMarkdownV2,
		})
		return ext.EndGroups
	}
	settings.AllowCaption = true
	database.SaveChannelSettings(settings, false)

	_, _ = message.Reply(b, mdparser.GetBold("Channel settings updated!").ToString(), &gotgbot.SendMessageOpts{
		ParseMode: gotgbot.ParseModeMarkdownV2,
	})

	return ext.EndGroups
}

func disallowCaptionResponse(b *gotgbot.Bot, ctx *ext.Context) error {
	message := ctx.EffectiveMessage
	idStr := ""
	var channelId int64

	// /disallowCaption ID
	// 0: /disallowCaption
	// 1: The channel ID
	myStrs := ctx.Args()
	if len(myStrs) < 2 {
		txt := mdparser.GetNormal("Usage: ")
		txt.Mono("/disallowCaption -100123456")
		_, _ = message.Reply(b, txt.ToString(), &gotgbot.SendMessageOpts{
			ParseMode: gotgbot.ParseModeMarkdownV2,
		})
		return ext.EndGroups
	}
	idStr = myStrs[1]

	channelId = ssg.ToInt64(idStr)
	if channelId >= 0 {
		txt := mdparser.GetNormal("Usage: ")
		txt.Mono("\t\t/disallowCaption -100123456\n")
		txt.Bold("Please make sure you have entered a correct channel ID.\n")
		txt.Normal("Channel IDs should always start with -100.")
		_, _ = message.Reply(b, txt.ToString(), &gotgbot.SendMessageOpts{
			ParseMode: gotgbot.ParseModeMarkdownV2,
		})
		return ext.EndGroups
	}

	settings := database.GetChannelSettings(channelId)
	if settings == nil {
		txt := mdparser.GetBold("Looks like this channel's settings doesn't exist in my database.\n")
		txt.Normal("You have to register the channel using:\n")
		txt.Mono("\t\t/register CHANNEL_ID (e.g. -10012345678)")
		_, _ = message.Reply(b, txt.ToString(), &gotgbot.SendMessageOpts{
			ParseMode: gotgbot.ParseModeMarkdownV2,
		})
		return ext.EndGroups
	}

	if !settings.AllowCaption {
		_, _ = message.Reply(b, mdparser.GetNormal("No new settings to be updated!").ToString(), &gotgbot.SendMessageOpts{
			ParseMode: gotgbot.ParseModeMarkdownV2,
		})
		return ext.EndGroups
	}
	settings.AllowCaption = false
	database.SaveChannelSettings(settings, false)

	_, _ = message.Reply(b, mdparser.GetBold("Channel settings updated!").ToString(), &gotgbot.SendMessageOpts{
		ParseMode: gotgbot.ParseModeMarkdownV2,
	})

	return ext.EndGroups
}

func allowButtonsResponse(b *gotgbot.Bot, ctx *ext.Context) error {
	message := ctx.EffectiveMessage
	idStr := ""
	var channelId int64

	// /allowFooter ID TEXT
	// 0: /allowButtons
	// 1: The channel ID
	myStrs := ctx.Args()
	if len(myStrs) < 2 {
		txt := mdparser.GetNormal("Usage: ")
		txt.Mono("/allowButtons -100123456")
		_, _ = message.Reply(b, txt.ToString(), &gotgbot.SendMessageOpts{
			ParseMode: gotgbot.ParseModeMarkdownV2,
		})
		return ext.EndGroups
	}
	idStr = myStrs[1]

	channelId = ssg.ToInt64(idStr)
	if channelId >= 0 {
		txt := mdparser.GetNormal("Usage: ")
		txt.Mono("\t\t/allowButtons -100123456\n")
		txt.Bold("Please make sure you have entered a correct channel ID.\n")
		txt.Normal("Channel IDs should always start with -100.")
		_, _ = message.Reply(b, txt.ToString(), &gotgbot.SendMessageOpts{
			ParseMode: gotgbot.ParseModeMarkdownV2,
		})
		return ext.EndGroups
	}

	settings := database.GetChannelSettings(channelId)
	if settings == nil {
		txt := mdparser.GetBold("Looks like this channel's settings doesn't exist in my database.\n")
		txt.Normal("You have to register the channel using:\n")
		txt.Mono("\t\t/register CHANNEL_ID (e.g. -10012345678)")
		_, _ = message.Reply(b, txt.ToString(), &gotgbot.SendMessageOpts{
			ParseMode: gotgbot.ParseModeMarkdownV2,
		})
		return ext.EndGroups
	}

	if settings.AllowButtons {
		_, _ = message.Reply(b, mdparser.GetNormal("No new settings to be updated!").ToString(), &gotgbot.SendMessageOpts{
			ParseMode: gotgbot.ParseModeMarkdownV2,
		})
		return ext.EndGroups
	}
	settings.AllowButtons = true
	database.SaveChannelSettings(settings, false)

	_, _ = message.Reply(b, mdparser.GetBold("Channel settings updated!").ToString(), &gotgbot.SendMessageOpts{
		ParseMode: gotgbot.ParseModeMarkdownV2,
	})

	return ext.EndGroups
}

func disallowButtonsResponse(b *gotgbot.Bot, ctx *ext.Context) error {
	message := ctx.EffectiveMessage
	idStr := ""
	var channelId int64

	// /allowFooter ID TEXT
	// 0: /disallowButtons
	// 1: The channel ID
	myStrs := ctx.Args()
	if len(myStrs) < 2 {
		txt := mdparser.GetNormal("Usage: ")
		txt.Mono("/disallowButtons -100123456")
		_, _ = message.Reply(b, txt.ToString(), &gotgbot.SendMessageOpts{
			ParseMode: gotgbot.ParseModeMarkdownV2,
		})
		return ext.EndGroups
	}
	idStr = myStrs[1]

	channelId = ssg.ToInt64(idStr)
	if channelId >= 0 {
		txt := mdparser.GetNormal("Usage: ")
		txt.Mono("\t\t/disallowButtons -100123456\n")
		txt.Bold("Please make sure you have entered a correct channel ID.\n")
		txt.Normal("Channel IDs should always start with -100.")
		_, _ = message.Reply(b, txt.ToString(), &gotgbot.SendMessageOpts{
			ParseMode: gotgbot.ParseModeMarkdownV2,
		})
		return ext.EndGroups
	}

	settings := database.GetChannelSettings(channelId)
	if settings == nil {
		txt := mdparser.GetBold("Looks like this channel's settings doesn't exist in my database.\n")
		txt.Normal("You have to register the channel using:\n")
		txt.Mono("\t\t/register CHANNEL_ID (e.g. -10012345678)")
		_, _ = message.Reply(b, txt.ToString(), &gotgbot.SendMessageOpts{
			ParseMode: gotgbot.ParseModeMarkdownV2,
		})
		return ext.EndGroups
	}

	if !settings.AllowButtons {
		_, _ = message.Reply(b, mdparser.GetNormal("No new settings to be updated!").ToString(), &gotgbot.SendMessageOpts{
			ParseMode: gotgbot.ParseModeMarkdownV2,
		})
		return ext.EndGroups
	}
	settings.AllowButtons = false
	database.SaveChannelSettings(settings, false)

	_, _ = message.Reply(b, mdparser.GetBold("Channel settings updated!").ToString(), &gotgbot.SendMessageOpts{
		ParseMode: gotgbot.ParseModeMarkdownV2,
	})

	return ext.EndGroups
}
