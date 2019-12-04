package main

import (
	"NewsSpider/CrawlerDistributed/config"
	"NewsSpider/CrawlerDistributed/saveDistributed/client"
	"NewsSpider/engine"
	"NewsSpider/scheduler"
	"NewsSpider/stcn/parser"
	"fmt"
)

func main() {

	const url = "http://kuaixun.stcn.com/index.shtml"
	itemChan, err := client.ItemSaver(fmt.Sprintf(":%d", config.ItemSaverPort))
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 3,
		ItemChan:    itemChan,
	}

	e.Run(engine.Request{
		Url:        url,
		ParserFunc: parser.NewsListParser,
	})
}
