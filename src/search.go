package main

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
	"github.com/buger/jsonparser"
)

func SearchWebtoonId(){
	var id string
	if WebtoonDownloadForm.wtype.CurrentIndex() == 0 {
		err, id = SearchNaverId(WebtoonDownloadForm.id.Text())
	} else if WebtoonDownloadForm.wtype.CurrentIndex() == 1 {
		err, id = SearchKakaoId(WebtoonDownloadForm.id.Text())
	} else if WebtoonDownloadForm.wtype.CurrentIndex() == 2 {
		err, id = SearchDaumId(WebtoonDownloadForm.id.Text())
	} else if WebtoonDownloadForm.wtype.CurrentIndex() == 3 {
		err, id = SearchLezhinId(WebtoonDownloadForm.id.Text())
	} else {
		mw.openWarningMessBox("Warning", "Please select type")
	}
	if err != nil {
		mw.openWarningMessBox("Warning", err.Error())
		return
	}
	WebtoonDownloadForm.id.SetText(id)
}
func SearchKakaoId(search string) (error, string) {
	resp, err := http.Get("https://page.kakao.com/search?word=" + url.QueryEscape(search))
	if err != nil {
		return err, ""
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return err, ""
	}

	NumImg := doc.Find(".css-4cffwv").Length()
	if NumImg <= 0 {
		return errors.New("No Search Result"), ""
	}
	firstResult := doc.Find(".css-4cffwv").First()
	firstResultUrl, _ := firstResult.Attr("href")
	err, firstResultId := getIdFromUrl(firstResultUrl)
	if err != nil {
		return err, ""
	}
	return nil, firstResultId
}

func SearchNaverId(search string) (error, string) {
	resp, err := http.Get("https://comic.naver.com/search.nhn?keyword=" + url.QueryEscape(search))
	if err != nil {
		return err, ""
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return err, ""
	}

	firstResultUrl, exist := doc.Find(".resultList").First().Find("li").Find("h5").Find("a").First().Attr("href")
	if !exist {
		return errors.New("No Search Result"), ""
	}
	err, firstResultId := getIdFromUrl(firstResultUrl)
	if err != nil {
		return err, ""
	}
	return nil, firstResultId
}


func SearchLezhinId(search string) (error, string) {
	resp, err := http.Get("https://dondog.lezhin.com/search?&v=2&type=comic&lang=ko&q=" + url.QueryEscape(search))
	if err != nil {
		return err, ""
	}
	defer resp.Body.Close()

	json, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err, ""
	}

	firstResultId, err := jsonparser.GetString(json, "sections", "[0]", "items", "[0]", "alias")
	if err != nil {
		return errors.New("No Search Result"), ""
	}
	return nil, firstResultId
}

func SearchDaumId(search string) (error, string) {
	resp, err := http.Get("http://webtoon.daum.net/data/pc/search/suggest?q=" + url.QueryEscape(search))
	if err != nil {
		return err, ""
	}
	defer resp.Body.Close()

	json, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err, ""
	}

	firstResultId, err := jsonparser.GetString(json, "data", "[0]", "nickname")
	if err != nil {
		return errors.New("No Search Result"), ""
	}
	return nil, firstResultId
}

func getIdFromUrl(url string) (error, string) {
	for i := 0; i < len(url); i++ {
		if url[i] == '=' {
			return nil, url[i+1:]
		}
	}
	return errors.New("Invalid Url"), ""
}
