package parser

import (
	"NewsSpider/engine"
	"regexp"
)

var newsListRe  = regexp.MustCompile(` <a href="(http://kuaixun.stcn.com/\d+/\d+/\d+.shtml)" title=".*" target="_blank">(.*)</a>`)
var nextPageRe = regexp.MustCompile(`<a href="(http://kuaixun.stcn.com/index_\d+.shtml)">\d+</a>`)

// 新闻列表解析器
func NewsListParser(contents []byte, url string) engine.ParseResult {
	matches := newsListRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		url := string(m[1])
		//title := string(m[2])
		//result.Items = append(result.Items, title)
		result.Requests = append(result.Requests, engine.Request{
			Url:        url,
			ParserFunc: NewsParser,
		})
	}
	// 翻页操作
	nextPages := nextPageRe.FindAllSubmatch(contents, -1)
	for _, nextPage := range nextPages{
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(nextPage[1]),
			ParserFunc: NewsListParser,
		})
	}
	return result
}

