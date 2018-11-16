package scrago

type Request struct {
	Url    string
	Method string
	Body   string
	Page *Page
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