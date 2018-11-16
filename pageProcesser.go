package scrago

type PageProcessor interface {
	Process(request *Request)
}