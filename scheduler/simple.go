package scheduler

import "NewsSpider/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

// 每个work 都有自己的channel
func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

func (s *SimpleScheduler) WorkerReady(chan engine.Request) {

}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	// 每个request 建立一个goroutine
	go func() { s.workerChan <- r }()
}
