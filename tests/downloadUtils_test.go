package tests

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/AnimeKaizoku/Punk05Robot/src/core/downloadUtils"
	twitterscraper "github.com/n0madic/twitter-scraper"
)

func TestGetTwitterPic(t *testing.T) {
	//utils.GetTwitterPhotoUrls("https://twitter.com/gabiran_/status/1590689796812582913?server=19")
	//utils.GetTwitterPhotoUrls("https://twitter.com/HitenKei/status/1591051073133113346?s=20&t=6jqNWiXWFRO3vhwwNiZExg")
	_, err := downloadUtils.GetTwitterMediaInfo("https://twitter.com/haori_crescendo/status/1586563553414172672?s=20&t=6jqNWiXWFRO3vhwwNiZExg")
	if err != nil {
		t.Error(err)
		return
	}

	s := twitterscraper.New()
	twitt, err := s.GetTweet("1590689796812582913")
	if err != nil {
		print(err)
		return
	}

	print(twitt.Photos)
	//req, err := http.NewRequest("GET", "https://pbs.twimg.com/media/FhNDI2_acAAcJ0W?format=jpg&name=medium", nil)
	req, err := http.NewRequest("GET", "https://twitter.com/gabiran_/status/1590689796812582913", nil)
	if err != nil {
		// handle err
		print(err)
	}
	//req.Header.Set("Authority", "pbs.twimg.com")
	req.Header.Set("Accept", "image/avif,image/webp,image/apng,image/svg+xml,image/*,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Referer", "https://twitter.com/")
	req.Header.Set("Sec-Ch-Ua", "\"Google Chrome\";v=\"107\", \"Chromium\";v=\"107\", \"Not=A?Brand\";v=\"24\"")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("Sec-Ch-Ua-Platform", "\"Windows\"")
	req.Header.Set("Sec-Fetch-Dest", "image")
	req.Header.Set("Sec-Fetch-Mode", "no-cors")
	req.Header.Set("Sec-Fetch-Site", "cross-site")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
		print(err)
	}
	defer resp.Body.Close()

	b, _ := io.ReadAll(resp.Body)
	myStr := string(b)
	doesContain := strings.Contains(myStr, "FhNDI2_acAAcJ0W")
	print(doesContain)
	print(myStr)
}

func TestGetPixivInfo(t *testing.T) {
	info, err := downloadUtils.GetPixivIllustrateInfo("https://www.pixiv.net/member_illust.php?mode=medium&illust_id=102825614")
	if err != nil {
		t.Error("when trying to use GetPixivIllustrateInfo:", err)
		return
	}
	if info.Error {
		t.Error("when trying to get illust info:", info.Message)
		return
	}

	info, err = downloadUtils.GetPixivIllustrateInfo("https://www.pixiv.net/artworks/102825614")
	if err != nil {
		t.Error("when trying to use GetPixivIllustrateInfoById:", err)
		return
	}

	if info.Error {
		t.Error("when trying to get illust info:", info.Message)
		return
	}

	info, err = downloadUtils.GetPixivIllustrateInfoById("101670857")
	if err != nil {
		t.Error("when trying to use GetPixivIllustrateInfoById:", err)
		return
	}

	if info.Error {
		t.Error("when trying to get illust info:", info.Message)
		return
	}
}

func TestPixivDownload01(t *testing.T) {
	info, err := downloadUtils.GetPixivIllustrateInfoById("101670857")
	if err != nil {
		t.Error("when trying to use GetPixivIllustrateInfoById:", err)
		return
	}

	if info.Error {
		t.Error("when trying to get illust info:", info.Message)
		return
	}

	picData, err := info.DownloadPage(0)
	if err != nil {
		t.Error("when trying to use DownloadPage:", err)
		return
	}

	if len(picData) == 0 {
		t.Error("got empty picData after using DownloadPage method")
		return
	}
}
