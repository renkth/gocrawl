package engine

import (
	"crawlZhenai/fetcher"
	"fmt"
	"log"
)

type SimpleEngine struct {
}

func (e *SimpleEngine) Run(seeds ...Request) {
	var requets []Request
	for _, r := range seeds {
		requets = append(requets, r)
	}
	for len(requets) > 0 {
		r := requets[0]
		requets = requets[1:]
		/*

			log.Printf("Fetching %s",r.Url)
			body, err := fetcher.Fetch(r.Url)
			if err != nil {
				log.Printf("Fetcher: error fetching url %s %v", r.Url, err)
				continue
			}
			parseResult := r.ParserFunc(body)

		*/
		parseResult, err := Worker(r)
		if err != nil {
			continue
		}
		fmt.Println(len(parseResult.Requests))
		requets = append(requets, parseResult.Requests...)
		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}
	}
}

func Worker(r Request) (ParseResult, error) {
	log.Printf("Fetching %s", r.Url)
	body, err := fetcher.Fetch(r.Url)

	if err != nil {
		log.Printf("Fetcher: error fetching url %s %v", r.Url, err)
		return ParseResult{}, err
	}

	return r.ParserFunc(body), nil
}
