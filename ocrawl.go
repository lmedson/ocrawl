package crawler

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func Crawl(url string) CrawlerResult {
	var crawler CrawlerResult

	crawler = CrawlerResult{
		baseURL:   url,
		linksList: []string{url},
	}

	var crawledList []string

	for len(crawler.linksList) != 0 {
		currentURL := crawler.linksList[0]

		if !Contains(crawledList, currentURL) {

			var linksRelateds []string

			result, _ := http.Get(currentURL)
			page, _ := goquery.NewDocumentFromResponse(result)
			// Add the url currently being crawled to crawleds list
			crawledList = append(crawledList, currentURL)

			page.Find("a").Each(func(i int, s *goquery.Selection) {
				// get found url
				foundURL, _ := s.Attr("href")
				// resolve current found url, checking if belongs to base url
				resolvedURL := ResolveUrls(foundURL, url)

				if len(resolvedURL) != 0 && !Contains(linksRelateds, resolvedURL) {
					/*
						if current resolved url belongs to base url,and not yet in current
						crawled url, push toarray with relashionships
					*/
					linksRelateds = append(linksRelateds, resolvedURL)
					// if the found url yet not craled, push to the list to be crawled
					if !Contains(crawledList, resolvedURL) && !Contains(crawler.linksList, resolvedURL) {
						crawler.linksList = append(crawler.linksList, resolvedURL)
					}
				}
			})

			newrelation := Relations{
				Page:         currentURL,
				RelatedLinks: linksRelateds,
			}

			crawler.RelationLinks = append(crawler.RelationLinks, newrelation)
			crawler.linksList = Remove(crawler.linksList, currentURL)
		}
	}

	crawler.Crawled = crawledList
	fmt.Printf("Crawling completed with %d urls mapped ", len(crawler.Crawled))

	return crawler
}
