package downloadUtils

import (
	"strconv"
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
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
		return p.GetProperUrl()
	}

	return strings.ReplaceAll(p.GetProperUrl(), "_p0", "_p"+strconv.Itoa(page))
}

func (p *PixivInfoBody) GetProperUrl() string {
	if p.Width >= 720 || p.Height >= 720 {
		return p.Urls.Regular
	}

	return p.Urls.Original
}

//---------------------------------------------------------

func (m *MediaUrlInfo) GetMediaGroup(caption string) []gotgbot.InputMedia {
	// TODO: Add support for more than only photo types.
	return m.getPhotoArray(caption)
}

func (m *MediaUrlInfo) getPhotoArray(caption string) []gotgbot.InputMedia {
	var myArray []gotgbot.InputMedia
	captionDone := false
	if len(m.Urls) != 0 {
		for i, current := range m.Urls {
			if i == len(m.Urls)-1 && !captionDone {
				myArray = append(myArray, gotgbot.InputMediaPhoto{
					Media:   current,
					Caption: caption,
				})
				captionDone = true
				continue
			}

			myArray = append(myArray, gotgbot.InputMediaPhoto{
				Media: current,
			})
		}
	}

	if len(m.Files) != 0 {
		for i, current := range m.Files {
			if i == len(m.Files)-1 && !captionDone {
				myArray = append(myArray, gotgbot.InputMediaPhoto{
					Media:   current.Data,
					Caption: caption,
				})
				captionDone = true
				continue
			}

			myArray = append(myArray, gotgbot.InputMediaPhoto{
				Media: current.Data,
			})
		}
	}
	return myArray
}
