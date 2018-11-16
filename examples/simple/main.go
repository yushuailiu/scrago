package main

import (
	"github.com/yushuailiu/scrago"
	"fmt"
)

type SimplePageProcessor struct {
}

func (s *SimplePageProcessor) Process(request *scrago.Request, page *scrago.Page) {
	item := &scrago.PageItem{}

	doc, err := page.GetDocParse()
	if err != nil {
		return
	}

}

func main() {
	spider := scrago.NewSpiderWithProcessor(&SimplePageProcessor{})
	_, err := spider.AddUrl("GET", "https://weibo.com/?category=1760")

	if err != nil {
		fmt.Println(err)
	}

	spider.Run()
}