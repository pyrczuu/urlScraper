package urlsgocraper

import (
	"log"
	"time"

	"github.com/gocolly/colly"
)

func collectURLs(url string, class string) []string {
	collector := colly.NewCollector()

	collector.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 2,
		RandomDelay: 3 * time.Second,
	})

	colly.UserAgent("Mozilla/5.0 (Windows NT; Win64; x64; rv109.0) Gecko/20100101 Firefox/121.0")

	var collected_urls []string

	collector.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL)
	})
	collector.OnResponse(func(r *colly.Response) {
		log.Println("Got a response from", r.Request.URL)
	})
	collector.OnError(func(r *colly.Response, e error) {
		log.Fatal("Error:", e)
	})
	collector.OnHTML(class, func(e *colly.HTMLElement) {
		collected_urls = append(collected_urls, e.Attr("href"))

	})

	collector.Visit(url)
	return collected_urls
}

func CollectNoFluffJobs() []string {
	urls := collectURLs("https://nofluffjobs.com/pl/artificial-intelligence?criteria=category%3Dsys-administrator,business-analyst,architecture,backend,data,ux,devops,erp,embedded,frontend,fullstack,game-dev,mobile,project-manager,security,support,testing,other", "a.posting-list-item")
	var formatted []string
	for _, url := range urls {
		formatted = append(formatted, "https://nofluffjobs.com"+url)
	}
	return formatted
}

func CollectJustJoinIT() []string {
	urls := collectURLs("https://justjoin.it/", "a.offer-card")
	var formatted []string
	for _, url := range urls {
		formatted = append(formatted, "https://justjoin.it"+url)
	}
	return formatted
}

func CollectPracujPL() []string {
	return collectURLs("https://it.pracuj.pl/praca", "[data-test=\"link-offer\"]")
}
