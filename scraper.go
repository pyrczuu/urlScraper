package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func CollectPracujPL() []string{
	url := "https://it.pracuj.pl/praca"
	collector := colly.NewCollector()

	var collected_urls []string

	collector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	collector.OnResponse(func(r *colly.Response) {
		fmt.Println("Got a response from", r.Request.URL)
	})
	collector.OnError(func(r *colly.Response, e error) {
		fmt.Println("Error:", e)
	})
	collector.OnHTML("a.tiles_cnb3rfy.core_n194fgoq", func(e *colly.HTMLElement) {
		collected_urls = append(collected_urls, e.Attr("href"))

	})
	collector.OnScraped(func(r *colly.Response) {
		for _, url := range collected_urls {
			fmt.Println(url)
		}
	})

	collector.Visit(url)
 return collected_urls

}
