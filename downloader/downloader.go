package downloader

import (
	"net/http"
	"strings"
	"io/ioutil"
	"compress/gzip"
	"github.com/yushuailiu/scrago/request"
	"github.com/yushuailiu/scrago/page"
)

type Downloader struct {
	request chan *request.Request
}

func NewDownloader() *Downloader {
	return &Downloader{
		request: make(chan *request.Request),
	}
}

func (d *Downloader) Do(request *request.Request, tryTimes int) *page.Page {

	c := &http.Client{}

	req, err := http.NewRequest(request.Method, request.Url, strings.NewReader(request.Body))

	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.77 Safari/537.36")


	if err != nil {
		// todo log
		return nil
	}

	var resp *http.Response
	// todo 这里是否添加 resp 状态的检查？或者添加自定义的判断是否成功得方法？
	for i := 0; i < tryTimes; i++ {
		resp, err = c.Do(req)
		if err == nil {
			break
		}
	}

	if err != nil {
		// todo log
		return nil
	}

	defer resp.Body.Close()

	var body string


	if resp.Header.Get("Content-Encoding") == "gzip" {
		compressedReader, err := gzip.NewReader(resp.Body)
		byteBody,err := ioutil.ReadAll(compressedReader)
		body = string(byteBody)
		if err != nil {
			// todo
			return nil
		}
	} else {
		byteBody,err := ioutil.ReadAll(resp.Body)
		body = string(byteBody)

		if err != nil {
			// todo
			return nil
		}
	}

	if err != nil {
		// todo log
		return nil
	}
	p := &page.Page{}
	p.SetStatus(resp.Status).SetStatusCode(resp.StatusCode).
		SetBody(body).SetHeader(resp.Header).SetCookies(resp.Cookies()).SetContentLength(resp.ContentLength)
	return p
}
