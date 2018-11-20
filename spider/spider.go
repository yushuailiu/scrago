package spider

import (
	"sync"
	"fmt"
	"net/http"
	"github.com/yushuailiu/scrago/pipeline"
	"github.com/yushuailiu/scrago/urlpool"
	"time"
	"github.com/yushuailiu/scrago/page_processer"
	"github.com/yushuailiu/scrago/downloader"
	"github.com/yushuailiu/scrago/request"
	"github.com/yushuailiu/scrago/resource_manage"
)

type Spider struct {
	sync.Mutex
	pageProcessor page_processer.PageProcessor

	freeDownloaderPool []*downloader.Downloader

	maxParallel int
	// 停止信号
	stopChannel chan struct{}

	// 是否运行中
	running bool

	tryTimes int

	pageSuccessFunc func(*http.Response) bool

	pipelines []pipeline.Pipeline

	urlPool urlpool.UrlPool

	parallelResource resource_manage.ResourceManage

	interval time.Duration

	header http.Header
}

const (
	DefaultMaxParallel = 1
	DefaultTryTimes    = 3
	DefaultInterval    = time.Second
)

func NewSpiderWithProcessor(pageProcessor page_processer.PageProcessor) *Spider {
	spider := &Spider{
		pageProcessor:    pageProcessor,
		tryTimes:         DefaultTryTimes,
		stopChannel:      make(chan struct{}),
		pageSuccessFunc:  defaultPageSuccessFunc,
		pipelines:        make([]pipeline.Pipeline, 0),
		urlPool:          urlpool.NewLocalQueueUrlPool(),
		parallelResource: resource_manage.NewResourceManageChan(DefaultMaxParallel),
		interval:         DefaultInterval,
		header: make(map[string][]string),
	}
	spider.setDefaultHeader()
	return spider
}


// defaultPageSuccessFunc 默认的判断页面是否成功方法
func defaultPageSuccessFunc(resp *http.Response) bool {
	return true
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
		select {
		case <-s.stopChannel:
			s.Close()
			break waitExit
		}
	}
	return s, nil
}

func (s *Spider) Close() {
	s.Lock()
	defer s.Unlock()

	s.running = false
}

func (s *Spider) startScrago() {
	for {
		req := s.urlPool.Pop()

		if req == nil && s.parallelResource.Has() == 0 {
			s.stopChannel <- struct{}{}
			return
		} else if req == nil {
			time.Sleep(time.Second * 1)
			continue
		}

		s.parallelResource.GetOne()

		d := downloader.NewDownloader()

		p := d.Do(req, s.tryTimes, s.header)
		s.pageProcessor.Process(req, p)

		time.Sleep(s.interval)

		for _, req := range p.GetNewRequests() {
			s.urlPool.Push(req)
		}

		for _, pipeline := range s.pipelines  {
			pipeline.Process(p)
		}

		s.parallelResource.FreeOne()
	}
}

func (s *Spider) AddUrl(method, url string) (*Spider, error) {
	req := request.NewRequest(method, url)
	s.Lock()
	defer s.Unlock()

	s.urlPool.Push(req)

	return s, nil
}

func (s *Spider) AddRequest(r *request.Request) *Spider {
	s.Lock()
	defer s.Unlock()
	s.urlPool.Push(r)
	return s
}

func (s *Spider) SetInterval(interval time.Duration) *Spider {
	s.interval = interval
	return s
}

func (s *Spider) SetParallelCount(parallelCount int) *Spider {
	s.parallelResource = resource_manage.NewResourceManageChan(parallelCount)
	return s
}

func (s *Spider) SetTryTimes(tryTimes int) *Spider {
	s.tryTimes = tryTimes
	return s
}

func (s *Spider) AddPipeline(p pipeline.Pipeline) *Spider {
	s.Lock()
	defer s.Unlock()
	s.pipelines = append(s.pipelines, p)
	return s
}

func (s *Spider) setDefaultHeader() {
	s.header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	s.header.Add("Accept-Encoding", "gzip, deflate")
	// todo  rand user-agent 还是保留作为默认值？
	s.header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.77 Safari/537.36")
}

func (s *Spider) AddHeader(key, val string) *Spider {
	s.header.Add(key, val)
	return s
}

func (s *Spider) SetHeader(key, val string) *Spider {
	s.header.Set(key, val)
	return s
}