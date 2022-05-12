package main

import (
	"log"

	"github.com/AnimeKaizoku/RepostingRobot/src/core/logging"
	"github.com/AnimeKaizoku/RepostingRobot/src/core/wotoConfig"
	"github.com/AnimeKaizoku/RepostingRobot/src/database"
	"github.com/AnimeKaizoku/RepostingRobot/src/plugins"
)

func main() {
	_, err := wotoConfig.LoadConfig()
	if err != nil {
		log.Fatal("Error parsing config file", err)
	}

	f := logging.LoadLogger()
	if f != nil {
		defer f()
	}

	err = database.StartDatabase()
	if err != nil {
		logging.Fatal("Error starting database", err)
	}

	err = plugins.StartTelegramBot()
	if err != nil {
		logging.Fatal("Failed to start the bot bot: ", err)
	}
}
