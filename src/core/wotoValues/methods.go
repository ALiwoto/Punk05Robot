package wotoValues

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
)

func (j *PendingJob) ShouldBeHandled() bool {
	return j.TimeDistance == 0 || time.Since(j.RegisteredTime) >= j.TimeDistance
}

func (j *PendingJob) GenerateButtons() *gotgbot.InlineKeyboardMarkup {
	if j.ButtonGenerator == nil {
		return nil
	}

	return j.ButtonGenerator(j)
}

func (j *PendingJob) GetPostCaption() string {
	if j.CaptionGetter == nil {
		return ""
	}

	return j.CaptionGetter(j)
}

//---------------------------------------------------------

func (i ButtonsUniqueId) IsEmpty() bool {
	return i == ""
}
