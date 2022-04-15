package repostPlugin

import (
	"fmt"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func (m *channelPost) CheckUpdate(b *gotgbot.Bot, u *gotgbot.Update) bool {
	if u.Message != nil {
		return m.Filter == nil && m.Filter(u.Message)
	}

	// if no channel and message is channel message
	if u.ChannelPost != nil {
		return m.Filter != nil && m.Filter(u.ChannelPost)
	}

	return false
}

func (m *channelPost) HandleUpdate(b *gotgbot.Bot, ctx *ext.Context) error {
	return m.Response(b, ctx)
}

func (m *channelPost) Name() string {
	return fmt.Sprintf("message_%p", m.Response)
}
