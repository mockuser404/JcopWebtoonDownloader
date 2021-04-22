package main

import (
	"github.com/lxn/walk"

	. "github.com/lxn/walk/declarative"
)

type cookieData struct {
	naverComicData   string
	kakaoPageData    string
}

type naverComicCookieForm struct {
	NIDAUT string
	NIDSES string
}

func (nc *naverComicCookieForm) RunDialog(owner walk.Form) (int, error) {
	var dlg *walk.Dialog
	var db *walk.DataBinder
	var acceptPB, cancelPB *walk.PushButton

	return Dialog{
		AssignTo:      &dlg,
		Title:         "Set Naver Comic Cookie Data",
		DefaultButton: &acceptPB,
		CancelButton:  &cancelPB,
		DataBinder: DataBinder{
			AssignTo:       &db,
			Name:           "nc",
			DataSource:     nc,
			ErrorPresenter: ToolTipErrorPresenter{},
		},
		MinSize: Size{300, 200},
		Layout:  VBox{},
		Children: []Widget{
			Composite{
				Layout: Grid{Columns: 2},
				Children: []Widget{
					Label{
						Text: "NID_AUT:",
					},
					LineEdit{
						Text: Bind("NIDAUT"),
					},
					Label{
						Text: "NID_SES:",
					},
					LineEdit{
						Text: Bind("NIDSES"),
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
							if err := db.Submit(); err != nil {
								mw.openErrorMessBox("Error", err.Error())
								return
							}

							dlg.Accept()
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

type kakaoPageCookieForm struct {
	Kawlp    string
	Kawlptea string
	Kawlt    string
	DeviceId string
}

func (kp *kakaoPageCookieForm) RunDialog(owner walk.Form) (int, error) {
	var dlg *walk.Dialog
	var db *walk.DataBinder
	var acceptPB, cancelPB *walk.PushButton

	return Dialog{
		AssignTo:      &dlg,
		Title:         "Set Kakao Page Cookie Data",
		DefaultButton: &acceptPB,
		CancelButton:  &cancelPB,
		DataBinder: DataBinder{
			AssignTo:       &db,
			Name:           "kp",
			DataSource:     kp,
			ErrorPresenter: ToolTipErrorPresenter{},
		},
		MinSize: Size{300, 200},
		Layout:  VBox{},
		Children: []Widget{
			Composite{
				Layout: Grid{Columns: 2},
				Children: []Widget{
					Label{
						Text: "_kawlp:",
					},
					LineEdit{
						Text: Bind("Kawlp"),
					},
					Label{
						Text: "_kawlptea:",
					},
					LineEdit{
						Text: Bind("Kawlptea"),
					},
					Label{
						Text: "_kawlt:",
					},
					LineEdit{
						Text: Bind("Kawlt"),
					},
					Label{
						Text: "DeviceId:",
					},
					LineEdit{
						Text: Bind("DeviceId"),
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
							if err := db.Submit(); err != nil {
								mw.openErrorMessBox("Error", err.Error())
								return
							}

							dlg.Accept()
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

func lezhinRunDialog(owner walk.Form) (int, error) {
	var dlg *walk.Dialog
	var accesstokenLineEdit *walk.LineEdit
	var acceptPB, cancelPB *walk.PushButton

	return Dialog{
		AssignTo:      &dlg,
		Title:         "Set Lezhin access token",
		DefaultButton: &acceptPB,
		CancelButton:  &cancelPB,
		MinSize:       Size{300, 200},
		Layout:        VBox{},
		Children: []Widget{
			Composite{
				Layout: Grid{Columns: 2},
				Children: []Widget{
					Label{
						Text: "access_token:",
					},
					LineEdit{
						AssignTo: &accesstokenLineEdit,
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
							WebtoonDownloadForm.accesstoken = accesstokenLineEdit.Text()
							SaveFormData()
							dlg.Cancel()
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
