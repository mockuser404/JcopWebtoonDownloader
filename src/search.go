package main

import (
	"errors"
	"net/url"
	"strings"

	"github.com/lxn/walk"
)

type EnvModel struct {
	walk.ListModelBase
	items []string
}

func (m *EnvModel) ItemCount() int {
	return len(m.items)
}

func (m *EnvModel) Value(index int) interface{} {
	return m.items[index]
}

func LoadEpis() {
	buff, err := url.Parse(WDdata.TitleIdURL.Text())
	if err != nil {
		Log(2, err)
	}

	values, err := url.ParseQuery(buff.RawQuery)
	if err != nil {
		Log(2, err)
	}
	switch WDdata.Type.CurrentIndex() {
	case -1:
		Log(2, errors.New("Please Select Website"))
	case 0:
		if values.Get("titleId") != "" {

			WDform.NaverComic.TitleId = values.Get("titleId")
			go func() {
				LoadingOn()
				err = WDform.NaverComic.GetEpiData()
				if err != nil {
					Log(2, err)
				}
				WDdata.StartControl.SetModel(&EnvModel{items: WDform.NaverComic.EpisodeName})
				WDdata.StopControl.SetModel(&EnvModel{items: WDform.NaverComic.EpisodeName})
				LoadingOff()
			}()
		} else {
			Log(2, errors.New("Can't find id from URL"))
		}
	case 1:
		if values.Get("seriesId") != "" {

			WDform.KakaoPage.TitleId = values.Get("seriesId")
			go func() {
				LoadingOn()
				err = WDform.KakaoPage.GetEpiData()
				if err != nil {
					Log(2, err)
				}
				WDdata.StartControl.SetModel(&EnvModel{items: WDform.KakaoPage.EpisodeName})
				WDdata.StopControl.SetModel(&EnvModel{items: WDform.KakaoPage.EpisodeName})
				LoadingOff()
			}()
		} else {
			Log(2, errors.New("Can't find id from URL"))
		}
	case 2:
		WDform.DaumWebtoon.TitleId = strings.Split(buff.Path, "/")[len(strings.Split(buff.Path, "/"))-1]
		go func() {
			LoadingOn()
			err = WDform.DaumWebtoon.GetEpiData()
			if err != nil {
				Log(2, err)
			}
			WDdata.StartControl.SetModel(&EnvModel{items: WDform.DaumWebtoon.EpisodeName})
			WDdata.StopControl.SetModel(&EnvModel{items: WDform.DaumWebtoon.EpisodeName})
			LoadingOff()
		}()
	case 3:
		WDform.LezhinComics.TitleId = strings.Split(buff.Path, "/")[len(strings.Split(buff.Path, "/"))-1]
		go func() {
			LoadingOn()
			err = WDform.LezhinComics.GetEpiData()
			if err != nil {
				Log(2, err)
			}
			WDdata.StartControl.SetModel(&EnvModel{items: WDform.LezhinComics.EpisodeName})
			WDdata.StopControl.SetModel(&EnvModel{items: WDform.LezhinComics.EpisodeName})
			LoadingOff()
		}()
	case 4:
		if values.Get("seriesId") != "" {

			WDform.KPepub.TitleId = values.Get("seriesId")
			go func() {
				LoadingOn()
				err = WDform.KPepub.GetEpiData()
				if err != nil {
					Log(2, err)
				}
				WDdata.StartControl.SetModel(&EnvModel{items: WDform.KPepub.EpisodeName})
				WDdata.StopControl.SetModel(&EnvModel{items: WDform.KPepub.EpisodeName})
				LoadingOff()
			}()
		} else {
			Log(2, errors.New("Can't find id from URL"))
		}
	}

	// m := &EnvModel{items: make([]string, 2)}
	// m.items[0] = "1"
	// m.items[1] = "3"
	// WDform.StartControl.SetModel(m)
	// log.Println(WDform.Start)
}
