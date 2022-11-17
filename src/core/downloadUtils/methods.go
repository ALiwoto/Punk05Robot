package downloadUtils

import (
	"strconv"
	"strings"
)

// ---------------------------------------------------------
func (p *PixivInfoResponse) DownloadPage(page int) ([]byte, error) {
	if p.Body == nil {
		return nil, ErrPixivBodyNil
	}

	return p.Body.DownloadPage(page)
}

// ---------------------------------------------------------
func (p *PixivInfoBody) DownloadPage(page int) ([]byte, error) {
	if page >= p.PageCount {
		return nil, ErrPixivPageInvalid
	}

	return DownloadPixivIllustByLink(p.GetDirectUrlByPage(page))
}

func (p *PixivInfoBody) GetDirectUrlByPage(page int) string {
	if page == 0 {
		return p.Urls.Original
	}

	rawUrl := strings.TrimSuffix(p.Urls.Original, "_p0.jpg")
	return rawUrl + "_p" + strconv.Itoa(page) + ".jpg"
}

//---------------------------------------------------------
