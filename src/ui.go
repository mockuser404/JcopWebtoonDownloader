package main

import (
	"errors"
	"strconv"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func askCookie(owner walk.Form, cookieType *string, requireCookies []string) (int, error) {
	var dlg *walk.Dialog
	accesstokenLineEdit := make([]*walk.LineEdit, len(requireCookies))
	var acceptPB, cancelPB *walk.PushButton

	var eachForm []Widget
	for i := range requireCookies {
		eachForm = append(eachForm, Label{
			Text: requireCookies[i],
		})
		eachForm = append(eachForm, LineEdit{
			AssignTo: &accesstokenLineEdit[i],
		})
	}

	return Dialog{
		AssignTo:      &dlg,
		Title:         "Set Cookie data",
		DefaultButton: &acceptPB,
		CancelButton:  &cancelPB,
		MinSize:       Size{Width: 300, Height: 200},
		Layout:        VBox{},
		Children: []Widget{
			Composite{
				Layout:   Grid{Columns: 2},
				Children: eachForm,
			},
			Composite{
				Layout: HBox{},
				Children: []Widget{
					HSpacer{},
					PushButton{
						AssignTo: &acceptPB,
						Text:     "OK",
						OnClicked: func() {
							*cookieType = ""
							for i := range requireCookies {
								*cookieType += requireCookies[i] + "=" + accesstokenLineEdit[i].Text() + "; "
							}
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

func lezhinRunDialog(owner walk.Form, cookieType *string, requireCookies []string) (int, error) {
	var dlg *walk.Dialog
	accesstokenLineEdit := make([]*walk.LineEdit, len(requireCookies)+1)
	var languageSelect *walk.ComboBox
	var acceptPB, cancelPB *walk.PushButton

	var eachForm []Widget
	for i := range requireCookies {
		eachForm = append(eachForm, Label{
			Text: requireCookies[i],
		})
		eachForm = append(eachForm, LineEdit{
			AssignTo: &accesstokenLineEdit[i],
		})
	}
	eachForm = append(eachForm, Label{
		Text: "access_token(NOT COOKIE)",
	})
	eachForm = append(eachForm, LineEdit{
		AssignTo: &accesstokenLineEdit[len(requireCookies)],
	})

	eachForm = append(eachForm, Label{
		Text: "Language",
	})
	eachForm = append(eachForm, ComboBox{
						
		Model:    getLezhinLangs(),
		AssignTo: &(languageSelect),

		BindingMember: "Real",
		DisplayMember: "Display",
	})

	return Dialog{
		AssignTo:      &dlg,
		Title:         "Set Lezhin access token",
		DefaultButton: &acceptPB,
		CancelButton:  &cancelPB,
		MinSize:       Size{Width: 300, Height: 200},
		Layout:        VBox{},
		Children: []Widget{
			Composite{
				Layout:   Grid{Columns: 2},
				Children: eachForm,
			},
			Composite{
				Layout: HBox{},
				Children: []Widget{
					HSpacer{},
					PushButton{
						AssignTo: &acceptPB,
						Text:     "OK",
						OnClicked: func() {
							*cookieType = ""
							for i := range requireCookies {
								*cookieType += requireCookies[i] + "=" + accesstokenLineEdit[i].Text() + "; "
							}
							WDform.LezhinComics.AccessToken = accesstokenLineEdit[len(requireCookies)].Text()
							WDform.LezhinComics.Language = languageSelect.Text()
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

type LezhinLanguage struct {
	Real    string
	Display string
}

func getLezhinLangs() []*LezhinLanguage {
	return []*LezhinLanguage{
		{"ko", "ko"},
		{"en", "en"},
	}
}

func setThread(owner walk.Form) (int, error) {
	var dlg *walk.Dialog
	var threadedit *walk.LineEdit
	var acceptPB, cancelPB *walk.PushButton

	return Dialog{
		AssignTo:      &dlg,
		Title:         "Set Cookie data",
		DefaultButton: &acceptPB,
		CancelButton:  &cancelPB,
		MinSize:       Size{Width: 300, Height: 200},
		Layout:        VBox{},
		Children: []Widget{
			Composite{
				Layout: Grid{Columns: 2},
				Children: []Widget{
					Label{
						Text: "Thread: ",
					},
					LineEdit{
						Text:     strconv.Itoa(WDdata.Thread),
						AssignTo: &threadedit,
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
							threadinNum, err := strconv.Atoi(threadedit.Text())
							if err != nil {
								Log(2, errors.New("Invalid num"))
								return
							}
							WDdata.Thread = threadinNum
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
