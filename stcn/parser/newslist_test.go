package parser

import (
	"NewsSpider/fetcher"
	"fmt"
	"testing"
)

func TestNewsListParser(t *testing.T) {
	const url  = "http://kuaixun.stcn.com/index.shtml"
	contents, err := fetcher.FetchNews(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf(string(contents))
	NewsListParser(contents)
}
