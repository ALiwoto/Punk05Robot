package plugins

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/AnimeKaizoku/RepostingRobot/src/core/logging"
	"github.com/AnimeKaizoku/RepostingRobot/src/core/wotoConfig"
	wv "github.com/AnimeKaizoku/RepostingRobot/src/core/wotoValues"
	"github.com/AnimeKaizoku/RepostingRobot/src/database"
	"github.com/AnimeKaizoku/ssg/ssg"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func StartTelegramBot() error {
	token := wotoConfig.GetBotToken()
	if len(token) == 0 {
		return errors.New("bot token is empty")
	}

	b, err := gotgbot.NewBot(token, &gotgbot.BotOpts{
		Client:      http.Client{},
		GetTimeout:  gotgbot.DefaultGetTimeout,
		PostTimeout: gotgbot.DefaultPostTimeout,
	})
	if err != nil {
		return err
	}

	utmp := ext.NewUpdater(nil)
	updater := &utmp
	err = updater.StartPolling(b, &ext.PollingOpts{
		DropPendingUpdates: wotoConfig.DropUpdates(),
	})
	if err != nil {
		return err
	}

	logging.Info(fmt.Sprintf("%s has started | ID: %d", b.Username, b.Id))

	wv.HelperBot = b
	wv.BotUpdater = updater

	LoadAllHandlers(updater.Dispatcher, wotoConfig.GetCmdPrefixes())

	return proccessJobs()
}

func proccessJobs() error {
	var handledJobs int
	var longHandledJobs int
	var err error
	for {
		handledJobs = 0
		time.Sleep(time.Second)

		if wv.PendingJobs == nil || wv.HaltJobs {
			return nil
		}

		if wv.PendingJobs.IsEmpty() || wv.PauseJobs {
			continue
		}

		wv.PendingJobs.ForEach(func(key uint64, job *wv.PendingJob) bool {
			if !job.ShouldBeHandled() || database.IsTmpIgnoring(job.Ctx.EffectiveChat.Id) {
				return false
			}

			if longHandledJobs > 2*wv.MaxJobsPerSecond {
				time.Sleep(30 * time.Second)
				longHandledJobs = 0
				return false
			}

			if handledJobs > wv.MaxJobsPerSecond {
				return false
			}

			if job.Handler == nil {
				return true
			}

			err = job.Handler(job)
			if err != nil {
				errStr := err.Error()
				myStrs := strings.Split(errStr, "Too Many Requests: retry after ")
				if len(myStrs) >= 2 {
					theSeconds := ssg.ToInt64(myStrs[1])
					if theSeconds > 0 {
						longHandledJobs = 0
						handledJobs = 0
						time.Sleep(time.Duration(theSeconds) * time.Second)
						return false
					}
				}
				logging.Errorf("Error while handling job %s: %v", key, err)
			}

			handledJobs++
			longHandledJobs++
			return true
		})
	}
}
