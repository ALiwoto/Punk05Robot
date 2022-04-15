package wotoConfig

type BotConfig struct {
	BotToken   string  `section:"general" key:"bot_token"`
	ChannelIds []int64 `section:"general" key:"channel_ids"`
}
