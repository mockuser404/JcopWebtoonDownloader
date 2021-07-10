package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"

	"github.com/buger/jsonparser"
)

func TestNewVersionCheck(t *testing.T) {
	resp, err := http.Get("https://api.github.com/repos/mynameispyo/JcopWebtoonDownloader/releases/latest")
	if err != nil {
		log.Println(1, err)
	}

	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(1, err)
	}
	_, err = jsonparser.GetString(content, "tag_name")
	if err != nil {
		log.Println(1, err)
	}

	assetsJson, err := jsonparser.GetUnsafeString(content, "assets")
	if err != nil {
		log.Println(1, err)
	}
	type assetsFormat struct {
		Name string `json:"name"`
		BrowserDownloadUrl string `json:"browser_download_url"`
	}
	var assets []assetsFormat
	err = json.Unmarshal([]byte(assetsJson), &assets)
	var downloadUrl string
	if err != nil {
		log.Println(2, err)
	}
	for i := range assets {
		lastExe := assets[i].Name[len(assets[i].Name)-4:len(assets[i].Name)]
		if lastExe != ".exe"{
			return
		}
		name := strings.Split(assets[i].Name[:len(assets[i].Name)-4], "_")
		arch := name[len(name)-1]
		if Architecture == arch {
			downloadUrl = assets[i].BrowserDownloadUrl
		}
	}
	if downloadUrl == "" {
		log.Println(1, errors.New("Can't find appropriate Architecture"))
	}else{
		log.Println(downloadUrl)
	}
	t.Error("test")
}
