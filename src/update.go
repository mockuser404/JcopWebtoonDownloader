package main

import (
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"

	"github.com/buger/jsonparser"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

const (
	Version = "v4.2.2"
)

func NewVersionCheck() {
	resp, err := http.Get("https://api.github.com/repos/mynameispyo/JcopWebtoonDownloader/releases/latest")
	if err != nil {
		Log(1, err)
	}

	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		Log(1, err)
	}
	newestVersion, err := jsonparser.GetString(content, "tag_name")
	if err != nil {
		Log(1, err)
	}

	downloadUrl, err := jsonparser.GetString(content, "assets", "[0]", "browser_download_url")
	if err != nil {
		Log(1, err)
	}

	if newestVersion != Version {
		if _, err := updateDialog(mw, newestVersion, downloadUrl); err != nil {
			Log(1, err)
		}
	} else {
		_, err = os.Stat(homeDir + appData + "\\tmp")
		if !os.IsNotExist(err) {
			err = os.RemoveAll(homeDir + appData + "\\tmp")
			if err != nil {
				Log(2, errors.New("Fail to remove tmp folder"))
			}
		}
	}

}

func updateDialog(owner walk.Form, newestVersion, downloadUrl string) (int, error) {
	var dlg *walk.Dialog
	// accesstokenLineEdit := make([]*walk.LineEdit, len(requireCookies))
	var acceptPB, cancelPB *walk.PushButton
	var progressValue *walk.ProgressBar
	return Dialog{
		AssignTo:      &dlg,
		Title:         "New Version available",
		DefaultButton: &acceptPB,
		CancelButton:  &cancelPB,
		MinSize:       Size{Width:300, Height:200},
		Layout:        VBox{},
		Children: []Widget{
			Composite{
				Layout: VBox{},
				Children: []Widget{
					Label{
						Text: "Install " + newestVersion,
					},
					ProgressBar{
						AssignTo: &(progressValue),
						Value:    0,
					},
				},
			},
			Composite{
				Layout: HBox{},
				Children: []Widget{
					HSpacer{},
					PushButton{
						AssignTo: &acceptPB,
						Text:     "OK",
						OnClicked: func() {
							// progressValue.SetValue()
							err = os.MkdirAll(homeDir+appData+"\\tmp", 0)
							if err != nil {
								Log(1, errors.New("Fail to make tmp directory"))
							}
							progressValue.SetValue(100)
							downloadFile(homeDir+appData+"\\tmp\\JcopWebtoonDownloaderSetup.exe", downloadUrl)
							err = exec.Command(homeDir + appData + "\\tmp\\JcopWebtoonDownloaderSetup.exe").Start()
							if err != nil {
								mw.openErrorMessBox("Error", err.Error())
							}
							os.Exit(0)

							// dlg.Cancel()
						},
					},
					PushButton{
						AssignTo:  &cancelPB,
						Text:      "Cancel",
						OnClicked: func() { dlg.Cancel() },
					},
				},
			},
		},
	}.Run(owner)
}

func downloadFile(fileName string, url string) error {
	out, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer out.Close()
	
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
