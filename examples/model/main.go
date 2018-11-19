// 抓取json 页面示例
package main

import (
	"github.com/yushuailiu/scrago/request"
	"github.com/yushuailiu/scrago/page"
	"github.com/yushuailiu/scrago"
	json2 "encoding/json"
	"github.com/yushuailiu/scrago/pipeline"
	"strconv"
)

func StringToInt(str string) int {
	if res, err := strconv.Atoi(str); err == nil {
		return res
	}
	return 0
}

type DBPipeline struct {
}

func (d *DBPipeline) Process(p *page.Page) {
	items := p.GetItems()
	for _, item := range items {
		user := User{
			UserName: item.GetField("user_name").(string),
			Avatar: item.GetField("avatar_url").(string),
			Intro: item.GetField("intro").(string),
			FollowCount: StringToInt(item.GetField("follow_count").(string)),
			FansCount: StringToInt(item.GetField("fans_count").(string)),
			UserType: int8(StringToInt(item.GetField("user_type").(string))),
			IsVip: int8(StringToInt(item.GetField("is_vip").(string))),
			PubshareCount: StringToInt(item.GetField("pubshare_count").(string)),
			Type: int8(StringToInt(item.GetField("type").(string))),
			AlbumCount: StringToInt(item.GetField("album_count").(string)),

		}
		hotUk, err := strconv.ParseInt(item.GetField("hot_uk").(string), 10, 64)
		if err != nil {
			user.HotUk = hotUk
		}
		count := 0
		DB.Table("user").Where("avatar = ?", user.Avatar).Count(&count)
		if count != 0 {
			continue
		}

		DB.Create(&user)
	}
}


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
		item.AddField("album_count", string(tempUser["album_count"].(json2.Number)))
		p.AddItem(item)
	}

}

func main() {
	spider := scrago.NewSpiderWithProcessor(&SimplePageProcessor{})
	spider.AddUrl("GET", "http://yun.baidu.com/pcloud/friend/gethotuserlist?type=1&from=feed&start=0&limit=24&channel=chunlei&clienttype=0&web=1")
	spider.AddPipeline(&pipeline.ConsolePipeline{})
	spider.AddPipeline(&DBPipeline{})
	spider.Run()
}
