package database

import (
	"github.com/AnimeKaizoku/RepostingRobot/src/core/logging"
	"github.com/AnimeKaizoku/RepostingRobot/src/core/wotoConfig"
	wv "github.com/AnimeKaizoku/RepostingRobot/src/core/wotoValues"
	"github.com/AnimeKaizoku/ssg/ssg"
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
		modelChannelAccessElement,
	)
	if err != nil {
		return err
	}

	err = LoadChannelsSettings()
	if err != nil {
		return err
	}

	err = LoadChannelAccessElements()
	if err != nil {
		return err
	}

	logging.Info("Auto-migrated database schema")

	// additional registering stuff
	RegisterNewChannels(wotoConfig.GetChannelIDs())

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

func LoadChannelAccessElements() error {
	var allElements []*wv.ChannelAccessElement

	lockDatabase()
	err := SESSION.Find(&allElements).Error
	unlockDatabase()

	if err != nil {
		return err
	}

	for _, current := range allElements {
		cacheAccessElement(current)
	}

	return nil
}

func GetChannelSettings(id int64) *wv.ChannelSettings {
	return channelsSettings.Get(id)
}

func GetUserAllAccess(userId int64) []*wv.ChannelAccessElement {
	return userAccessChannels.GetValue(userId)
}

// RegisterNewChannels registers new default settings for the given
// channel-ids (if and only if the channel is not already registered.)
func RegisterNewChannels(ids []int64) {
	if len(ids) == 0 {
		return
	}

	var settings *wv.ChannelSettings
	for _, current := range ids {
		settings = GetChannelSettings(current)
		if settings != nil {
			// settings already exists, skip
			continue
		}

		SaveChannelSettings(&wv.ChannelSettings{
			ChannelId: current,
		}, true)
	}
}

func SaveChannelSettings(settings *wv.ChannelSettings, cache bool) {
	lockDatabase()
	tx := SESSION.Begin()
	tx.Save(settings)
	tx.Commit()
	unlockDatabase()

	if cache {
		if settings.AccessMap == nil {
			settings.AccessMap = ssg.NewSafeMap[int64, wv.ChannelAccessElement]()
		}

		channelsSettings.Add(settings.ChannelId, settings)
	}
}

func SaveAccessElement(element *wv.ChannelAccessElement, cache bool) {
	lockDatabase()
	tx := SESSION.Begin()
	tx.Save(element)
	tx.Commit()
	unlockDatabase()

	if cache {
		cacheAccessElement(element)
	}
}

func cacheAccessElement(element *wv.ChannelAccessElement) {
	settings := GetChannelSettings(element.ChannelId)
	if settings == nil {
		return
	}

	settings.AccessMap.Add(element.UserId, element)
	userAllAccess := userAccessChannels.GetValue(element.UserId)
	userAllAccess = append(userAllAccess, element)
	userAccessChannels.Set(element.UserId, userAllAccess)
}

func IsChannelRegistered(id int64) bool {
	return channelsSettings.Exists(id)
}

func settingskeyGetter(s *wv.ChannelSettings) int64 {
	if s.AccessMap == nil {
		s.AccessMap = ssg.NewSafeMap[int64, wv.ChannelAccessElement]()
	}

	return s.ChannelId
}

func lockDatabase() {
	mutex.Lock()
}

func unlockDatabase() {
	mutex.Unlock()
}
