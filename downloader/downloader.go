package downloader

import (
	"net/http"
	"strings"
	"io/ioutil"
	"compress/gzip"
	"github.com/yushuailiu/scrago/request"
	"github.com/yushuailiu/scrago/page"
	"fmt"
)

type Downloader struct {
	request chan *request.Request
}

func NewDownloader() *Downloader {
	return &Downloader{
		request: make(chan *request.Request),
	}
}

func (d *Downloader) Do(request *request.Request, tryTimes int, header http.Header) *page.Page {

	c := &http.Client{}

	req, err := http.NewRequest(request.Method, request.Url, strings.NewReader(request.Body))

	for key, vals := range header {
		for _, val := range vals {
			req.Header.Add(key, val)
		}
	}

	for key, vals := range request.Header {
		for _, val := range vals {
			req.Header.Add(key, val)
		}
	}

	fmt.Println(req.Header)
	fmt.Println(req.URL)

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
