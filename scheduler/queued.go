package scheduler

import "NewsSpider/engine"

type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

// 每个work 都有自己的channel
func (s *QueuedScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (s *QueuedScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}

func (s *QueuedScheduler) WorkerReady(w chan engine.Request) {
	s.workerChan <- w
}

func (s *QueuedScheduler) Run() {
	s.workerChan = make(chan chan engine.Request)
	s.requestChan = make(chan engine.Request)

	go func() {
		// 创建2个队列
		var requestQ []engine.Request
		var workerQ []chan engine.Request

		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			// 当 requestQ 队列大于0, workerQ 队列大于0 ，分别在2个队列中取出一个赋值给channel
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeWorker = workerQ[0]
				activeRequest = requestQ[0]
			}

			select {
			case r := <-s.requestChan:
				requestQ = append(requestQ, r)
			case w := <-s.workerChan:
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}
		}
	}()
}
