package main

import (
	"NewsSpider/engine"
	"NewsSpider/newsSave"
	"NewsSpider/scheduler"
	"NewsSpider/stcn/parser"
)

func main() {

	const url = "http://kuaixun.stcn.com/index.shtml"
	itemChan, err := newsSave.ItemSaver("stock_info")
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 3,
		ItemChan:    itemChan,
	}

	e.Run(engine.Request{
		Url:    url,
		Parser: engine.NewFuncParser(parser.NewsListParser, "NewsListParser"),
	})
}
