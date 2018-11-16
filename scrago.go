package scrago

import (
	"github.com/yushuailiu/scrago/spider"
	"github.com/yushuailiu/scrago/page_processer"
	"os"
	"log"
)

var (
	Logger        = defaultLogger()
)

func NewSpider() *spider.Spider {
	return spider.NewSpiderWithProcessor(nil)
}

func NewSpiderWithProcessor(pageProcessor page_processer.PageProcessor) *spider.Spider {
	return spider.NewSpiderWithProcessor(pageProcessor)
}

func defaultLogger() *log.Logger {
	return log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
}