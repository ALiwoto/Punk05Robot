package database

import (
	"github.com/AnimeKaizoku/RepostingRobot/src/core/logging"
	wv "github.com/AnimeKaizoku/RepostingRobot/src/core/wotoValues"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var SESSION *gorm.DB

func StartDatabase() error {
	var err error
	var db *gorm.DB

	db, err = gorm.Open(
		sqlite.Open("reposting-robot.db"),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Error),
		},
	)
	if err != nil {
		return err
	}

	SESSION = db
	wv.DatabaseSession = db

	logging.Info("Database connected ")

	// create tables if they don't exist
	err = SESSION.AutoMigrate(
		modelChannelsSettings,
	)
	if err != nil {
		return err
	}

	err = LoadChannelsSettings()
	if err != nil {
		return err
	}

	logging.Info("Auto-migrated database schema")

	return nil
}

func LoadChannelsSettings() error {
	var allSettings []*wv.ChannelSettings

	lockDatabase()
	err := SESSION.Find(&allSettings).Error
	unlockDatabase()

	if err != nil {
		return err
	}

	if len(allSettings) != 0 {
		channelsSettings.AddPointerList(settingskeyGetter, allSettings...)
	}

	return nil
}

func settingskeyGetter(s *wv.ChannelSettings) int64 {
	return s.ChannelId
}

func lockDatabase() {
	mutex.Lock()
}

func unlockDatabase() {
	mutex.Unlock()
}
