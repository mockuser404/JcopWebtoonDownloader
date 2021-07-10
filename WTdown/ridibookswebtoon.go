package WTdown

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/jiajunhuang/gotasks/pool"
)

type RidiWT struct {
	TitleId string
	Cookies string // _kpawlt _kpawltea _kpawlst

	epis        []string
	EpisodeName []string
}

const (
	RIDIWT_HOME_URL = "https://ridibooks.com/books/"
	RIDIWT_IMG_API  = "https://view.ridibooks.com/generate/"

// 	KAKAO_SINGLES_API  = "https://api2-page.kakao.com/api/v5/store/singles"
// 	KAKAO_IMG_API      = "https://api2-page.kakao.com/api/v1/inven/get_download_data/web"
// 	KAKAO_BASE_IMG_URL = "http://page-edge-jz.kakao.com/sdownload/resource/"
)

func (kp *RidiWT) Download(start, stop, thread int, folder string) (int, error) {

	if len(kp.epis) <= 0 {
		return 2, errors.New("Wrong Webtoon Id")
	}

	for episode := start; episode <= stop; episode++ {
		if episode > len(kp.epis) {
			return 2, errors.New("Can't find Epi" + strconv.Itoa(episode+1))
		}
		response, err := kp.getImgURL(kp.epis[episode])
		if err != nil {
			return 1, err
		}
		if len(*response) <= 0 {
			return 2, errors.New("Epi" + strconv.Itoa(episode+1) + " - Can't find images")
		}

		err = os.MkdirAll(folder+"/"+strconv.Itoa(episode+1), os.ModePerm)
		if err != nil {
			return 1, err
		}
		// errchan := make(chan error, len(response))
		gopool := pool.NewGoPool(pool.WithMaxLimit(thread))

		for anchor := range *response {
			func(anchor int) {
				gopool.Submit(func() {
					downloadFileSingle((*response)[anchor], folder+"/"+strconv.Itoa(episode+1)+"/"+strconv.Itoa(anchor+1)+".jpg")
				})
			}(anchor)
			// go downloadFile(KAKAO_BASE_IMG_URL+response[anchor], folder+"/"+strconv.Itoa(episode)+"/"+strconv.Itoa(anchor+1)+".jpg", errchan)
		}

		gopool.Wait()
		// for i := 0; i < len(response); i++ {
		// 	tmp_err := <-errchan
		// 	if tmp_err != nil {
		// 		return 2, err
		// 	}
		// }

		makeHTML(episode+1, len(*response), kp.TitleId, folder+"/"+strconv.Itoa(episode+1)+"/"+strconv.Itoa(episode+1)+".html")
	}
	return 0, nil
}

func (rd *RidiWT) getImgURL(productId string) (*[]string, error) {
	header := make(map[string]string)
	header["Cookie"] = rd.Cookies

	resp, err := requestWithCookieNBody(RIDIWT_IMG_API+productId, "GET", header, make(map[string]string))
	if err != nil {
		return nil, err
	}
	downloadData, err := ioutil.ReadAll(resp)
	if err != nil {
		return nil, err
	}

	output := make([]string, 0)

	type eachFiles struct {
		Pages []map[string]string `json:"pages"`
	}

	var outputFiles eachFiles
	json.Unmarshal(downloadData, &outputFiles)
	for i := range outputFiles.Pages {
		output = append(output, outputFiles.Pages[i]["src"])
	}
	return &output, nil
}

func (rd *RidiWT) GetEpiData() error {
	rd.epis = make([]string, 0)
	rd.EpisodeName = make([]string, 0)

	header := make(map[string]string)
	header["Cookie"] = rd.Cookies

	resp, err := requestWithCookieNBody(RIDIWT_HOME_URL+rd.TitleId, "GET", header, make(map[string]string))
	if err != nil {
		return err
	}
	doc, err := goquery.NewDocumentFromReader(resp)
	if err != nil {
		return err
	}
	doc.Find(".js_book_checkbox_input").Each(func(_ int, s *goquery.Selection) {
		id, state := s.Attr("value")
		if state {
			rd.epis = append(rd.epis, id)
		}
	})
	doc.Find(".js_book_title").Each(func(_ int, s *goquery.Selection) {
		title := s.Text()
		if title != "" {
			rd.EpisodeName = append(rd.EpisodeName, title)
		}
	})
	return nil
}
