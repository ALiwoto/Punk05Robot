package channelsPlugin

import (
	"github.com/ALiwoto/mdparser/mdparser"
	"github.com/AnimeKaizoku/RepostingRobot/src/core/wotoConfig"
	wv "github.com/AnimeKaizoku/RepostingRobot/src/core/wotoValues"
	"github.com/AnimeKaizoku/RepostingRobot/src/database"
	"github.com/AnimeKaizoku/ssg/ssg"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func registerCommandResponse(b *gotgbot.Bot, ctx *ext.Context) error {
	user := ctx.EffectiveUser
	msg := ctx.EffectiveMessage
	if wotoConfig.IsSudoUser(user.Id) {
		return ext.ContinueGroups
	}

	myStrs := ssg.Split(msg.Text)
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
