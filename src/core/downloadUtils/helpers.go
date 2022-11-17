package downloadUtils

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func IsSupportedUploadingUrl(value string) bool {
	value = strings.ReplaceAll(strings.ReplaceAll(strings.ToLower(value), "https://", ""), "http://", "")
	for current := range UrlUploaderHandlers {
		if strings.HasPrefix(value, current) {
			return true
		}
	}

	return false
}

func GetUrlUploaderHandler(theUrl string) MediaDownloadHandler {
	if theUrl == "" {
		return nil
	}

	theUrl = strings.ReplaceAll(strings.ReplaceAll(strings.ToLower(theUrl), "https://", ""), "http://", "")
	for key, handlerValue := range UrlUploaderHandlers {
		if strings.HasPrefix(theUrl, key) {
			return handlerValue
		}
	}

	return nil
}

func GetTwitterMediaInfo(postLink string) (*MediaUrlInfo, error) {
	myUrl, err := url.Parse(postLink)
	if err != nil {
		return nil, err
	}

	myStrs := strings.Split(myUrl.Path, "/")
	postId := myStrs[len(myStrs)-1]
	if postId == "" {
		return nil, errors.New("empty post-id specified, make sure the post link is correct")
	}

	theTwit, err := TwitterClient.GetTweet(postId)
	if err != nil {
		return nil, err
	}

	// profile, err := TwitterClient.GetProfile(theTwit.Username)
	// if err != nil {
	// return nil, err
	// }

	return &MediaUrlInfo{
		Urls:  theTwit.Photos,
		Owner: theTwit.Username,
	}, nil
}

func GetPixivMediaInfo(postLink string) (*MediaUrlInfo, error) {
	rawInfo, err := GetPixivIllustrateInfo(postLink)
	if err != nil {
		return nil, err
	} else if rawInfo.Error {
		return nil, errors.New(rawInfo.Message)
	}

	mediaInfo := new(MediaUrlInfo)
	mediaInfo.Owner = rawInfo.Body.UserAccount

	var currentPicData []byte
	for currentPage := 0; currentPage < rawInfo.Body.PageCount; currentPage++ {
		currentPicData, err = rawInfo.DownloadPage(currentPage)
		if err != nil || len(currentPicData) == 0 {
			continue
		}

		mediaInfo.Files = append(mediaInfo.Files, &MediaFile{
			Data: currentPicData,
		})
	}

	return mediaInfo, nil
}

//---------------------------------------------------------

func GetPixivIllustrateInfo(linkUrl string) (*PixivInfoResponse, error) {

	return nil, nil
}

func GetPixivIllustrateInfoById(illustId string) (*PixivInfoResponse, error) {
	req, err := http.NewRequest("GET", "https://pixiv.net/ajax/illust/"+illustId, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Referer", "https://pixiv.net/")
	req.Header.Set("Sec-Ch-Ua", "\"Google Chrome\";v=\"107\", \"Chromium\";v=\"107\", \"Not=A?Brand\";v=\"24\"")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("Sec-Ch-Ua-Platform", "\"Windows\"")
	req.Header.Set("Sec-Fetch-Dest", "image")
	req.Header.Set("Sec-Fetch-Mode", "no-cors")
	req.Header.Set("Sec-Fetch-Site", "cross-site")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	allData, _ := io.ReadAll(resp.Body)
	pixivResp := new(PixivInfoResponse)

	err = json.Unmarshal(allData, pixivResp)
	if err != nil {
		return nil, err
	}

	return pixivResp, nil
}

func DownloadPixivIllustByLink(linkUrl string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, linkUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authority", "i.pximg.net")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Referer", "https://www.pixiv.net/")
	req.Header.Set("Sec-Ch-Ua", "^^Google")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
