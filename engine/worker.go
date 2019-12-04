package engine

import (
	"NewsSpider/fetcher"
	"log"
)

func worker(r Request) (ParseResult, error) {
	body, err := fetcher.FetchNews(r.Url)
	if err != nil {
		log.Printf("Fetcher: error" + "fetching url %s: %v", r.Url, err)
		return ParseResult{}, err
	}
	return r.ParserFunc(body, r.Url), nil
}

