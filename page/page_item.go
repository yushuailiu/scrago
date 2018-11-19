package page

import (
	"github.com/kataras/iris/core/errors"
)


var (
	FieldNotFount = errors.New("field not found")
)

type pageItem struct {
	fields map[string]interface{}
}

func NewPageItem() *pageItem {
	return &pageItem{
		fields: make(map[string]interface{}),
	}
}
func (p *pageItem) AddField(field string, val interface{}) {
	p.fields[field] = val
}

func (p *pageItem) GetField(field string) interface{} {
	return p.fields[field]
}