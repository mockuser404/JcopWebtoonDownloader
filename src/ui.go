package main

import (
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
		MinSize:       Size{300, 200},
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
							WDform.LezhinComics.AccessToken = accesstokenLineEdit.Text()
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
