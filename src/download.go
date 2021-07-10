package main

import (
	"errors"
	"os"

	"github.com/lxn/walk"
)

type DownloadData struct {
	Type         *walk.ComboBox
	TitleIdURL   *walk.LineEdit
	StartControl *walk.ListBox
	Start        *EnvModel
	StopControl  *walk.ListBox
	Stop         *EnvModel
	AccessToken  string
	Thread       int
	Folder       string
}

func WebtoonDownload() {
	if WDdata.StartControl.CurrentIndex() == -1 || WDdata.StopControl.CurrentIndex() == -1 {
		Log(2, errors.New("Please select episode"))
		return
	}
	LoadingOn()
	switch WDdata.Type.CurrentIndex() {
	case -1:
		Log(2, errors.New("Please Select Website"))
	case 0:
		makeNamspaceFolder(WDform.NaverComic.TitleId)
		code, err := WDform.NaverComic.Download(WDdata.StartControl.CurrentIndex(), WDdata.StopControl.CurrentIndex(), WDdata.Thread, WDdata.Folder+"\\"+WDform.NaverComic.TitleId)
		if err != nil {
			Log(code, err)
		}
	case 1:
		makeNamspaceFolder(WDform.KakaoPage.TitleId)
		code, err := WDform.KakaoPage.Download(WDdata.StartControl.CurrentIndex(), WDdata.StopControl.CurrentIndex(), WDdata.Thread, WDdata.Folder+"\\"+WDform.KakaoPage.TitleId)
		if err != nil {
			Log(code, err)
		}
	case 2:
		makeNamspaceFolder(WDform.DaumWebtoon.TitleId)
		code, err := WDform.DaumWebtoon.Download(WDdata.StartControl.CurrentIndex(), WDdata.StopControl.CurrentIndex(), WDdata.Thread, WDdata.Folder+"\\"+WDform.DaumWebtoon.TitleId)
		if err != nil {
			Log(code, err)
		}
	case 3:
		makeNamspaceFolder(WDform.LezhinComics.TitleId)
		code, err := WDform.LezhinComics.Download(WDdata.StartControl.CurrentIndex(), WDdata.StopControl.CurrentIndex(), WDdata.Thread, WDdata.Folder+"\\"+WDform.LezhinComics.TitleId)
		if err != nil {
			Log(code, err)
		}
	case 4:
		makeNamspaceFolder(WDform.KPepub.TitleId)
		code, err := WDform.KPepub.Download(WDdata.StartControl.CurrentIndex(), WDdata.StopControl.CurrentIndex(), WDdata.Folder+"\\"+WDform.KPepub.TitleId)
		if err != nil {
			Log(code, err)
		}
	case 5:
		makeNamspaceFolder(WDform.RidiWT.TitleId)
		code, err := WDform.RidiWT.Download(WDdata.StartControl.CurrentIndex(), WDdata.StopControl.CurrentIndex(), WDdata.Thread, WDdata.Folder+"\\"+WDform.RidiWT.TitleId)
		if err != nil {
			Log(code, err)
		}
	}
	LoadingOff()
}

func makeNamspaceFolder(id string) {
	_, err = os.Stat(WDdata.Folder + "\\" + id)
	if os.IsNotExist(err) {
		err = os.MkdirAll(WDdata.Folder+"\\"+id, os.ModePerm)
		if err != nil {
			mw.openErrorMessBox("Error", err.Error())
			return
		}
	}
}
