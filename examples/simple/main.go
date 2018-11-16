package main

import (
	"github.com/yushuailiu/scrago"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"github.com/yushuailiu/scrago/request"
	"github.com/yushuailiu/scrago/page"
)

type SimplePageProcessor struct {
}

func (s *SimplePageProcessor) Process(req *request.Request, p *page.Page) {
	scrago.Logger.Println(p)
	doc, err := p.GetDocParse()
	if err != nil {
		return
	}

	doc.Find("#ip_list tr").Each(func(i int, selection *goquery.Selection) {
		if selection.Find("th").Length() == 0 {
			item := page.NewPageItem()

			tds := selection.Find("td")
			if country, ok := tds.Eq(0).Find("img").First().Attr("alt"); ok {
				item.AddField("country", country)
			}
			item.AddField("ip", tds.Eq(1).Text())
			item.AddField("port", tds.Eq(2).Text())
			item.AddField("city", strings.TrimSpace(tds.Eq(3).Text()))
			item.AddField("hide", tds.Eq(4).Text())
			item.AddField("protocol", tds.Eq(5).Text())

			p.AddItem(item)
		}
	})

	doc.Find(".pagination a[href]").Each(func(i int, selection *goquery.Selection) {
		href, _ := selection.Attr("href")
		newReq := request.NewGetRequest("http://www.xicidaili.com/" + href)
		scrago.Logger.Println(newReq.Url)
		p.AddRequest(newReq)
	})
}

func main() {
	spider := scrago.NewSpiderWithProcessor(&SimplePageProcessor{})
	_, err := spider.AddUrl("GET", "http://www.xicidaili.com/nn/")

	if err != nil {
		fmt.Println(err)
	}

	spider.Run()
}