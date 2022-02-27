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
	queue = make(chan string)
	hasVisited = make(map[string]bool)
)

func main() {
	arguments := os.Args[1:]
	
	if len(arguments) == 0 {
		fmt.Println("Missing URL, e.g. go-webscrapper https://www.trendyol.com/")
		os.Exit(1)
	}

	baseURL:= arguments[0]
	go func() {
		queue <- baseURL
	}()

	for href := range queue {
		if !hasVisited[href] && isSameDomain(href, baseURL){
			crawlUrl(href)
		}
		
	}
}

func isSameDomain(href, baseUrl string) bool {
	uri, err := url.Parse(href)
	if err != nil {
		return false
	}
	parentUri, err := url.Parse(baseUrl)
	if err != nil {
		return false
	}

	if uri.Host != parentUri.Host{
		return false
	}

	return true 
}

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
		absoluteURL := toFixedURL(link.Href, href)

		go func() {
			queue <- absoluteURL
		}()
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
