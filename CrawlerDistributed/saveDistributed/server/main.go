package main

import (
	"NewsSpider/CrawlerDistributed/config"
	"NewsSpider/CrawlerDistributed/rpcSuppert"
	"NewsSpider/CrawlerDistributed/saveDistributed"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
)

func main() {
	log.Fatal(serverRpc(fmt.Sprintf(":%d", config.ItemSaverPort), config.ElasticIndex))

	//err := serverRpc(":1234", "stock_info")
	//if err != nil {
	//	panic(err)
	//}
}

func serverRpc(host, index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return err
	}

	return rpcSuppert.ServerRpc(host, saveDistributed.ItemSaverService{
		Client: client,
		Index:  index,
	})
}
