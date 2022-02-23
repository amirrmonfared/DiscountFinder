package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/steelx/extractlinks"
)

var(
	config = &tls.Config{
		InsecureSkipVerify: true,
	}

	transport = &http.Transport{
		TLSClientConfig: config,
	}

	netClient = &http.Client{
		Transport: transport,
	}
)

func main() {
	arguments := os.Args[1:]
	
	if len(arguments) == 0 {
		fmt.Println("Missing URL, e.g. go-webscrapper http://js.org/")
		os.Exit(1)
	}

	baseURL := arguments[0]
	fmt.Println("baseURL", baseURL)

	crawlUrl(baseURL)
}


func crawlUrl(href string) {
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
		crawlUrl(toFixedURL(link.Href, href))
	}

}

func toFixedURL(href, baseUrl string) string {
	uri, err := url.Parse(href)
	if err != nil {
		return "cannot parse uri"
	}
	
	base, err := url.Parse(baseUrl)
	if err != nil {
		return "cannot parse base url"
	}

	toFixedUri := base.ResolveReference(uri)

	return toFixedUri.String()
}
