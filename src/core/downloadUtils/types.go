package downloadUtils

// MediaUrlInfo contains information about a media obtained from a certain url in a certain
// media platform (such as twitter, pixiv, etc etc...)
type MediaUrlInfo struct {
	// Urls field is an array of direct url to the medias, owned by
	// same person/account that posted it.
	Urls []string

	// Files field specifies the files that has to be manually
	// uploaded to the tg by the bot.
	// In the case `Urls` is nil or empty, this field might be non-empty.
	Files []*MediaFile

	// Owner is the owner of the media(s) posted with that url.
	Owner string
}

type MediaFile struct {
	Data []byte
}

type PixivInfoResponse struct {
	Error   bool           `json:"error"`
	Message string         `json:"message"`
	Body    *PixivInfoBody `json:"body"`
}

type PixivInfoBody struct {
	IllustId      string          `json:"illustId"`
	IllustTitle   string          `json:"illustTitle"`
	UserId        string          `json:"userId"`
	UserName      string          `json:"userName"`
	UserAccount   string          `json:"userAccount"`
	Urls          *PixivUrls      `json:"urls"`
	IllustComment string          `json:"illustComment"`
	Id            string          `json:"id"`
	Title         string          `json:"title"`
	Description   string          `json:"description"`
	IllustType    PixivIllustType `json:"illustType"`
	Sl            int             `json:"sl"`
	Width         int             `json:"width"`
	Height        int             `json:"height"`
	PageCount     int             `json:"pageCount"`
	ViewCount     int             `json:"viewCount"`
	IsUnlisted    bool            `json:"isUnlisted"`
	AiType        int             `json:"aiType"`
}

type PixivIllustType int

type PixivUrls struct {
	Mini     string `json:"mini"`
	Thumb    string `json:"thumb"`
	Small    string `json:"small"`
	Regular  string `json:"regular"`
	Original string `json:"original"`
}

type MediaDownloadHandler func(inputUrl string) (*MediaUrlInfo, error)
