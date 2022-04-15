package wotoConfig

type BotConfig struct {
	BotToken string  `section:"general" key:"bot_token"`
	Owners   []int64 `section:"general" key:"owners"`
}
