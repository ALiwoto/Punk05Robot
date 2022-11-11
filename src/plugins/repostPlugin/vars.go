package repostPlugin

import (
	"github.com/AnimeKaizoku/ssg/ssg"
)

var (
	repeatCheckerMap      = _getRepeatCheckerMap()
	mediaGroupMessagesMap = _getMediaGroupMessagesMap()
)

var (
	jobsKeyGenerator = ssg.NewNumIdGenerator[uint64]()
)
