package utils

import twitterLib "github.com/n0madic/twitter-scraper"

var TwitterClient = twitterLib.New()

var SupportedUploadingUrl = map[string]bool{
	"twitter.com/": true,
}
