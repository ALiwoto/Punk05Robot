package wotoConfig

import (
	"github.com/AnimeKaizoku/ssg/ssg"
)

func ParseConfig(filename string) (*BotConfig, error) {
	if ConfigSettings != nil {
		return ConfigSettings, nil
	}
	config := &BotConfig{}

	err := ssg.ParseConfig(config, filename)
	if err != nil {
		return nil, err
	}

	ConfigSettings = config
	channelIds = make(map[int64]bool)
	for _, id := range config.ChannelIds {
		channelIds[id] = true
	}

	return ConfigSettings, nil
}

func LoadConfig() (*BotConfig, error) {
	return ParseConfig("config.ini")
}

func GetBotToken() string {
	if ConfigSettings != nil {
		return ConfigSettings.BotToken
	}
	return ""
}

func DropUpdates() bool {
	return false
}

func GetCmdPrefixes() []rune {
	return []rune{'/', '!'}
}

func IsChannelAllowed(id int64) bool {
	return channelIds[id]
}
