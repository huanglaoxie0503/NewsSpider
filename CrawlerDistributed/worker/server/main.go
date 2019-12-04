package main

import (
	"NewsSpider/CrawlerDistributed/config"
	"NewsSpider/CrawlerDistributed/rpcSuppert"
	"NewsSpider/CrawlerDistributed/worker"
	"log"
)

func main() {
	log.Fatal(rpcSuppert.ServerRpc(":%d", config.WorkerPort0), worker.CrawlService{})
}
