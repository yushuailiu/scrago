package scrago

import (
	"net/http"
	"strings"
)

type Downloader struct {
	request chan *Request
	spider *Spider
}

func (d *Downloader) Run() {
	go func() {
		for request := range d.request {
			if request == nil {
				// todo 减少运行数量
				return
			}

			d.do(request)

			d.spider.addFreeDownloader(d)
		}
	}()
}

func (d *Downloader) do(request *Request) {
	c := &http.Client{}

	req, err := http.NewRequest(request.Method, request.Url, strings.NewReader(request.Body))

	if err != nil {
		// todo log
		return
	}

	resp, err := c.Do(req)
	if err != nil {
		// todo log
		return
	}
	p := &Page{}

	request.Page = p
	request.Resp = resp

	d.spider.pageProcessor.Process(request)
}