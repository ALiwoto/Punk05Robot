package downloadUtils

import (
	"errors"

	twitterLib "github.com/n0madic/twitter-scraper"
)

var TwitterClient = twitterLib.New()

var UrlUploaderHandlers = map[string]MediaDownloadHandler{
	"twitter.com/": GetTwitterMediaInfo,
	"pixiv.net/":   GetPixivMediaInfo,
}

// pixiv error vars
var (
	ErrPixivBodyNil     = errors.New("pixiv body is nil")
	ErrPixivPageInvalid = errors.New("invalid page specified, page number should start from 0")
	ErrPixivUrlInvalid  = errors.New("the provided url is invalid")
)
