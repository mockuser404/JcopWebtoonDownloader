package WTdown

// import (
// 	"errors"
// 	"net/http"
// 	"os"
// 	"strconv"

// 	"github.com/jiajunhuang/gotasks/pool"

// 	"github.com/PuerkitoBio/goquery"
// )

// type Webtooncom struct {
// 	TitleId string
// 	Cookies string

// 	epis        []string
// 	EpisodeName []string
// }

// const (
// 	// NAVER_BASE_EPISODES_DATA_URL = "https://m.comic.naver.com/webtoon/list.nhn?sortOrder=ASC&titleId="
// 	// NAVER_BASE_IMG_DATA_URL      = "https://m.comic.naver.com/webtoon/detail.nhn?titleId="
// )

// // error code, error
// // 0:nil, 1:error, 2:warning, 3:info
// func (nc *Webtooncom) Download(start, stop, thread int, folder string) (int, error) {
// 	for episode := start; episode <= stop; episode++ {

// 		var dataURL string

// 		header := make(map[string]string, 2)
// 		header["User-Agent"] = USER_AGENT
// 		header["Cookie"] = nc.Cookies
// 		resp, err := requestWithCookieNBody(NAVER_BASE_IMG_DATA_URL+nc.TitleId+"&no="+nc.epis[episode], "GET", header, nil)
// 		if err != nil {
// 			return 1, err
// 		}

// 		doc, err := goquery.NewDocumentFromReader(resp)
// 		if err != nil {
// 			return 1, err
// 		}

// 		NumImg := doc.Find(".lazy").Length()
// 		if NumImg <= 0 {
// 			return 2, errors.New("Epi " + strconv.Itoa(episode+1) + " - Can't find Images")
// 		}
// 		err = os.MkdirAll(folder+"/"+strconv.Itoa(episode+1), os.ModePerm)
// 		if err != nil {
// 			return 1, errors.New("Epi " + strconv.Itoa(episode+1) + " - Can't make folder")
// 		}
// 		// errchan := make(chan error, NumImg)
// 		gopool := pool.NewGoPool(pool.WithMaxLimit(thread))

// 		doc.Find(".lazy").Each(func(j int, s *goquery.Selection) {
// 			dataURL, _ = s.Attr("data-src")
// 			func(dataURL string) {
// 				gopool.Submit(func() {
// 					downloadFileSingle(dataURL, folder+"/"+strconv.Itoa(episode+1)+"/"+strconv.Itoa(j+1)+".jpg")
// 				})
// 			}(dataURL)
// 			// go downloadFile(string(dataURL), folder+"/"+strconv.Itoa(episode)+"/"+strconv.Itoa(j+1)+".jpg", errchan)
// 		})

// 		gopool.Wait()
// 		// for i := 0; i < NumImg; i++ {
// 		// 	err = <-errchan
// 		// 	if err != nil {
// 		// 		return 1, err
// 		// 	}
// 		// }
// 		if err != makeHTML(episode+1, NumImg, nc.TitleId, folder+"/"+strconv.Itoa(episode+1)+"/"+strconv.Itoa(episode+1)+".html") {
// 			return 1, errors.New("Epi " + strconv.Itoa(episode+1) + " - Can't make HTML viewer")
// 		}
// 	}

// 	return 0, nil
// }

// func (nc *NaverComic) GetEpiData() error {
// 	nc.epis = make([]string, 0)
// 	nc.EpisodeName = make([]string, 0)

// 	err, totalPages := nc.getTotalPages()
// 	if err != nil {
// 		return err
// 	}
// 	intTotalPages, err := strconv.Atoi(totalPages)

// 	for pageNum := 1; pageNum <= intTotalPages; pageNum++ {
// 		nc.getDataFromEachPage(pageNum)
// 	}
// 	return nil
// }

// func (nc *NaverComic) getTotalPages() (error, string) {
// 	req, err := http.NewRequest("GET", NAVER_BASE_EPISODES_DATA_URL+nc.TitleId, nil)
// 	if err != nil {
// 		return err, ""
// 	}

// 	req.Header.Set("User-Agent", USER_AGENT)
// 	req.Header.Set("Cookie", nc.Cookies)
// 	resp, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		return err, ""
// 	}
// 	defer resp.Body.Close()

// 	// resp, err := requestWithCookieNBody(BASE_EPISODES_DATA_URL, nil)

// 	doc, err := goquery.NewDocumentFromReader(resp.Body)
// 	if err != nil {
// 		return err, ""
// 	}
// 	totalPages := doc.Find(".total").Text()

// 	return nil, totalPages
// }

// func (nc *NaverComic) getDataFromEachPage(page int) error {
// 	req, err := http.NewRequest("GET", NAVER_BASE_EPISODES_DATA_URL+nc.TitleId+"&page="+strconv.Itoa(page), nil)
// 	if err != nil {
// 		return err
// 	}

// 	req.Header.Set("User-Agent", USER_AGENT)
// 	req.Header.Set("Cookie", nc.Cookies)
// 	resp, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		return err
// 	}
// 	defer resp.Body.Close()

// 	doc, err := goquery.NewDocumentFromReader(resp.Body)
// 	if err != nil {
// 		return err
// 	}
// 	doc.Find(".item").Each(func(_ int, s *goquery.Selection) {
// 		_, state := s.Attr("data-title-id")
// 		if state {
// 			strnumEpi, _ := s.Attr("data-no")
// 			numEpi, _ := strconv.Atoi(strnumEpi)
// 			// if err != nil {
// 			// 	log.Println("data-no is not a int")
// 			// }
// 			nc.epis = append(nc.epis, strconv.Itoa(numEpi))
// 			nc.EpisodeName = append(nc.EpisodeName, s.Find(".name").Text())
// 		}
// 	})
// 	return err
// }