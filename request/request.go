package request

import (
	"crypto/md5"
	"fmt"
)

type Request struct {
	Url    string
	Method string
	Body   string
}

func NewRequest(method, url string) *Request {
	return &Request{
		Method: method,
		Url:    url,
	}
}

func NewGetRequest(url string) *Request {
	return &Request{
		Method: "GET",
		Url:    url,
	}
}

func NewPostRequest(url string) *Request {
	return &Request{
		Method: "POST",
		Url:    url,
	}
}

func (r *Request) SetBody(body string) *Request {
	r.Body = body
	return r
}

func (r *Request) Key() string {
	key := md5.Sum([]byte(r.Method + " " + r.Url + " " + r.Body))

	return fmt.Sprintf("%x", key)
}