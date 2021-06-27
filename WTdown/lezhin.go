package WTdown

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/buger/jsonparser"
	"github.com/jiajunhuang/gotasks/pool"
)

type LezhinComics struct {
	TitleId string
	Cookies string // _kpawlt _kpawltea _kpawlst
	AccessToken string

	epis        []string
	EpisodeName []string
}

const (
	LEZHIN_SINGLE_URL   = "https://www.lezhin.com/ko/comic/"
	LEZHIN_BASE_IMG_API = "https://www.lezhin.com/api/v2/inventory_groups/comic_viewer_k"
	LEZHIN_BASE_IMG_URL = "https://cdn.lezhin.com/v2"
)

func (lc *LezhinComics) Download(start, stop, thread int, folder string) (int, error) {
	for episode := start; episode <= stop; episode++ {

		reqHeader := make(map[string]string, 2)
		reqHeader["x-lz-locale"] = "ko_KR"
		if len(lc.epis) < episode {
			return 2, errors.New("Can't find Epi " + strconv.Itoa(episode+1))
		}
		resp, err := requestWithCookieNBody(LEZHIN_BASE_IMG_API+"?alias="+lc.TitleId+"&name="+lc.epis[episode]+"&type=comic_episode", "GET", reqHeader, make(map[string]string, 0))
		if err != nil {
			return 1, err
		}
		downloadData, err := ioutil.ReadAll(resp)
		if err != nil {
			return 1, err
		}
		imgs := make([]string, 0)

		type eachFiles struct {
			Path string `json:"path"`
		}

		var outputFiles []eachFiles
		files, err := jsonparser.GetUnsafeString(downloadData, "data", "extra", "episode", "scrollsInfo")
		json.Unmarshal([]byte(files), &outputFiles)
		for i := range outputFiles {
			imgs = append(imgs, outputFiles[i].Path)
		}

		if len(imgs) <= 0 {
			return 2, errors.New("Can't find Images")
		}

		err = os.MkdirAll(folder+"/"+strconv.Itoa(episode+1), os.ModePerm)
		if err != nil {
			return 1, err
		}
		// errchan := make(chan error, len(imgs))
		gopool := pool.NewGoPool(pool.WithMaxLimit(thread))
		for anchor := range imgs {
			func(anchor int) {
				gopool.Submit(func() {
					downloadFileSingle(LEZHIN_BASE_IMG_URL+imgs[anchor]+"?access_token="+lc.AccessToken, folder+"/"+strconv.Itoa(episode+1)+"/"+strconv.Itoa(anchor+1)+".jpg")
				})
			}(anchor)
		}
		gopool.Wait()
		// for i := 0; i < len(imgs); i++ {
		// 	err = <-errchan
		// 	if err != nil {
		// 		return 1, err
		// 	}
		// }

		makeHTML(episode+1, len(imgs), lc.TitleId, folder+"/"+strconv.Itoa(episode+1)+"/"+strconv.Itoa(episode+1)+".html")
	}
	return 0, nil
}

func (lc *LezhinComics) GetEpiData() error {
	lc.epis = make([]string, 0)
	lc.EpisodeName = make([]string, 0)

	req, err := http.NewRequest("GET", LEZHIN_SINGLE_URL+lc.TitleId, nil)
	// req, err := http.NewRequest("GET", "https://www.lezhin.com/ko/comic/girlwetwall", nil)
	if err != nil {
		return err
	}

	// req.Header.Set("User-Agent", USER_AGENT)
	req.Header.Set("Cookie", lc.Cookies)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	source, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	first := strings.Index(string(source), "all: [") + 5
	last := strings.Index(string(source[first:]), "\n")

	type eachEpi struct {
		Name    string            `json:"name"`
		Display map[string]string `json:"display"`
	}
	var titleNepi []eachEpi
	err = json.Unmarshal(source[first:first+last-1], &titleNepi)
	if err != nil {
		fmt.Println("error:", err)
	}
	for each := range titleNepi {
		lc.epis = append(lc.epis, titleNepi[each].Name)
		lc.EpisodeName = append(lc.EpisodeName, titleNepi[each].Display["title"])
	}
	reverseStrArray(&lc.epis)
	reverseStrArray(&lc.EpisodeName)
	return nil
}
