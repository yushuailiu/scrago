package urlpool

import "github.com/yushuailiu/scrago/request"

type UrlPool interface {
	Push(req *request.Request)
	Pop() *request.Request
}
