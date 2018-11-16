package resource_manage

type resourceManageChan struct {
	cap int
	ch  chan int
}

func NewResourceManageChan(cap int) *resourceManageChan {
	return &resourceManageChan{
		cap: cap,
		ch:  make(chan int, cap),
	}
}

func (r *resourceManageChan) GetOne() {
	r.ch <- 1
}

func (r *resourceManageChan) FreeOne() {
	<- r.ch
}

func (r *resourceManageChan) Left() int {
	return r.cap - int(len(r.ch))
}

func (r *resourceManageChan) Has() int {
	return len(r.ch)
}