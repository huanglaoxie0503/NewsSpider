package main

import (
	"NewsSpider/CrawlerDistributed/rpcSuppert"
	"NewsSpider/engine"
	"NewsSpider/model"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T)  {
	const host  = ":1234"

	go serverRpc(host, "test1")

	time.Sleep(time.Second)

	client, err := rpcSuppert.NewClient(host)
	if err != nil {
		panic(err)
	}
	item := engine.Item{
		Url: "http://kuaixun.stcn.com/2019/1204/15526045.shtml",
		Type: "news",
		Id: "15526045",
		Payload:model.NewsFields{
			Title: "三只松鼠：提前实现百亿目标，松鼠进入新时期",
			Content: "2019年12月03日午间，公司发布公告称，截止12月3日，三只松鼠2019年已实现100亿销售额（含税口径）。国信证券认为该业绩超市场预期，百亿目标提前一年实现，松鼠进入后百亿时代。休闲食品行业一般二三季度是消费者培育期（即淡季），一四季度是收获期（旺季）。公司作为线上休闲食品龙头，当前正稳步向线下渠道拓展。我们维持19-21年EPS分别为0.97/1.22/1.53元，当前股价对应PE分别为59/47/37，未来一年合理估值区间为75.67-77.29元，维持“增持”评级。",
			PublishTime: "2019-12-04 14:19",
		},
	}

	result := ""
	err = client.Call("ItemSaverService.Save", item, &result)
	if err != nil || result != "ok"{
		t.Errorf("result: %s: err: %s", result, err)
	}
}
