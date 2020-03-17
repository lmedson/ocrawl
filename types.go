package crawler

type CrawlerResult struct {
	baseURL       string
	linksList     []string
	RelationLinks []Relations
	Crawled       []string
}

type Relations struct {
	Page         string
	RelatedLinks []string
}
