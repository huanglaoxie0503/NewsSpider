package engine

import (
	"NewsSpider/fetcher"
	"log"
)

func Worker(r Request) (ParseResult, error) {
	body, err := fetcher.FetchNews(r.Url)
	if err != nil {
		log.Printf("Fetcher: error"+"fetching url %s: %v", r.Url, err)
		return ParseResult{}, err
	}
	return r.Parser.Parser(body, r.Url), nil
}
