package main

import (
	"NewsSpider/CrawlerDistributed/saveDistributed/client"
	"NewsSpider/engine"
	"NewsSpider/scheduler"
	"NewsSpider/stcn/parser"
)

func main() {

	const url  = "http://kuaixun.stcn.com/index.shtml"
	itemChan, err := client.ItemSaver(":1234")
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		WorkerCount: 3,
		ItemChan: itemChan,
	}

	e.Run(engine.Request{
		Url:        url,
		ParserFunc: parser.NewsListParser,
	})
}

