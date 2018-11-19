package pipeline

import "github.com/yushuailiu/scrago/page"

type Pipeline interface {
	Process(page *page.Page)
}
