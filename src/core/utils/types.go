package utils

// MediaUrlInfo contains information about a media obtained from a certain url in a certain
// media platform (such as twitter, pixiv, etc etc...)
type MediaUrlInfo struct {
	// Urls field is an array of direct url to the medias, owned by
	// same person/account that posted it.
	Urls []string

	// Owner is the owner of the media(s) posted with that url.
	Owner string
}
