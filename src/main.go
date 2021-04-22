package main

import (
	"os/exec"

	"github.com/lxn/walk"

	. "github.com/lxn/walk/declarative"
)

var (
	mw                  = new(MyMainWindow)
	WebtoonDownloadForm = DownloadForm{}
	buttonLog           = new(walk.PushButton)
)

func main() {
	MainWindow{
		Icon:     "img\\downloader.ico",
		AssignTo: &mw.MainWindow,
		Title:    "Jcop Webtoon Downloader",
		MinSize:  Size{420, 140},
		Layout:   VBox{},
		MenuItems: []MenuItem{
			Menu{
				Text: "&Setting",
				Items: []MenuItem{
					Menu{
						Text: "Set Data",
						Items: []MenuItem{
							Action{
								Text: "Naver Comic",
								OnTriggered: setNaverComicCookieData,
							},
							Action{
								Text: "Kakao Page",
								OnTriggered: setKakaoPageCookieData,
							},
							Action{
								Text: "Lezhin Comics",
								OnTriggered: func() {
									lezhinRunDialog(mw)
								},
							},
						},
					},
					Action{
						Text: "Set Default Directory",
						OnTriggered: func() {

							dlg := new(walk.FileDialog)

							dlg.Title = "Select Folder"

							if ok, err := dlg.ShowBrowseFolder(mw); err != nil {
								mw.openErrorMessBox("Error", err.Error())
								return
							} else if !ok {
								return
							}
							err = setDefaultDir(dlg.FilePath)
							if err != nil {
								mw.openErrorMessBox("Error", err.Error())
							}
							WebtoonDownloadForm.folder = dlg.FilePath
						},
					},
				},
			},
			Menu{
				Text: "&Tool",
				Items: []MenuItem{
					Action{
						Text: "Open Default Directory",
						OnTriggered: func() {
							err = exec.Command("rundll32", "url.dll,FileProtocolHandler", WebtoonDownloadForm.folder).Start()
							if err != nil {
								mw.openErrorMessBox("Error", err.Error())
							}
						},
					},
				},
			},
			Menu{
				Text: "&Help",
				Items: []MenuItem{
					Action{
						Text: "Version",
						OnTriggered: func() {
							mw.openInfoMessBox("Version", ProgramVersion)
						},
					},
					Action{
						Text: "Vist Blog",
						OnTriggered: func() {
							err = exec.Command("rundll32", "url.dll,FileProtocolHandler", BlogUrl).Start()
							if err != nil {
								mw.openErrorMessBox("Error", err.Error())
							}
						},
					},
				},
			},
		},
		Children: []Widget{
			Composite{
				Layout: Grid{Columns: 2, Spacing: 10},
				Children: []Widget{

					Label{
						Font: Font{PointSize: 12, Bold: true},
						Text: "Type",
					},
					ComboBox{
						Font:     Font{PointSize: 12, Bold: true},
						Model:    GetWebtoonTypes(),
						AssignTo: &(WebtoonDownloadForm.wtype),

						BindingMember: "Id",
						DisplayMember: "Type",
					},

					Label{
						Font: Font{PointSize: 12, Bold: true},
						Text: "ID",
					},
					Composite{
						Layout: HBox{Spacing: 3, MarginsZero: true},
						Children: []Widget{
							LineEdit{
								Font:     Font{PointSize: 12, Bold: true},
								AssignTo: &(WebtoonDownloadForm.id),
							},
							PushButton{
								Font: Font{PointSize: 12, Bold: true},
								MaxSize: Size{35,10},
								Text: "🔍",
								OnClicked: SearchWebtoonId,
							},
						},
					},

					Label{
						Font: Font{PointSize: 12, Bold: true},
						Text: "Episodes",
					},
					Composite{
						Layout: HBox{Spacing: 3, MarginsZero: true},
						Children: []Widget{
							LineEdit{
								Font:     Font{PointSize: 12, Bold: true},
								AssignTo: &(WebtoonDownloadForm.start),
							},
							Label{
								Font: Font{PointSize: 12, Bold: true},
								Text: "~",
							},
							LineEdit{
								Font:     Font{PointSize: 12, Bold: true},
								AssignTo: &(WebtoonDownloadForm.stop),
							},
						},
					},
				},
			},
			Composite{
				Layout: HBox{},
				Children: []Widget{
					PushButton{

						Font: Font{PointSize: 12, Bold: true},
						// AssignTo: &acceptPB,
						AssignTo: &buttonLog,
						Text:     "Download",
						OnClicked: func() {

							go webtoonDownload()

						},
					},
				},
			},
		},
	}.Create()
	loadSettingData()

	mw.SetBounds(walk.Rectangle{0, 0, 420, 140}) // You can use GetSystemMetrics from the `win` package to get the screen resolution

	mw.Run()
}

type MyMainWindow struct {
	*walk.MainWindow
}

func (mw *MyMainWindow) openInfoMessBox(title, message string) {
	walk.MsgBox(mw, title, message, walk.MsgBoxIconInformation)
}

func (mw *MyMainWindow) openWarningMessBox(title, message string) {
	walk.MsgBox(mw, title, message, walk.MsgBoxIconWarning)
}

func (mw *MyMainWindow) openErrorMessBox(title, message string) {
	walk.MsgBox(mw, title, message, walk.MsgBoxIconError)
}

type WebtoonType struct {
	Id   int
	Type string
}

func GetWebtoonTypes() []*WebtoonType {
	return []*WebtoonType{
		{1, "Naver Comic"},
		{2, "Kakao Page"},
		{3, "Daum Webtoon"},
		{4, "Lezhin Comics"},
	}
}
