package page

import (
	"net/http"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"fmt"
	"github.com/yushuailiu/scrago/request"
	"github.com/bitly/go-simplejson"
)


type Page struct {
	status     string // e.g. "200 OK"
	statusCode int    // e.g. 200
	body string
	header http.Header
	cookies []*http.Cookie

	contentLength int64

	items []*pageItem

	newRequests []*request.Request
}

func (p *Page)GetNewRequests() []*request.Request {
	return  p.newRequests
}

func (p *Page) AddRequest(req *request.Request) {
	p.newRequests = append(p.newRequests, req)
}

func (p *Page) SetStatus(status string) *Page {
	p.status = status
	return p
}

func (p *Page) SetStatusCode(statusCode int) *Page {
	p.statusCode = statusCode
	return p
}

func (p *Page) SetBody(body string) *Page {
	p.body = body
	return p
}

func (p *Page) SetHeader(header http.Header) *Page {
	p.header = header
	return p
}

func (p *Page) SetCookies(cookies []*http.Cookie) *Page {
	p.cookies = cookies
	return p
}

func (p *Page) SetContentLength(contentLength int64) *Page {
	p.contentLength = contentLength
	return p
}


func (p *Page) GetBody() string {
	return p.body
}

func (p *Page) GetDocParse() (*goquery.Document, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(p.body))
	if err != nil {
		return nil, fmt.Errorf("doc parse error: %s", err)
	}
	return doc, nil
}

func (p *Page) GetJsonParse() (*simplejson.Json, error) {
	jsonData, err := simplejson.NewJson([]byte(p.body))
	return jsonData, err
}

func (p *Page) AddItem(item *pageItem) *Page {
	p.items = append(p.items, item)
	return p
}

func (p *Page) GetItems() []*pageItem {
	return p.items
}