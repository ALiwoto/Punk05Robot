package database

import (
	"sync"

	wv "github.com/AnimeKaizoku/RepostingRobot/src/core/wotoValues"
	"github.com/AnimeKaizoku/ssg/ssg"
)

var (
	mutex                 = &sync.Mutex{}
	modelChannelsSettings = &wv.ChannelSettings{}
)

var (
	channelsSettings = ssg.NewSafeMap[int64, wv.ChannelSettings]()
)
