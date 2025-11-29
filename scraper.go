package urlsgocraper

import (
	"context"
	"github.com/PuerkitoBio/goquery"
	//"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/chromedp"
	"log"
	"math/rand"
	"strings"
	"time"
)

func collectURLs(url string, selector string) []string {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Timeout
	ctx, cancel = context.WithTimeout(ctx, time.Duration(rand.Intn(800)+300)*time.Millisecond)
	defer cancel()

	var html string

	// Pobierz pełny HTML
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.WaitReady("body", chromedp.ByQuery),
		chromedp.OuterHTML("html", &html, chromedp.ByQuery),
	)
	if err != nil {
		log.Fatal("Chromedp error:", err)
	}

	// Parsowanie HTML goquery
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatal("Goquery error:", err)
	}

	var collected []string

	doc.Find(selector).Each(func(_ int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists {
			collected = append(collected, href)
		}
	})

	return collected
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
