package crawler

import (
	"fmt"
	"net/http"
	"strings"

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

func CrawlAssets(url string) CrawlerResult {
	result := Crawl(url)
	for key := range result.RelationLinks {
		result.AssetsMapList = append(result.AssetsMapList, AssetsMap{
			Page:   result.RelationLinks[key].Page,
			Images: []Img{},
			Css:    []string{},
			Js:     []string{},
		})

		res, _ := http.Get(result.RelationLinks[key].Page)
		page, _ := goquery.NewDocumentFromResponse(res)

		page.Find("img").Each(func(i int, s *goquery.Selection) {
			alt, _ := s.Attr("alt")
			src, _ := s.Attr("src")

			resolvedURL := ResolveUrls(src, url)

			if (len(resolvedURL) > 0) && (len(alt) > 0) {
				result.AssetsMapList[key].Images = append(result.AssetsMapList[key].Images, Img{
					ImageLink: resolvedURL,
					ImageName: alt,
				})
			}
		})

		page.Find("link").Each(func(i int, s *goquery.Selection) {
			link, _ := s.Attr("href")

			if strings.HasSuffix(link, ".css") {
				if strings.HasPrefix(link, "/") {
					result.AssetsMapList[key].Css = append(result.AssetsMapList[key].Css, url+link)
				} else {
					result.AssetsMapList[key].Css = append(result.AssetsMapList[key].Css, link)
				}
			}
		})

		page.Find("script").Each(func(i int, s *goquery.Selection) {
			link, _ := s.Attr("src")

			if len(link) > 0 {
				if strings.HasPrefix(link, "/") {
					result.AssetsMapList[key].Js = append(result.AssetsMapList[key].Js, url+link)
				} else {
					result.AssetsMapList[key].Js = append(result.AssetsMapList[key].Js, link)
				}
			}
		})
	}
	return result
}
