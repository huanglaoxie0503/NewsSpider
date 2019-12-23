package engine

import (
	"log"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan    chan Item
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (c *ConcurrentEngine) Run(seeds ...Request) {
	//in := make(chan Request)
	// 初始化结果输出channel
	out := make(chan ParseResult)
	//c.Scheduler.ConfigureMasterWorkerChan(in)
	c.Scheduler.Run()

	// 传入 request 返回结果给 out channel
	for i := 0; i < c.WorkerCount; i++ {
		creatWorker(c.Scheduler.WorkerChan(), out, c.Scheduler)
	}

	// url 去重后把 request 提交给 调度器（Scheduler）
	for _, r := range seeds {
		if isDuplicate(r.Url) {
			log.Printf("Duplicate request url list: "+"%s", r.Url)
			continue
		}
		c.Scheduler.Submit(r)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			// 每一个 item 开一个 goroutine
			go func() { c.ItemChan <- item }()
		}
		// url 去重
		for _, request := range result.Requests {
			if isDuplicate(request.Url) {
				log.Printf("Duplicate request: "+"%s", request.Url)
				continue
			}
			c.Scheduler.Submit(request)
		}
	}
}

func creatWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			// tell scheduler i`m ready
			ready.WorkerReady(in)
			request := <-in
			result, err := Worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

// Url 去重
var visitedUrls = make(map[string]bool)

func isDuplicate(url string) bool {
	if visitedUrls[url] {
		return true
	}
	visitedUrls[url] = true
	return false
}
