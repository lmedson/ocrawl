# OCrawl - A simple crawler to map sites relations 
![GO](https://img.shields.io/badge/go-1.13.6-blue.svg)
## Installation

Install your terminal type:
`$ go get github.com/lmedson/ocrawl`

## Setup a site to crawl

First, you need to create a go file and set a url to be crawled. To visualize the relations graphically, it is possible to plot a graph through an html file, just use the plot function of the graph, passing the result of what was crawled and a desired name, to the html file.

## Plot example:

```go
    func main(){
        crawledData := crawler.Crawl("https://clojure.org/")
        crawler.Plot(crawledData, "index") // to plot a index.html file with the graph
    }
```

## Crawling assets example:

```go
    func main(){
        crawler.CrawlAssets("https://clojure.org/") // you can to plot or parse to json the outputed data
    }
```

## Formating output data example:

```go
    func main(){
        crawledData := crawler.Crawl("https://clojure.org/")
        crawler.JsonParse(crawledData, "data")
    }
```
## Runinng

Make sure you have installed all the dependencies. Run your created file, with above code:

`$ go run <your-file-name>.go`.

## Result

After run the crawl and plot a .html file with the graph, you can open it in your favorite browser and see the relations infos passing the mouse cursor by the nodes and through line connections.

The package structure folder

```bash
.
├── .gitignore                  # File with ignored files
├── LICENSE                     # Our kind of license
├── ocrawl.go                   # The main method to crawl
├── README.md                   # Readme with how to use the crawler
├── types.go                    # Types used at crawler
├── Gopkg.lock                  # It locks the version of packages
├── Gopkg.toml                  # Controls the import instructions, used by the lock file
├── <filename>.html             # After graph plot, you will have this output file
├── <filename>.json             # After json parse crawled data, you will have this output file
└── utils.py                    # File with some helpers
```
