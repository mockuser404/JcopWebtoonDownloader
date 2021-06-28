package main

import (
	"log"
	"os/exec"

	"github.com/lxn/walk"
	"github.com/mynameispyo/JcopWebtoonDownloader/WTdown"

	. "github.com/lxn/walk/declarative"
)

var (
	mw = new(MyMainWindow)

	WDdata = DownloadData{Start: &EnvModel{items: make([]string, 0)},Stop: &EnvModel{items: make([]string, 0)}, Thread: 70}
	WDform = WTdown.WTdown{}

	buttonLog = new(walk.PushButton)
	err       error
)

func main() {
	err = MainWindow{
		Icon:     "img\\downloader.ico",
		AssignTo: &mw.MainWindow,
		Title:    "Jcop Webtoon Downloader",
		MinSize:  Size{Width:420, Height:140},
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
								OnTriggered: setLezhinComicCookieData,
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
							WDdata.Folder = dlg.FilePath
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
							err = exec.Command("rundll32", "url.dll,FileProtocolHandler", WDdata.Folder).Start()
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
							mw.openInfoMessBox("Version", Version)
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
			HSplitter{
				Column: 3,
				Children: []Widget{
					ListBox{
						AssignTo: &(WDdata.StartControl),
						Model:    WDdata.Start,
					},
					ListBox{
						AssignTo: &(WDdata.StopControl),
						Model:    WDdata.Stop,
					},
					Composite{
						Layout: VBox{},
						Children: []Widget{
							Composite{
								Layout: Grid{Columns: 2, Spacing: 10},
								Children: []Widget{

									Label{
										Font: Font{PointSize: 12},
										Text: "Website",
									},
									ComboBox{
										Font:     Font{PointSize: 12},
										Model:    GetWebtoonTypes(),
										AssignTo: &(WDdata.Type),

										BindingMember: "Id",
										DisplayMember: "Type",
									},

									Label{
										Font: Font{PointSize: 12},
										Text: "URL",
									},
									Composite{
										Layout: HBox{Spacing: 3, MarginsZero: true},
										Children: []Widget{
											LineEdit{
												Font:     Font{PointSize: 12},
												AssignTo: &(WDdata.TitleIdURL),
											},
											PushButton{
												Font:      Font{PointSize: 12},
												MaxSize:   Size{Width:35, Height:10},
												Text:      "🔍",
												OnClicked: LoadEpis,
											},
										},
									},
								},
							},
							PushButton{

								Font: Font{PointSize: 12},
								// AssignTo: &acceptPB,
								AssignTo: &buttonLog,
								Text:     "Download",
								OnClicked: func() {

									go WebtoonDownload()

								},
							},
						},
					},
				},
			},
		},
		// Children: []Widget{},
	}.Create()

	if err != nil {
		log.Println(err)
	}

	mw.SetBounds(walk.Rectangle{X:0, Y:0, Width:500, Height:140}) // You can use GetSystemMetrics from the `win` package to get the screen resolution

	loadSettingData()
	NewVersionCheck()
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
