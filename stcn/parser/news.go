package parser

import (
	"NewsSpider/engine"
	"NewsSpider/model"
	"regexp"
)

var contentRe  = regexp.MustCompile(`<p>(.*)</p>`)
var titleRe  = regexp.MustCompile(`<h2>(.*)</h2>`)
var timeRe  = regexp.MustCompile(`<div class="info">(.*)<span>来源：证券时报网</span></div>`)
var idRe = regexp.MustCompile(`var contentid = (\d+);`)

func NewsParser(contents []byte, url string) engine.ParseResult {
	newsPro := model.NewsFields{}
	newsPro.Content = extractString(contents, contentRe)
	newsPro.Title = extractString(contents, titleRe)
	newsPro.PublishTime = extractString(contents, timeRe)

	result := engine.ParseResult{
		Items: []engine.Item{
			{
				Url: 	url,
				Type:	"news",
				Id:		extractString(contents, idRe),
				Payload:	newsPro,
			},
		},
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	}else {
		return ""
	}
}
