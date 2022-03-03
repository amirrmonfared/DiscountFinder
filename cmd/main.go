package main

import (
	"crypto/tls"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/amirrmonfared/WebCrawler/internal/crawler"
	"github.com/amirrmonfared/WebCrawler/util"
	_ "github.com/lib/pq"
	"github.com/steelx/extractlinks"
)

var (
	config = &tls.Config{
		InsecureSkipVerify: true,
	}

	transport = &http.Transport{
		TLSClientConfig: config,
	}

	netClient = &http.Client{
		Transport: transport,
	}
	queue      = make(chan string)
	hasVisited = make(map[string]bool)
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	fmt.Println("connected to database:")

	defer conn.Close()
	arguments := os.Args[1:]

	if len(arguments) == 0 {
		fmt.Println("Missing URL, e.g. go-webscrapper https://www.trendyol.com/")
		os.Exit(1)
	}

	baseURL := arguments[0]
	go func() {
		queue <- baseURL
	}()

	for href := range queue {
		if !hasVisited[href] && crawler.IsSameDomain(href, baseURL) {
			crawlUrl(href)
		}

	}
}


// crawlUrl take links from href and crawling all url in the page
// I used extractlink package to extract links from the html body 
//also uses a goroutine to queue up an HTTP request for each link that has been crawled.
func crawlUrl(href string) {
	hasVisited[href] = true
	fmt.Printf("Crawling url -> %v \n", href)
	respones, err := netClient.Get(href)
	if err != nil {
		log.Fatal("cannot get baseUrl:", err)
	}
	defer respones.Body.Close()

	links, err := extractlinks.All(respones.Body)
	if err != nil {
		log.Fatal("cannot extract links:", err)
	}

	for _, link := range links {
		absoluteURL := crawler.ToFixedURL(link.Href, href)

		go func() {
			queue <- absoluteURL
		}()
	}

}
