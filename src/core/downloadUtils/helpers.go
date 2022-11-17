package downloadUtils

import (
	"encoding/json"
	"io"
	"net/http"
)

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
