package downloadUtils

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
