package scrago

type PageItem struct {
	fields map[string]interface{}
}

func (p *PageItem) AddField(field string, val interface{}) {
	p.fields[field] = val
}