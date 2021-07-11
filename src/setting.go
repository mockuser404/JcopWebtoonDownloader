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
	requireCookies := make([]string, 2)
	requireCookies[0] = "RSESSION"
	requireCookies[1] = "cc"
	if _, err := lezhinRunDialog(mw, &(WDform.LezhinComics.Cookies), requireCookies); err != nil {
		Log(1, err)
	}
}

func setRidibooksWebtoonCookieData() {
	requireCookies := make([]string, 1)
	requireCookies[0] = "ridi-at"
	if _, err := askCookie(mw, &(WDform.RidiWT.Cookies), requireCookies); err != nil {
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
	// err = loadDefaultDir()
	// if err != nil {
	// 	mw.openErrorMessBox("Error", err.Error())
	// }

	err = loadFormData()
	if err != nil {
		mw.openErrorMessBox("Error", err.Error())
	}
}

// func loadDefaultDir() error {
// 	defaultDir, err := os.Open(homeDir + appData + "\\bgdata.json")
// 	defer defaultDir.Close()

// 	if err != nil {
// 		mkfile, err := os.Create(homeDir + appData + "\\bgdata.json")
// 		defer mkfile.Close()
// 		if err != nil {
// 			return err
// 		}
// 		mkfile.WriteString(homeDir + "\\Documents\\Jcop Webtoon Downloader")
// 		WDdata.Folder = homeDir + "\\Documents\\Jcop Webtoon Downloader"
// 		return nil
// 	}
// 	dat := make([]byte, 9999)
// 	size, err := defaultDir.Read(dat)

// 	if err != nil {
// 		return err
// 	}

// 	WDdata.Folder = string(dat[:size])
// 	return nil
// }

// func setDefaultDir(path string) error {
// 	defaultDir, err := os.Create(homeDir + appData + "\\DefaultDir")
// 	defer defaultDir.Close()
// 	if err != nil {
// 		return err
// 	}
// 	defaultDir.WriteString(path)
// 	return nil
// }

type dataForm struct {
	DefaultDir  string `json:"defaultDir"`
	Thread      int    `json:"thread"`
	Lang        string `json:"lang"`
	Naver       string `json:"naver"`
	Kakao       string `json:"kakao"`
	Lezhin      string `json:"lezhin"`
	Accesstoken string `json:"accesstoken"`
	Ridiwt      string `json:"ridiwt"`
}

func loadFormData() error {
	var readDataForm dataForm
	defaultDir, err := os.Open(homeDir + appData + "\\FormData.json")
	defer defaultDir.Close()

	if err != nil {
		mkfile, err := os.Create(homeDir + appData + "\\FormData.json")
		defer mkfile.Close()
		if err != nil {
			return err
		}
		readDataForm.DefaultDir = homeDir + "\\Documents\\Jcop Webtoon Downloader"
		readDataForm.Thread = 70
		readDataForm.Lang = "ko"
		emptyJson, err := json.Marshal(readDataForm)
		if err != nil {
			return err
		}
		mkfile.WriteString(string(emptyJson))
		WDdata.Folder = homeDir + "\\Documents\\Jcop Webtoon Downloader"
		return nil
	}
	dat := make([]byte, 19999)
	size, err := defaultDir.Read(dat)

	if err != nil {
		return err
	}
	err = json.Unmarshal(dat[:size], &readDataForm)
	if err != nil {
		return nil
	}
	WDdata.Folder = readDataForm.DefaultDir
	WDform.NaverComic.Cookies = readDataForm.Naver
	WDform.KakaoPage.Cookies = readDataForm.Kakao
	WDform.LezhinComics.Cookies = readDataForm.Lezhin
	WDform.LezhinComics.AccessToken = readDataForm.Accesstoken
	WDform.RidiWT.Cookies = readDataForm.Ridiwt
	if readDataForm.Lang == ""{
		WDform.LezhinComics.Language = "ko"
	}else{
		WDform.LezhinComics.Language = readDataForm.Lang
	}
	if readDataForm.Thread == 0 {
		WDdata.Thread = 70
	} else {
		WDdata.Thread = readDataForm.Thread
	}

	return nil
}

func SaveFormData() error {
	newFormData := dataForm{
		DefaultDir:  WDdata.Folder,
		Thread:      WDdata.Thread,
		Lang:        WDform.LezhinComics.Language,
		Naver:       WDform.NaverComic.Cookies,
		Kakao:       WDform.KakaoPage.Cookies,
		Lezhin:      WDform.LezhinComics.Cookies,
		Accesstoken: WDform.LezhinComics.AccessToken,
		Ridiwt:      WDform.RidiWT.Cookies,
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
