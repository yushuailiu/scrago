package scrago

import (
	"sync"
	"fmt"
	"sync/atomic"
	"log"
	"os"
)

type Spider struct {
	sync.Mutex
	cond          *sync.Cond
	requests      []*Request
	pageProcessor PageProcessor

	freeDownloaderPool []*Downloader

	maxParallel int
	// 停止信号
	stopChannel chan struct{}

	// 是否运行中
	running bool

	// 正在抓取数量
	runningCount int32

	logger *log.Logger
}

var (
	SpiderStopErr = fmt.Errorf("spider has stoped")
)

const (
	DefaultMaxParallel = 5
)

func NewSpider() *Spider {
	return NewSpiderWithProcessor(nil)
}

func (s *Spider) Run() (*Spider, error) {
	if s.pageProcessor == nil {
		return s, fmt.Errorf("please set page processor")
	}

	s.Lock()
	s.running = true
	s.Unlock()

	go s.startScrago()

waitExit:
	for {
		s.logger.Println("for")
		select {
		case <-s.stopChannel:
			s.Close()
			s.logger.Println("exit")
			break waitExit
		}
	}
	return s, nil
}

func (s *Spider) Close() {
	s.Lock()
	defer s.Unlock()

	for _, d := range s.freeDownloaderPool {
		d.request <- nil
	}

	s.logger.Println("close")

	s.freeDownloaderPool = nil
	s.running = false
	s.cond.Signal()
}

func (s *Spider) startScrago() {
	for {
		s.Lock()

		if s.running == false {
			s.Unlock()
			return
		}

		requests := s.requests
		if len(requests) == 0 {
			s.Unlock()
			continue
		}

		request := requests[0]
		s.requests = requests[1:]

		wait := true
		if int(atomic.LoadInt32(&s.runningCount)) >= s.maxParallel {
			wait = true
		} else {
			wait = false
		}

		if wait {
			for {
				s.cond.Wait()
				if !s.running {
					s.Unlock()
					return
				}
				frees := s.freeDownloaderPool
				num := len(frees)
				if num == 0 {
					continue
				}
				download := frees[0]
				s.freeDownloaderPool = frees[1:]
				download.request <- request
				break
			}
		} else {
			d := &Downloader{
				request: make(chan *Request),
				spider:  s,
			}
			d.Run()
			d.request <- request
			atomic.AddInt32(&s.runningCount, 1)
		}
		s.Unlock()
	}
}

func NewSpiderWithProcessor(pageProcessor PageProcessor) *Spider {
	spider := &Spider{
		pageProcessor: pageProcessor,
		maxParallel:   DefaultMaxParallel,
		logger:        defaultLogger(),
		stopChannel:   make(chan struct{}),
	}
	spider.cond = sync.NewCond(spider)
	return spider
}

func (s *Spider) AddUrl(method, url string) (*Spider, error) {
	request := NewGetRequest(url)
	s.Lock()
	defer s.Unlock()
	s.requests = append(s.requests, request)
	return s, nil
}

func (s *Spider) addFreeDownloader(d *Downloader) error {
	s.Lock()
	defer s.Unlock()
	if !s.running {
		return SpiderStopErr
	}

	atomic.AddInt32(&s.runningCount, -1)

	s.freeDownloaderPool = append(s.freeDownloaderPool, d)
	s.cond.Signal()

	s.logger.Println(s.runningCount)
	s.logger.Println(len(s.requests))
	if s.runningCount == 0 && len(s.requests) == 0 {
		s.logger.Println("fadsf")
		s.stopChannel <- struct{}{}
		s.logger.Println("fffff")
	}

	return nil
}

func defaultLogger() *log.Logger {
	return log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
}
