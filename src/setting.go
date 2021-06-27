package main

import (
	"encoding/json"
	"os"
	"os/user"
)

const (
	BlogUrl = "https://blog.naver.com/the3countrys/222106929101"

	appData string = "\\AppData\\Local\\JcopWebtoonDownloader"
)

var (
	homeDir string
)

func setNaverComicCookieData() {
	requireCookies := make([]string, 2)
	requireCookies[0] = "NID_AUT"
	requireCookies[1] = "NID_SES"
	if _, err := askCookie(mw, &(WDform.NaverComic.Cookies), requireCookies); err != nil {
		Log(1, err)
	}
}

func setKakaoPageCookieData() {
	requireCookies := make([]string, 3)
	requireCookies[0] = "_kpawlt"
	requireCookies[1] = "_kpawltea"
	requireCookies[2] = "_kpawlst"
	if _, err := askCookie(mw, &(WDform.KakaoPage.Cookies), requireCookies); err != nil {
		Log(1, err)
	}
}

func setLezhinComicCookieData() {
	if _, err := lezhinRunDialog(mw); err != nil {
		Log(1, err)
	}
}

func loadSettingData() {
	user, err := user.Current()
	if err != nil {
		mw.openErrorMessBox("Error", err.Error())
	}
	homeDir = user.HomeDir
	file, err := os.Open(homeDir + appData)
	defer file.Close()
	if err != nil {
		err = os.MkdirAll(user.HomeDir+appData, os.ModePerm)
		if err != nil {
			mw.openErrorMessBox("Error", err.Error())
		}
	}
	err = loadDefaultDir()
	if err != nil {
		mw.openErrorMessBox("Error", err.Error())
	}

	err = loadFormData()
	if err != nil {
		mw.openErrorMessBox("Error", err.Error())
	}
}

func loadDefaultDir() error {
	defaultDir, err := os.Open(homeDir + appData + "\\DefaultDir")
	defer defaultDir.Close()

	if err != nil {
		mkfile, err := os.Create(homeDir + appData + "\\DefaultDir")
		defer mkfile.Close()
		if err != nil {
			return err
		}
		mkfile.WriteString(homeDir + "\\Documents\\Jcop Webtoon Downloader")
		WDdata.Folder = homeDir + "\\Documents\\Jcop Webtoon Downloader"
		return nil
	}
	dat := make([]byte, 9999)
	size, err := defaultDir.Read(dat)

	if err != nil {
		return err
	}
	WDdata.Folder = string(dat[:size])
	return nil
}

func setDefaultDir(path string) error {
	defaultDir, err := os.Create(homeDir + appData + "\\DefaultDir")
	defer defaultDir.Close()
	if err != nil {
		return err
	}
	defaultDir.WriteString(path)
	return nil
}

type dataForm struct {
	Naver       string `json:"naver"`
	Kakao       string `json:"kakao"`
	Accesstoken string `json:"accesstoken"`
}

func loadFormData() error {
	defaultDir, err := os.Open(homeDir + appData + "\\FormData.json")
	defer defaultDir.Close()

	if err != nil {
		mkfile, err := os.Create(homeDir + appData + "\\FormData.json")
		defer mkfile.Close()
		if err != nil {
			return err
		}
		mkfile.WriteString(`{"naver":"","kakao":"","accesstoken":""}`)
		WDdata.Folder = homeDir + "\\Documents\\Jcop Webtoon Downloader"
		return nil
	}
	dat := make([]byte, 9999)
	size, err := defaultDir.Read(dat)

	if err != nil {
		return err
	}
	var readDataForm dataForm
	err = json.Unmarshal(dat[:size], &readDataForm)
	if err != nil {
		return nil
	}
	WDform.NaverComic.Cookies = readDataForm.Naver
	WDform.KakaoPage.Cookies = readDataForm.Kakao
	WDform.LezhinComics.AccessToken = readDataForm.Accesstoken
	return nil
}

func SaveFormData() error {
	newFormData := dataForm{
		Naver:       WDform.NaverComic.Cookies,
		Kakao:       WDform.KakaoPage.Cookies,
		Accesstoken: WDform.LezhinComics.AccessToken,
	}
	stringFormData, err := json.Marshal(newFormData)
	if err != nil {
		return err
	}
	formDataWriter, err := os.Create(homeDir + appData + "\\FormData.json")
	defer formDataWriter.Close()
	if err != nil {
		return err
	}
	formDataWriter.Write(stringFormData)
	return nil
}
