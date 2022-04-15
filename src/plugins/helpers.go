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
	var jobs map[string]wv.PendingJob
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

		jobs = wv.PendingJobs.ToNormalMap()
		if len(jobs) == 0 {
			continue
		}

		for key, job := range jobs {
			if longHandledJobs > 2*wv.MaxJobsPerSecond {
				time.Sleep(30 * time.Second)
				longHandledJobs = 0
				break
			}

			if handledJobs > wv.MaxJobsPerSecond {
				break
			}

			if job.Handler == nil {
				wv.PendingJobs.Delete(key)
				continue
			}

			err = job.Handler(&job)
			if err != nil {
				errStr := err.Error()
				myStrs := strings.Split(errStr, "Too Many Requests: retry after ")
				if len(myStrs) >= 2 {
					theSeconds := ssg.ToInt64(myStrs[1])
					if theSeconds > 0 {
						longHandledJobs = 0
						handledJobs = 0
						time.Sleep(time.Duration(theSeconds) * time.Second)
						continue
					}
				}
				logging.Errorf("Error while handling job %s: %v", key, err)
			}

			wv.PendingJobs.Delete(key)
			handledJobs++
			longHandledJobs++
		}

	}
}
