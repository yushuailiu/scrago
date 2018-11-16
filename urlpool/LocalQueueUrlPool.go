package urlpool

import (
	"container/list"
	"sync"
	"github.com/yushuailiu/scrago/request"
)

type LocalQueueUrlPool struct {
	rmRepeat bool

	historyUrls map[string]struct{}
	list *list.List
	sync.Mutex
}

func NewLocalQueueUrlPool() UrlPool {
	return &LocalQueueUrlPool{
		rmRepeat: true,
		historyUrls: make(map[string]struct{}),
		list: list.New(),
	}
}

func (l *LocalQueueUrlPool) Push(req *request.Request) {
	l.Lock()
	defer l.Unlock()

	if _, ok := l.historyUrls[req.Key()]; ok && l.rmRepeat {
		return
	}

	l.historyUrls[req.Key()] = struct{}{}

	l.list.PushBack(req)

	return
}

func (l *LocalQueueUrlPool) Pop() *request.Request {
	l.Lock()
	defer l.Unlock()

	if l.list.Len() == 0 {
		return nil
	}

	element := l.list.Front()
	req := element.Value.(*request.Request)

	l.list.Remove(element)
	return req
}