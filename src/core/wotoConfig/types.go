package wotoConfig

type BotConfig struct {
	BotToken  string  `section:"general" key:"bot_token"`
	SudoUsers []int64 `section:"general" key:"sudo_users"`
}
