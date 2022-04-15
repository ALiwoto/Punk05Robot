package repostPlugin

import (
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters"
)

type channelPost struct {
	Filter   filters.Message
	Response handlers.Response
}
