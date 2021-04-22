package main

import (
	"encoding/json"
	"os"
	"os/user"

	"github.com/lxn/walk"
)

const (
	ProgramVersion = "v3.1.1"
	BlogUrl        = "https://blog.naver.com/the3countrys/222106929101"

	appData string = "\\AppData\\Local\\JcopWebtoonDownloader"
)

var (
	homeDir string
)

func setNaverComicCookieData() {
	naverComicData := &naverComicCookieForm{}
	if cmd, err := naverComicData.RunDialog(mw); err != nil {
		mw.openErrorMessBox("Error", err.Error())
	} else if cmd == walk.DlgCmdOK {
		WebtoonDownloadForm.cookie.naverComicData = "NID_AUT=" + naverComicData.NIDAUT + "; NID_SES=" + naverComicData.NIDSES + ";"
		SaveFormData()
	}
}

func setKakaoPageCookieData() {
	kakaoPageData := &kakaoPageCookieForm{}
	if cmd, err := kakaoPageData.RunDialog(mw); err != nil {
		mw.openErrorMessBox("Error", err.Error())
	} else if cmd == walk.DlgCmdOK {
		WebtoonDownloadForm.cookie.kakaoPageData = `_kawlt=` + kakaoPageData.Kawlt + `; _kawlp=` + kakaoPageData.Kawlp + `; _kawlptea=` + kakaoPageData.Kawlptea + `;`
		WebtoonDownloadForm.deviceId = kakaoPageData.DeviceId
		SaveFormData()
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
		WebtoonDownloadForm.folder = homeDir + "\\Documents\\Jcop Webtoon Downloader"
		return nil
	}
	dat := make([]byte, 9999)
	size, err := defaultDir.Read(dat)

	if err != nil {
		return err
	}
	WebtoonDownloadForm.folder = string(dat[:size])
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
	DeviceId    string `json:"deviceid"`
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
		mkfile.WriteString(`{"naver":"","kakao":"","accesstoken":"","deviceid":""}`)
		WebtoonDownloadForm.folder = homeDir + "\\Documents\\Jcop Webtoon Downloader"
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
	WebtoonDownloadForm.cookie.naverComicData = readDataForm.Naver
	WebtoonDownloadForm.cookie.kakaoPageData = readDataForm.Kakao
	WebtoonDownloadForm.accesstoken = readDataForm.Accesstoken
	WebtoonDownloadForm.deviceId = readDataForm.DeviceId
	return nil
}

func SaveFormData() error {
	newFormData := dataForm{
		Naver:       WebtoonDownloadForm.cookie.naverComicData,
		Kakao:       WebtoonDownloadForm.cookie.kakaoPageData,
		Accesstoken: WebtoonDownloadForm.accesstoken,
		DeviceId:    WebtoonDownloadForm.deviceId,
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
