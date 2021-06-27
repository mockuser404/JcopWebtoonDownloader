package WTdown

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/buger/jsonparser"
	"github.com/jiajunhuang/gotasks/pool"
)

type DaumWebtoon struct {
	TitleId string
	Cookies string // _kpawlt _kpawltea _kpawlst

	epis        []string
	EpisodeName []string
}

const (
	DAUM_SINGLES_API  = "http://webtoon.daum.net/data/pc/webtoon/view/"
	DAUM_IMG_API      = "http://webtoon.daum.net/data/pc/webtoon/viewer_images/"
	DAUM_BASE_IMG_URL = "http://page-edge-jz.kakao.com/sdownload/resource/"
)

func (dw *DaumWebtoon) Download(start, stop, thread int, folder string) (int, error) {

	if len(dw.epis) <= 0 {
		return 1, errors.New("Wrong Webtoon Id")
	}
	for episode := start; episode <= stop; episode++ {
		if episode > len(dw.epis) {
			return 2, errors.New("Can't find Epi " + strconv.Itoa(episode+1))
		}
		response, err := dw.getImgURL(dw.epis[episode])
		if err != nil {
			return 1, err
		}
		if len(response) <= 0 {
			return 2, errors.New("Epi" + strconv.Itoa(episode+1) + " - Can't find images")
		}

		err = os.MkdirAll(folder+"/"+strconv.Itoa(episode+1), os.ModePerm)
		if err != nil {
			return 1, err
		}
		// errchan := make(chan error, len(response))
		gopool := pool.NewGoPool(pool.WithMaxLimit(thread))

		for anchor := range response {
			func(anchor int) {
				gopool.Submit(func() {
					downloadFileSingle(response[anchor], folder+"/"+strconv.Itoa(episode+1)+"/"+strconv.Itoa(anchor+1)+".jpg")
				})
			}(anchor)
			// go downloadFile(response[anchor], folder+"/"+strconv.Itoa(episode)+"/"+strconv.Itoa(anchor+1)+".jpg", errchan)
		}
		gopool.Wait()
		// for i := 0; i < len(response); i++ {
		// 	err = <-errchan
		// 	if err != nil {
		// 		return 1, err
		// 	}
		// }
		makeHTML(episode+1, len(response), dw.TitleId, folder+"/"+strconv.Itoa(episode+1)+"/"+strconv.Itoa(episode+1)+".html")
	}
	return 0, nil
}

func (dw *DaumWebtoon) getImgURL(productId string) ([]string, error) {
	resp, err := http.Get(DAUM_IMG_API + productId)
	if err != nil {
		return nil, err
	}
	downloadData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	output := make([]string, 0)

	type eachFiles struct {
		Url string `json:"url"`
	}

	var outputFiles []eachFiles
	files, err := jsonparser.GetUnsafeString(downloadData, "data")
	json.Unmarshal([]byte(files), &outputFiles)
	for i := range outputFiles {
		output = append(output, outputFiles[i].Url)
	}
	return output, nil
}

func (dw *DaumWebtoon) GetEpiData() error {
	dw.epis = make([]string, 0)
	dw.EpisodeName = make([]string, 0)

	resp, err := http.Get(DAUM_SINGLES_API + dw.TitleId)
	if err != nil {
		return err
	}
	downloadData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	type eachFiles struct {
		Id int `json:"id"`
		Title string `json:"title"`
	}

	var outputFiles []eachFiles
	files, err := jsonparser.GetUnsafeString(downloadData, "data", "webtoon", "webtoonEpisodes")
	if err != nil {
		return nil
	}
	json.Unmarshal([]byte(files), &outputFiles)

	for i := range outputFiles {
		dw.epis = append(dw.epis, strconv.Itoa(outputFiles[i].Id))
		dw.EpisodeName = append(dw.EpisodeName, outputFiles[i].Title)
	}

	order, err := jsonparser.GetString(downloadData, "data", "webtoon", "sort")
	if err != nil {
		return nil
	}
	if order != "asc" {
		reverseStrArray(&dw.epis)
		reverseStrArray(&dw.EpisodeName)
	}
	return nil
}

// func (dw *DaumWebtoon) LoadEpisodesTitle() error {
// 	resp, err := http.Get(DAUM_SINGLES_API + dw.TitleId)
// 	if err != nil {
// 		return err
// 	}
// 	downloadData, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return err
// 	}

// 	type eachFiles struct {
// 		Title string `json:"title"`
// 	}

// 	var outputFiles []eachFiles
// 	files, err := jsonparser.GetUnsafeString(downloadData, "data", "webtoon", "webtoonEpisodes")
// 	json.Unmarshal([]byte(files), &outputFiles)

// 	dw.EpisodeName = make([]string, 0)
// 	for i := range outputFiles {
// 		dw.EpisodeName = append(dw.EpisodeName, outputFiles[i].Title)
// 	}

// 	order, err := jsonparser.GetString(downloadData, "data", "webtoon", "sort")
// 	if err != nil {
// 		return nil
// 	}
// 	if order != "asc" {
// 		reverseStrArray(&dw.EpisodeName)
// 	}
// 	return nil
// }
