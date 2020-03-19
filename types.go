package crawler

type CrawlerResult struct {
	baseURL       string
	linksList     []string
	AssetsMapList []AssetsMap `json:"assetsMapList"`
	RelationLinks []Relations `json:"relationLinks"`
	Crawled       []string    `json:"crawled"`
}

type Relations struct {
	Page         string   `json:"page"`
	RelatedLinks []string `json:"relatedLinks"`
}

type AssetsMap struct {
	Page   string   `json:"page"`
	Js     []string `json:"js"`
	Css    []string `json:"css"`
	Images []Img    `json:"images"`
}

type Img struct {
	ImageName string `json:"imageName"`
	ImageLink string `json:"imageLink"`
}
