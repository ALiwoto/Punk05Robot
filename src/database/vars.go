package database

import (
	"sync"

	wv "github.com/AnimeKaizoku/Punk05Robot/src/core/wotoValues"
	"github.com/AnimeKaizoku/ssg/ssg"
)

var (
	mutex                     = &sync.Mutex{}
	modelChannelsSettings     = &wv.ChannelSettings{}
	modelChannelAccessElement = &wv.ChannelAccessElement{}
)

var (
	channelsSettings   = ssg.NewSafeMap[int64, wv.ChannelSettings]()
	userAccessChannels = ssg.NewSafeMap[int64, []*wv.ChannelAccessElement]()
)
