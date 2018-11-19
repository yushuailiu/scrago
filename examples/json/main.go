// 抓取json 页面示例
package main

import (
	"github.com/yushuailiu/scrago/request"
	"github.com/yushuailiu/scrago/page"
	"github.com/yushuailiu/scrago"
	json2 "encoding/json"
	"github.com/yushuailiu/scrago/pipeline"
)

type SimplePageProcessor struct {
}

func (s *SimplePageProcessor) Process(req *request.Request, p *page.Page) {
	json, err := p.GetJsonParse()
	if err != nil {
		scrago.Logger.Println(err)
		return
	}

	users, err := json.Get("hotuser_list").Array()

	if err != nil {
		scrago.Logger.Println(err)
		return
	}

	for _, user := range users {
		tempUser := user.(map[string]interface{})
		item := page.NewPageItem()
		item.AddField("user_name", tempUser["hot_uname"].(string))
		item.AddField("intro", tempUser["intro"].(string))
		item.AddField("avatar_url", tempUser["avatar_url"].(string))
		item.AddField("follow_count", string(tempUser["follow_count"].(json2.Number)))
		item.AddField("fans_count", string(tempUser["fans_count"].(json2.Number)))
		item.AddField("user_type", string(tempUser["user_type"].(json2.Number)))
		item.AddField("is_vip", string(tempUser["is_vip"].(json2.Number)))
		item.AddField("hot_uk", string(tempUser["hot_uk"].(json2.Number)))
		item.AddField("pubshare_count", string(tempUser["pubshare_count"].(json2.Number)))
		item.AddField("type", string(tempUser["type"].(json2.Number)))
		p.AddItem(item)
	}

}

func main() {
	spider := scrago.NewSpiderWithProcessor(&SimplePageProcessor{})
	spider.AddUrl("GET", "http://yun.baidu.com/pcloud/friend/gethotuserlist?type=1&from=feed&start=0&limit=24&channel=chunlei&clienttype=0&web=1")
	spider.AddPipeline(&pipeline.ConsolePipeline{})
	spider.Run()
}
