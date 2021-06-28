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
		MinSize:       Size{Width:300, Height:200},
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

func lezhinRunDialog(owner walk.Form,cookieType *string, requireCookies []string) (int, error) {
	var dlg *walk.Dialog
	accesstokenLineEdit := make([]*walk.LineEdit, len(requireCookies)+1)
	// requireCookies = append(requireCookies, "access_token")
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

	return Dialog{
		AssignTo:      &dlg,
		Title:         "Set Lezhin access token",
		DefaultButton: &acceptPB,
		CancelButton:  &cancelPB,
		MinSize:       Size{Width:300, Height:200},
		Layout:        VBox{},
		Children: []Widget{
			Composite{
				Layout: Grid{Columns: 2},
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
