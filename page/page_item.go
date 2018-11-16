package page

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
