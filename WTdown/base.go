package WTdown

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

var (
	err error
)

const (
	USER_AGENT = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/50.0.2661.102 Safari/537.36"
)

func requestWithCookieNBody(urlname, method string, header map[string]string, body map[string]string) (io.Reader, error) {
	data := &url.Values{}
	for i := range body {
		data.Add(i, body[i])
	}

	req, err := http.NewRequest(method, urlname, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return nil, err
	}
	// req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	// req.Header.Add("User-Agent", USER_AGENT)

	for i := range header {
		req.Header.Add(i, header[i])
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}

func downloadFile(URL, fileName string, errchan chan error) {
	out, err := os.Create(fileName)
	if err != nil {
		errchan <- err
	}
	defer out.Close()

	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		errchan <- err
	}

	req.Header.Set("User-Agent", USER_AGENT)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		errchan <- err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		errchan <- err
	}

	errchan <- nil
}

func downloadFileSingle(URL, fileName string) error {

	out, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer out.Close()

	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return err
	}

	req.Header.Set("User-Agent", USER_AGENT)
	resp, err := http.DefaultClient.Do(req)
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

func makeHTML(episode, NumImg int, titleId, outFilePath string) error {
	content := "<html><head><meta charset='UTF-8'><meta name='viewport' content='width=device-width, initial-scale=1.0'><style>body, html{margin: 0;border: 0;padding: 0;}@media only screen and (max-width: 700px) {img {width: 100%;}}</style><title>Episode " + strconv.Itoa(episode) + " (" + titleId + ")</title></head><body><center>"
	for l := 1; l <= NumImg; l++ {
		content += "<img src='"
		content += strconv.Itoa(l)
		content += ".jpg'><br>"
	}
	content += "</body></center></html>"

	err = ioutil.WriteFile(outFilePath, []byte(content), 0)
	if err != nil {
		return err
	}
	return nil
}

func reverseStrArray(arr *[]string){
	for i, j := 0, len(*arr)-1; i < j; i, j = i+1, j-1 {
		(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
	}
}