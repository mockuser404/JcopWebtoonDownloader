package main

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/buger/jsonparser"
)

func main() {
	err, res := searchDaumId("교수인형")
	if err != nil {
		log.Println(err)
	}
	log.Println(res)
}
func searchDaumId(search string) (error, string) {
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
