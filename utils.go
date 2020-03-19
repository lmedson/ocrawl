package crawler

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-echarts/go-echarts/charts"
)

func Remove(urlList []string, urlToRemove string) []string {
	index := 0

	for _, a := range urlList {
		if a == urlToRemove {
			return append(urlList[:index], urlList[index+1:]...)
		}
		index++
	}
	return urlList
}

func Contains(links []string, linkToFind string) bool {
	for _, element := range links {
		if element == linkToFind {
			return true
		}
	}
	return false
}

func ResolveUrls(link string, baseURL string) string {
	if strings.HasPrefix(link, "/") {
		return baseURL + link
	} else if strings.HasPrefix(link, baseURL) {
		return link
	} else {
		return ""
	}
}

func Plot(res CrawlerResult, fileName string) {
	var urlKeys []charts.GraphNode

	// here we separate all the urls crawled, how a node
	for i := range res.Crawled {
		urlKeys = append(urlKeys, charts.GraphNode{Name: res.Crawled[i]})
	}

	graph := charts.NewGraph()
	graph.SetGlobalOptions(charts.TitleOpts{Title: "Relations Graph"})

	graph.Add("Url Relations Map", urlKeys, func() []charts.GraphLink {
		links := make([]charts.GraphLink, 0)
		// set the relations in map
		for i := 0; i < len(urlKeys); i++ {
			for j := 0; j < len(res.RelationLinks[i].RelatedLinks); j++ {
				links = append(links,
					charts.GraphLink{Source: urlKeys[i].Name, Target: res.RelationLinks[i].RelatedLinks[j]})
			}
		}
		return links
	}(),
		charts.GraphOpts{Force: charts.GraphForce{Repulsion: 5000}},
	)

	// Create the file owner of graph
	plottedGraph, err := os.Create(fileName + ".html")

	if err != nil {
		log.Println(err)
	}

	// render the graph
	graph.Render(plottedGraph)
}

func JsonParse(crawledData CrawlerResult, fileName string) {
	parsedData, _ := json.Marshal(crawledData)

	f, _ := os.Create(fileName + ".json")
	f.WriteString(string(parsedData))
	f.Close()

	fmt.Printf("Your json file with crawled data was created how %s.json", fileName)
}
