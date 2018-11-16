package scrago

import (
	"net/http"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"fmt"
)


type Page struct {
	status     string // e.g. "200 OK"
	statusCode int    // e.g. 200
	body string
	header http.Header
	cookies []*http.Cookie

	contentLength int

	items []*PageItem
}

func (p *Page) GetDocParse() (*goquery.Document, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(p.body))
	if err != nil {
		return nil, fmt.Errorf("doc parse error: %s", err)
	}
	return doc, nil
}