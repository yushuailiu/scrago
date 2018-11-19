package pipeline

import (
	"github.com/yushuailiu/scrago/page"
	"fmt"
)

type ConsolePipeline struct {
}

func (c *ConsolePipeline) Process(p *page.Page)  {
	for _, item := range p.GetItems() {
		fmt.Println(item)
	}
}
