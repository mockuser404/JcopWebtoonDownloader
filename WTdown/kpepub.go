package WTdown

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/buger/jsonparser"
)

type KPepub struct {
	TitleId string
	Cookies string // _kpawlt _kpawltea _kpawlst

	epis        []string
	EpisodeName []string
}

const (
	// KAKAO_SINGLES_API  = "https://api2-page.kakao.com/api/v5/store/singles"
	// KAKAO_IMG_API      = "https://api2-page.kakao.com/api/v1/inven/get_download_data/web"
	// KAKAO_BASE_IMG_URL = "http://page-edge-jz.kakao.com/sdownload/resource/"
	EPUB_VIEWER_API = "https://page.kakao.com/viewer?productId="
	EPUB_TEXT_API   = "https://dn-img-page.kakao.com/download/resource?kid="
)

func (kp *KPepub) Download(start, stop int, folder string) (int, error) {

	if len(kp.epis) <= 0 {
		return 2, errors.New("Wrong Webtoon Id")
	}

	for episode := start; episode <= stop; episode++ {
		if episode > len(kp.epis) {
			return 2, errors.New("Can't find Epi" + strconv.Itoa(episode+1))
		}
		response, err := kp.getTextURL(kp.epis[episode])
		if err != nil {
			return 1, err
		}
		if response == "" {
			return 2, errors.New("Epi " + strconv.Itoa(episode+1) + " - Can't find text")
		}

		// err = os.MkdirAll(folder+"/"+strconv.Itoa(episode+1), os.ModePerm)
		// if err != nil {
		// 	return 1, err
		// }
		err = kp.downloadText(response, folder+"/"+"/"+strconv.Itoa(episode+1)+".html")
		if err != nil {
			return 1, err
		}
		// downloadFileSingle(EPUB_TEXT_API+response, folder+"/"+strconv.Itoa(episode+1)+"/"+strconv.Itoa(anchor+1)+".jpg")

		// makeHTML(episode+1, len(response), kp.TitleId, folder+"/"+strconv.Itoa(episode+1)+"/"+strconv.Itoa(episode+1)+".html")
	}
	return 0, nil
}

func (kp *KPepub) getTextURL(productId string) (string, error) {
	// data := make(map[string]string)
	// data["productId"] = productId

	header := make(map[string]string)
	header["Cookie"] = kp.Cookies
	header["Content-Type"] = "application/x-www-form-urlencoded"

	resp, err := requestWithCookieNBody(EPUB_VIEWER_API+productId, "GET", header, make(map[string]string))
	if err != nil {
		return "", err
	}

	doc, err := goquery.NewDocumentFromReader(resp)
	if err != nil {
		return "", err
	}
	downloadData := doc.Find("#__NEXT_DATA__").Text()

	files, err := jsonparser.GetUnsafeString([]byte(downloadData), "props", "initialState", "product", "productMap", productId, "singleForMeta", "epubViewerId")

	return string(files), nil
}

func (kp *KPepub) downloadText(epubViewerId, outFile string) error {
	req, err := http.NewRequest("GET", EPUB_TEXT_API+epubViewerId, nil)
	if err != nil {
		return err
	}

	header := make(map[string]string)
	header["Cookie"] = kp.Cookies
	header["Content-Type"] = "application/x-www-form-urlencoded"
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	start := strings.Index(string(content), "onMainJsonLoaded(") + 17
	stop := strings.LastIndex(string(content[start:]), ");") - 1
	// log.Println(string(content[start : start+stop]))
	outcontent, err := jsonparser.GetString(content[start:start+stop], "body")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(outFile, []byte(`<head><meta charset="utf-8"></head>`+outcontent), 0)
	if err != nil {
		return err
	}

	return nil
}

func (kp *KPepub) GetEpiData() error {
	kp.epis = make([]string, 0)
	kp.EpisodeName = make([]string, 0)

	c := 0
	for {
		data := make(map[string]string)
		data["seriesid"] = kp.TitleId
		data["page"] = strconv.Itoa(c)

		header := make(map[string]string)
		header["Content-Type"] = "application/x-www-form-urlencoded"

		resp, err := requestWithCookieNBody(KAKAO_SINGLES_API, "POST", header, data)
		if err != nil {
			return err
		}

		downloadData, err := ioutil.ReadAll(resp)
		if err != nil {
			return err
		}
		// log.Println(string(downloadData))
		type idFormat struct {
			Id    int    `json:"id"`
			Title string `json:"title"`
		}
		var outputFiles []idFormat
		singles, err := jsonparser.GetUnsafeString(downloadData, "singles")
		if err != nil {
			return err
		}
		json.Unmarshal([]byte(singles), &outputFiles)
		// log.Println(outputFiles)
		for i := range outputFiles {
			kp.epis = append(kp.epis, strconv.Itoa(outputFiles[i].Id))
			kp.EpisodeName = append(kp.EpisodeName, outputFiles[i].Title)
		}
		if len(outputFiles) <= 0 {
			break
		}
		c += 1
	}
	// log.Println(kp.epis)
	return nil
}
