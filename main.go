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
)

//to skip ssl certificate
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
		fmt.Println("Missing URL, Just enter product page Url like : https://www.trendyol.com/erkek-t-shirt-x-g2-c73?pi=2")
		os.Exit(1)
	}

	baseURL := arguments[0]
	go func() {
		queue <- baseURL
	}()



	for href := range queue {
		if !hasVisited[href] && crawler.IsSameDomain(href, baseURL) {
			CrawlUrl(href)
			crawler.Scraper(href)
		}

	}
}


// crawlUrl take links from href and crawling all url in the page
// I used extractlink package to extract links from the html body 
//also uses a goroutine to queue up an HTTP request for each link that has been crawled.
func CrawlUrl(href string) {
	hasVisited[href] = true
	fmt.Printf("Crawling url -> %v \n", href)
	respones, err := netClient.Get(href)
	if err != nil {
		log.Fatal("cannot get baseUrl:", err)
	}
	defer respones.Body.Close()

	links, err := crawler.ExtractLink(respones.Body)
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
