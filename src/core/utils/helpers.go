package utils

import (
	"errors"
	"log"
	"net/url"
	"strconv"
	"strings"

	"github.com/ALiwoto/mdparser/mdparser"
	"github.com/AnimeKaizoku/Punk05Robot/src/core/logging"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func GetTwitterPhotoUrls(postLink string) (*MediaUrlInfo, error) {
	myUrl, err := url.Parse(postLink)
	if err != nil {
		return nil, err
	}

	myStrs := strings.Split(myUrl.Path, "/")
	postId := myStrs[len(myStrs)-1]
	if postId == "" {
		return nil, errors.New("empty post-id specified, make sure the post link is correct")
	}

	theTwit, err := TwitterClient.GetTweet(postId)
	if err != nil {
		return nil, err
	}

	profile, err := TwitterClient.GetProfile(theTwit.Username)
	if err != nil {
		return nil, err
	}

	return &MediaUrlInfo{
		Urls:  theTwit.Photos,
		Owner: profile.Name,
	}, nil
}

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
	_, err := m.Reply(b, md.ToString(), &gotgbot.SendMessageOpts{
		ParseMode: gotgbot.ParseModeMarkdownV2,
	})
	if err != nil {
		log.Println(err)
	}

	return nil
}

func SafeReply(b *gotgbot.Bot, ctx *ext.Context, output string) error {
	msg := ctx.EffectiveMessage
	if len(output) < 4096 {
		_, err := msg.Reply(b, output,
			&gotgbot.SendMessageOpts{ParseMode: gotgbot.ParseModeMarkdownV2})
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
