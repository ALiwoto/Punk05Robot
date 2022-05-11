package wotoValues

import "time"

func (j *PendingJob) ShouldBeHandled() bool {
	return j.TimeDistance == 0 || time.Since(j.RegisteredTime) >= j.TimeDistance
}
