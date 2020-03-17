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

			crawledList = append(crawledList, currentURL)

			page.Find("a").Each(func(i int, s *goquery.Selection) {
				foundURL, _ := s.Attr("href")
				resolvedURL := ResolveUrls(foundURL, url)

				if len(resolvedURL) != 0 && !Contains(linksRelateds, resolvedURL) {
					linksRelateds = append(linksRelateds, resolvedURL)

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
