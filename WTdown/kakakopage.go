package WTdown

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/buger/jsonparser"
	"github.com/jiajunhuang/gotasks/pool"
)

type KakaoPage struct {
	TitleId string
	Cookies string // _kpawlt _kpawltea _kpawlst

	epis        []string
	EpisodeName []string
}

const (
	KAKAO_SINGLES_API  = "https://api2-page.kakao.com/api/v5/store/singles"
	KAKAO_IMG_API      = "https://api2-page.kakao.com/api/v1/inven/get_download_data/web"
	KAKAO_BASE_IMG_URL = "http://page-edge-jz.kakao.com/sdownload/resource/"
)

func (kp *KakaoPage) Download(start, stop, thread int, folder string) (int, error) {

	// _, err = os.Stat(WebtoonDownloadForm.folder + "\\" + WebtoonDownloadForm.id.Text())
	// if os.IsNotExist(err) {
	// 	err = os.MkdirAll(WebtoonDownloadForm.folder+"\\"+WebtoonDownloadForm.id.Text(), os.ModePerm)
	// 	if err != nil {
	// 		mw.openErrorMessBox("Error", err.Error())
	// 		return
	// 	}
	// }

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
					downloadFileSingle(KAKAO_BASE_IMG_URL+(*response)[anchor], folder+"/"+strconv.Itoa(episode+1)+"/"+strconv.Itoa(anchor+1)+".jpg")
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

func (kp *KakaoPage) getImgURL(productId string) (*[]string, error) {
	data := make(map[string]string)
	data["productId"] = productId

	header := make(map[string]string)
	header["Cookie"] = kp.Cookies
	header["user-agent"] = USER_AGENT
	header["Content-Type"] = "application/x-www-form-urlencoded"

	resp, err := requestWithCookieNBody(KAKAO_IMG_API, "POST", header, data)
	if err != nil {
		return nil, err
	}
	downloadData, err := ioutil.ReadAll(resp)
	if err != nil {
		return nil, err
	}

	output := make([]string, 0)

	type eachFiles struct {
		SecureUrl string `json:"secureUrl"`
	}

	var outputFiles []eachFiles
	files, err := jsonparser.GetUnsafeString(downloadData, "downloadData", "members", "files")
	json.Unmarshal([]byte(files), &outputFiles)
	for i := range outputFiles {
		output = append(output, outputFiles[i].SecureUrl)
	}
	return &output, nil
}

func (kp *KakaoPage) GetEpiData() error {
	kp.epis = make([]string, 0)
	kp.EpisodeName = make([]string, 0)

	c := 0
	for {
		data := make(map[string]string)
		data["seriesid"] = kp.TitleId
		data["page"] = strconv.Itoa(c)

		header := make(map[string]string)
		header["user-agent"] = USER_AGENT
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
			Id int `json:"id"`
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

// func (kp *KakaoPage) LoadEpisodesTitle() error {
// 	kp.episodeName = make([]string, 0)
// 	c := 0
// 	for {
// 		data := make(map[string]string)
// 		data["seriesid"] = kp.TitleId
// 		data["page"] = strconv.Itoa(c)

// 		header := make(map[string]string)
// 		header["Content-Type"] = "application/x-www-form-urlencoded"

// 		resp, err := requestWithCookieNBody(KAKAO_SINGLES_API, "POST", header, data)
// 		if err != nil {
// 			return err
// 		}

// 		downloadData, err := ioutil.ReadAll(resp)
// 		if err != nil {
// 			return err
// 		}

// 		type idFormat struct {
// 			Title string `json:"title"`
// 		}
// 		var outputFiles []idFormat
// 		singles, err := jsonparser.GetUnsafeString(downloadData, "singles")
// 		if err != nil {
// 			return err
// 		}
// 		json.Unmarshal([]byte(singles), &outputFiles)
// 		for i := range outputFiles {
// 			kp.episodeName = append(kp.episodeName, outputFiles[i].Title)
// 		}
// 		if len(outputFiles) <= 0 {
// 			break
// 		}
// 		c += 1
// 	}
// 	return nil
// }
