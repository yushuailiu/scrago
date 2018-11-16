package page_processer

import (
	"github.com/yushuailiu/scrago/request"
	"github.com/yushuailiu/scrago/page"
)

type PageProcessor interface {
	Process(request *request.Request, page *page.Page)
}