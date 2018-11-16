package main

import (
	"github.com/yushuailiu/scrago"
	"fmt"
)

type SimplePageProcessor struct {
}

func (s *SimplePageProcessor) Process(request *scrago.Request) {
	fmt.Println(request.Resp.Status)
}

func main() {
	spider := scrago.NewSpiderWithProcessor(&SimplePageProcessor{})
	_, err := spider.AddUrl("GET", "http://yun.baidu.com/pcloud/friend/gethotuserlist?type=1&from=feed&start=0&limit=24&channel=chunlei&clienttype=0&web=1")
	spider.AddUrl("GET", "http://yun.baidu.com/pcloud/friend/gethotuserlist?type=1&from=feed&start=0&limit=24&channel=chunlei&clienttype=0&web=1")
	spider.AddUrl("GET", "http://yun.baidu.com/pcloud/friend/gethotuserlist?type=1&from=feed&start=0&limit=24&channel=chunlei&clienttype=0&web=1")

	if err != nil {
		fmt.Println(err)
	}

	spider.Run()
}