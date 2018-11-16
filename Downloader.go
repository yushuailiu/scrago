package scrago

import (
	"net/http"
	"strings"
	"sync/atomic"
	"io/ioutil"
)

type Downloader struct {
	request chan *Request
	spider  *Spider
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

	tryTimes := int(atomic.LoadInt32(&d.spider.tryTimes))
	var resp *http.Response
	// todo 这里是否添加 resp 状态的检查？或者添加自定义的判断是否成功得方法？
	for i := 0; i < tryTimes; i++ {
		resp, err = c.Do(req)
		if err == nil && d.spider.pageSuccessFunc(resp) {
			break
		}
	}

	if err != nil {
		// todo log
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		// todo log
		return
	}
	p := &Page{
		status:     resp.Status,
		statusCode: resp.StatusCode,
		body:       string(body),
		header:     resp.Header,
		cookies:    resp.Cookies(),
	}

	d.spider.pageProcessor.Process(request, p)
}
