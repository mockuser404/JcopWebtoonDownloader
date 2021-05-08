package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/buger/jsonparser"
	"github.com/lxn/walk"
)

const naverBaseURL string = "https://comic.naver.com/webtoon/detail.nhn"
const lezhinBaseURL string = "https://www.lezhin.com/api/v2/inventory_groups/comic_viewer_k"
const UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/50.0.2661.102 Safari/537.36"

var err error

func webtoonDownload() {
	if WebtoonDownloadForm.wtype.CurrentIndex() == 0 {
		WebtoonDownloadForm.NaverComicDownload(WebtoonDownloadForm.folder + "\\" + WebtoonDownloadForm.id.Text())
	} else if WebtoonDownloadForm.wtype.CurrentIndex() == 1 {
		WebtoonDownloadForm.KakaoPageDownload(WebtoonDownloadForm.folder + "\\" + WebtoonDownloadForm.id.Text())
	} else if WebtoonDownloadForm.wtype.CurrentIndex() == 2 {
		WebtoonDownloadForm.daumWebtoonDownload(WebtoonDownloadForm.folder + "\\" + WebtoonDownloadForm.id.Text())
	} else if WebtoonDownloadForm.wtype.CurrentIndex() == 3 {
		WebtoonDownloadForm.lezhinComicsDownload(WebtoonDownloadForm.folder + "\\" + WebtoonDownloadForm.id.Text())
	} else {
		mw.openWarningMessBox("Warning", "Please select type")
	}
}

type DownloadForm struct {
	wtype       *walk.ComboBox
	id          *walk.LineEdit
	start       *walk.LineEdit
	stop        *walk.LineEdit
	deviceId    string
	folder      string
	cookie      cookieData
	accesstoken string
}

func (df *DownloadForm) NaverComicDownload(folder string) {
	var processing float32 = 0.0
	start, err := strconv.Atoi(WebtoonDownloadForm.start.Text())
	if err != nil {
		mw.openWarningMessBox("Warning", "Invalid episode")
		return
	}
	stop, err := strconv.Atoi(WebtoonDownloadForm.stop.Text())
	if err != nil {
		mw.openWarningMessBox("Warning", "Invalid episode")
		return
	}

	_, err = os.Stat(WebtoonDownloadForm.folder + "\\" + WebtoonDownloadForm.id.Text())
	if os.IsNotExist(err) {
		err = os.MkdirAll(WebtoonDownloadForm.folder+"\\"+WebtoonDownloadForm.id.Text(), os.ModePerm)
		if err != nil {
			mw.openErrorMessBox("Error", err.Error())
			return
		}
	}

	c := 0
	buttonLog.SetEnabled(false)
	buttonLog.SetText("Downloading... " + fmt.Sprintf("%.2f", processing) + "%")
	for episode := start; episode <= stop; episode++ {
		titleID := df.id.Text()

		var dataURL string

		resp, err := requestWithCookie(naverBaseURL+"?titleId="+titleID+"&no="+strconv.Itoa(episode), "GET", df.cookie.naverComicData)
		if err != nil {
			mw.openErrorMessBox("Error", err.Error())
			resetDownloadButton()
			return
		}

		doc, err := goquery.NewDocumentFromReader(resp)
		if err != nil {
			mw.openErrorMessBox("Error", err.Error())
			resetDownloadButton()
			return
		}

		NumImg := doc.Find(".wt_viewer").Find("img").Length()

		if NumImg <= 0 {
			mw.openWarningMessBox("Warning", "Can't find Images")
			resetDownloadButton()
			return
		}

		err = os.MkdirAll(folder+"/"+strconv.Itoa(episode), os.ModePerm)
		if err != nil {
			mw.openErrorMessBox("Error", err.Error())
			resetDownloadButton()
			return
		}

		errchan := make(chan error, NumImg)
		doc.Find(".wt_viewer").Find("img").Each(func(j int, s *goquery.Selection) {
			dataURL, _ = s.Attr("src")
			go downloadFile(string(dataURL), folder+"/"+strconv.Itoa(episode)+"/"+strconv.Itoa(j+1)+".jpg", errchan)
		})
		for i := 0; i < NumImg; i++ {
			err = <-errchan
			if err != nil {
				mw.openErrorMessBox("Error", err.Error())
				resetDownloadButton()
				return
			}
		}

		content := "<html><head><meta charset='UTF-8'><meta name='viewport' content='width=device-width, initial-scale=1.0'><style>body, html{margin: 0;border: 0;padding: 0;}@media only screen and (max-width: 700px) {img {width: 100%;}}</style><title>Episode " + strconv.Itoa(episode) + " (" + titleID + ")</title></head><body><center>"
		for l := 1; l <= NumImg; l++ {
			content += "<img src='"
			content += strconv.Itoa(l)
			content += ".jpg'><br>"
		}
		content += "</body></center></html>"

		err = ioutil.WriteFile(folder+"/"+strconv.Itoa(episode)+"/"+strconv.Itoa(episode)+".html", []byte(content), 0)
		if err != nil {
			mw.openErrorMessBox("Error", err.Error())
			resetDownloadButton()
			return
		}

		c += 1
		processing = (float32(c) / float32(stop-start+1)) * 100
		buttonLog.SetText("Downloading... " + fmt.Sprintf("%.2f", processing) + "%")
	}
	buttonLog.SetText("Download")
	buttonLog.SetEnabled(true)
	mw.openInfoMessBox("Info", "Successfully Downloaded")
}

func (df *DownloadForm) KakaoPageDownload(folder string) {
	var processing float32 = 0.0
	start, err := strconv.Atoi(WebtoonDownloadForm.start.Text())
	if err != nil {
		mw.openWarningMessBox("Warning", "Invalid episode")
		return
	}
	stop, err := strconv.Atoi(WebtoonDownloadForm.stop.Text())
	if err != nil {
		mw.openWarningMessBox("Warning", "Invalid episode")
		return
	}

	_, err = os.Stat(WebtoonDownloadForm.folder + "\\" + WebtoonDownloadForm.id.Text())
	if os.IsNotExist(err) {
		err = os.MkdirAll(WebtoonDownloadForm.folder+"\\"+WebtoonDownloadForm.id.Text(), os.ModePerm)
		if err != nil {
			mw.openErrorMessBox("Error", err.Error())
			return
		}
	}

	c := 0
	buttonLog.SetEnabled(false)
	buttonLog.SetText("Downloading... " + fmt.Sprintf("%.2f", processing) + "%")

	// dl

	titleID := df.id.Text()

	ids, err := df.KakaoGetTitlesURL(titleID)
	if err != nil {
		mw.openWarningMessBox("Warning", err.Error())
		resetDownloadButton()
		return
	}

	if len(ids) <= 0 {
		mw.openWarningMessBox("Warning", "Wrong Webtoon Id")
		resetDownloadButton()
		return
	}

	for episode := start; episode <= stop; episode++ {
		if episode > len(ids) {
			mw.openWarningMessBox("Warning", "Can't find Epi"+strconv.Itoa(episode))
			resetDownloadButton()
			return
		}
		response, err := df.KakaoGetImgURL(ids[episode-1])
		if err != nil {
			mw.openErrorMessBox("Error", err.Error())
			resetDownloadButton()
			return
		}
		if len(response) <= 0 {
			mw.openWarningMessBox("Warning", "Epi"+strconv.Itoa(episode)+" - Can't find images")
			resetDownloadButton()
			return
		}

		err = os.MkdirAll(folder+"/"+strconv.Itoa(episode), os.ModePerm)
		if err != nil {
			mw.openErrorMessBox("Error", err.Error())
			resetDownloadButton()
			return
		}
		errchan := make(chan error, len(response))
		for anchor := range response {

			go downloadFile("http://page-edge-jz.kakao.com/sdownload/resource/"+response[anchor], folder+"/"+strconv.Itoa(episode)+"/"+strconv.Itoa(anchor+1)+".jpg", errchan)
		}
		for i := 0; i < len(response); i++ {
			err = <-errchan
			if err != nil {
				mw.openErrorMessBox("Error", err.Error())
				resetDownloadButton()
				return
			}
		}

		content := "<html><head><meta charset='UTF-8'><meta name='viewport' content='width=device-width, initial-scale=1.0'><style>body, html{margin: 0;border: 0;padding: 0;}@media only screen and (max-width: 700px) {img {width: 100%;}}</style><title>Episode " + strconv.Itoa(episode) + " (" + titleID + ")</title></head><body><center>"
		for l := 1; l <= len(response); l++ {
			content += "<img src='"
			content += strconv.Itoa(l)
			content += ".jpg'><br>"
		}
		content += "</body></center></html>"

		err = ioutil.WriteFile(folder+"/"+strconv.Itoa(episode)+"/"+strconv.Itoa(episode)+".html", []byte(content), 0)
		c += 1
		processing = (float32(c) / float32(stop-start+1)) * 100
		buttonLog.SetText("Downloading... " + fmt.Sprintf("%.2f", processing) + "%")
	}
	buttonLog.SetText("Download")
	buttonLog.SetEnabled(true)
	mw.openInfoMessBox("Info", "Successfully Downloaded")
}

func (df *DownloadForm) KakaoGetImgURL(productId string) ([]string, error) {
	data := make(map[string]string)
	data["productId"] = productId
	data["deviceId"] = df.deviceId
	resp, err := requestWithCookieNBody("https://api2-page.kakao.com/api/v1/inven/get_download_data/web", "POST", df.cookie.kakaoPageData, data)
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
	return output, nil
}

func (df *DownloadForm) KakaoGetTitlesURL(seriesid string) ([]string, error) {
	output := make([]string, 0)
	c := 0
	for {
		data := make(map[string]string)
		data["seriesid"] = seriesid
		data["page"] = strconv.Itoa(c)

		resp, err := requestWithCookieNBody("https://api2-page.kakao.com/api/v5/store/singles", "POST", "", data)
		if err != nil {
			return nil, err
		}

		downloadData, err := ioutil.ReadAll(resp)
		if err != nil {
			return nil, err
		}

		type idFormat struct {
			Id int `json:"id"`
		}
		var outputFiles []idFormat
		singles, err := jsonparser.GetUnsafeString(downloadData, "singles")
		if err != nil {
			return nil, err
		}
		json.Unmarshal([]byte(singles), &outputFiles)
		for i := range outputFiles {
			output = append(output, strconv.Itoa(outputFiles[i].Id))
		}
		if len(outputFiles) <= 0 {
			break
		}
		c += 1
	}
	return output, nil
}

func (df *DownloadForm) daumWebtoonDownload(folder string) {
	var processing float32 = 0.0
	start, err := strconv.Atoi(WebtoonDownloadForm.start.Text())
	if err != nil {
		mw.openWarningMessBox("Warning", "Invalid episode")
		return
	}
	stop, err := strconv.Atoi(WebtoonDownloadForm.stop.Text())
	if err != nil {
		mw.openWarningMessBox("Warning", "Invalid episode")
		return
	}

	_, err = os.Stat(WebtoonDownloadForm.folder + "\\" + WebtoonDownloadForm.id.Text())
	if os.IsNotExist(err) {
		err = os.MkdirAll(WebtoonDownloadForm.folder+"\\"+WebtoonDownloadForm.id.Text(), os.ModePerm)
		if err != nil {
			mw.openErrorMessBox("Error", err.Error())
			return
		}
	}

	c := 0
	buttonLog.SetEnabled(false)
	buttonLog.SetText("Downloading... " + fmt.Sprintf("%.2f", processing) + "%")

	// dl

	titleID := df.id.Text()

	ids, err := df.daumGetTitlesURL(titleID)
	if err != nil {
		mw.openErrorMessBox("Error", err.Error())
		resetDownloadButton()
		return
	}

	if len(ids) <= 0 {
		mw.openErrorMessBox("Error", "Wrong Webtoon Id")
		resetDownloadButton()
		return
	}
	for episode := start; episode <= stop; episode++ {
		if episode > len(ids) {
			mw.openWarningMessBox("Warning", "Can't find Epi"+strconv.Itoa(episode))
			resetDownloadButton()
			return
		}
		response, err := df.daumGetImgURL(ids[episode-1])
		if err != nil {
			mw.openErrorMessBox("Error", err.Error())
			resetDownloadButton()
			return
		}
		if len(response) <= 0 {
			mw.openWarningMessBox("Warning", "Epi"+strconv.Itoa(episode)+" - Can't find images")
			resetDownloadButton()
			return
		}

		err = os.MkdirAll(folder+"/"+strconv.Itoa(episode), os.ModePerm)
		if err != nil {
			mw.openErrorMessBox("Error", err.Error())
			resetDownloadButton()
			return
		}
		errchan := make(chan error, len(response))
		for anchor := range response {

			go downloadFile(response[anchor], folder+"/"+strconv.Itoa(episode)+"/"+strconv.Itoa(anchor+1)+".jpg", errchan)
		}
		for i := 0; i < len(response); i++ {
			err = <-errchan
			if err != nil {
				mw.openErrorMessBox("Error", err.Error())
				resetDownloadButton()
				return
			}
		}

		content := "<html><head><meta charset='UTF-8'><meta name='viewport' content='width=device-width, initial-scale=1.0'><style>body, html{margin: 0;border: 0;padding: 0;}@media only screen and (max-width: 700px) {img {width: 100%;}}</style><title>Episode " + strconv.Itoa(episode) + " (" + titleID + ")</title></head><body><center>"
		for l := 1; l <= len(response); l++ {
			content += "<img src='"
			content += strconv.Itoa(l)
			content += ".jpg'><br>"
		}
		content += "</body></center></html>"

		err = ioutil.WriteFile(folder+"/"+strconv.Itoa(episode)+"/"+strconv.Itoa(episode)+".html", []byte(content), 0)
		c += 1
		processing = (float32(c) / float32(stop-start+1)) * 100
		buttonLog.SetText("Downloading... " + fmt.Sprintf("%.2f", processing) + "%")
	}
	buttonLog.SetText("Download")
	buttonLog.SetEnabled(true)
	mw.openInfoMessBox("Info", "Successfully Downloaded")
}

func (df *DownloadForm) daumGetImgURL(productId string) ([]string, error) {
	resp, err := requestWithCookie("http://webtoon.daum.net/data/pc/webtoon/viewer_images/"+productId, "GET", "")
	if err != nil {
		return nil, err
	}
	downloadData, err := ioutil.ReadAll(resp)
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

func (df *DownloadForm) daumGetTitlesURL(seriesid string) ([]string, error) {
	resp, err := requestWithCookie("http://webtoon.daum.net/data/pc/webtoon/view/"+seriesid, "POST", "")
	if err != nil {
		return nil, err
	}
	downloadData, err := ioutil.ReadAll(resp)
	if err != nil {
		return nil, err
	}
	output := make([]int, 0)

	type eachFiles struct {
		Id int `json:"id"`
	}

	var outputFiles []eachFiles
	files, err := jsonparser.GetUnsafeString(downloadData, "data", "webtoon", "webtoonEpisodes")
	json.Unmarshal([]byte(files), &outputFiles)
	for i := range outputFiles {
		output = append(output, outputFiles[i].Id)
	}
	sort.Ints(output)
	strOutput := make([]string, len(output))
	for c := range output{
		strOutput[c] = strconv.Itoa(output[c])
	}
	return strOutput, nil
}

func (df *DownloadForm) lezhinComicsDownload(folder string) {
	var processing float32 = 0.0
	start, err := strconv.Atoi(WebtoonDownloadForm.start.Text())
	if err != nil {
		mw.openWarningMessBox("Warning", "Invalid episode")
		return
	}
	stop, err := strconv.Atoi(WebtoonDownloadForm.stop.Text())
	if err != nil {
		mw.openWarningMessBox("Warning", "Invalid episode")
		return
	}

	_, err = os.Stat(WebtoonDownloadForm.folder + "\\" + WebtoonDownloadForm.id.Text())
	if os.IsNotExist(err) {
		err = os.MkdirAll(WebtoonDownloadForm.folder+"\\"+WebtoonDownloadForm.id.Text(), os.ModePerm)
		if err != nil {
			mw.openErrorMessBox("Error", err.Error())
			return
		}
	}

	c := 0
	buttonLog.SetEnabled(false)
	buttonLog.SetText("Downloading... " + fmt.Sprintf("%.2f", processing) + "%")
	for episode := start; episode <= stop; episode++ {
		titleID := df.id.Text()

		reqHeader := make(map[string]string, 2)
		reqHeader["x-lz-locale"] = "ko_KR"
		resp, err := requestWithheader(lezhinBaseURL+"?alias="+df.id.Text()+"&name="+strconv.Itoa(episode)+"&type=comic_episode", "GET", reqHeader)
		if err != nil {
			mw.openErrorMessBox("Error", err.Error())
			resetDownloadButton()
			return
		}

		downloadData, err := ioutil.ReadAll(resp)
		if err != nil {
			mw.openErrorMessBox("Error", err.Error())
			resetDownloadButton()
			return
		}
		imgs := make([]string, 0)

		type eachFiles struct {
			Path string `json:"path"`
		}

		var outputFiles []eachFiles
		files, err := jsonparser.GetUnsafeString(downloadData, "data", "extra", "episode", "scrollsInfo")
		json.Unmarshal([]byte(files), &outputFiles)
		for i := range outputFiles {
			imgs = append(imgs, outputFiles[i].Path)
		}

		if len(imgs) <= 0 {
			mw.openWarningMessBox("Warning", "Can't find Images")
			resetDownloadButton()
			return
		}

		err = os.MkdirAll(folder+"/"+strconv.Itoa(episode), os.ModePerm)
		if err != nil {
			mw.openErrorMessBox("Error", err.Error())
			resetDownloadButton()
			return
		}
		errchan := make(chan error, len(imgs))
		for anchor := range imgs {
			go downloadFile("https://cdn.lezhin.com/v2"+imgs[anchor]+"?access_token="+df.accesstoken, folder+"/"+strconv.Itoa(episode)+"/"+strconv.Itoa(anchor+1)+".jpg", errchan)
		}
		for i := 0; i < len(imgs); i++ {
			err = <-errchan
			if err != nil {
				mw.openErrorMessBox("Error", err.Error())
				resetDownloadButton()
				return
			}
		}

		content := "<html><head><meta charset='UTF-8'><meta name='viewport' content='width=device-width, initial-scale=1.0'><style>body, html{margin: 0;border: 0;padding: 0;}@media only screen and (max-width: 700px) {img {width: 100%;}}</style><title>Episode " + strconv.Itoa(episode) + " (" + titleID + ")</title></head><body><center>"
		for l := 1; l <= len(imgs); l++ {
			content += "<img src='"
			content += strconv.Itoa(l)
			content += ".jpg'><br>"
		}
		content += "</body></center></html>"

		err = ioutil.WriteFile(folder+"/"+strconv.Itoa(episode)+"/"+strconv.Itoa(episode)+".html", []byte(content), 0)
		c += 1
		processing = (float32(c) / float32(stop-start+1)) * 100
		buttonLog.SetText("Downloading... " + fmt.Sprintf("%.2f", processing) + "%")
	}
	buttonLog.SetText("Download")
	buttonLog.SetEnabled(true)
	mw.openInfoMessBox("Info", "Successfully Downloaded")
}

func downloadFile(URL, fileName string, errchan chan error) {
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		errchan <- err
	}

	req.Header.Set("User-Agent", UserAgent)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		errchan <- err
	}

	var fileContent []byte
	fileContent, err = ioutil.ReadAll(res.Body)
	err = ioutil.WriteFile(fileName, fileContent, 0)
	if err != nil {
		errchan <- err
	}

	errchan <- nil
}

func requestWithCookie(url, method, cookie string) (io.Reader, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", UserAgent)
	req.Header.Set("Cookie", cookie)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}

func requestWithCookieNBody(urlname, method, cookie string, body map[string]string) (io.Reader, error) {
	data := &url.Values{}
	for i := range body {
		data.Add(i, body[i])
	}

	req, err := http.NewRequest(method, urlname, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", UserAgent)
	req.Header.Set("Cookie", cookie)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}

func requestWithheader(urlname, method string, header map[string]string) (io.Reader, error) {

	req, err := http.NewRequest(method, urlname, nil)
	if err != nil {
		return nil, err
	}
	for key, val := range header {
		req.Header.Add(key, val)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}

func resetDownloadButton() {
	buttonLog.SetText("Download")
	buttonLog.SetEnabled(true)
}
