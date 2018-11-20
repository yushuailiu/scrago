package request

import (
	"crypto/md5"
	"fmt"
	"net/http"
)

type Request struct {
	Url    string
	Method string
	Body   string
	Header http.Header
}

func NewRequest(method, url string) *Request {
	return &Request{
		Method: method,
		Url:    url,
		Header: make(map[string][]string),
	}
}

func NewGetRequest(url string) *Request {
	return NewRequest("GET", url)
}

func NewPostRequest(url string) *Request {
	return NewRequest("POST", url)
}

func (r *Request) SetBody(body string) *Request {
	r.Body = body
	return r
}

func (r *Request) Key() string {
	key := md5.Sum([]byte(r.Method + " " + r.Url + " " + r.Body))

	return fmt.Sprintf("%x", key)
}