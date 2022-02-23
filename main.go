package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	"github.com/steelx/extractlinks"
)

func main() {
	baseURL := "https://www.trendyol.com/"

	config := &tls.Config{
		InsecureSkipVerify: true,
	}

	transport := &http.Transport{
		TLSClientConfig: config,
	}

	netClient := &http.Client{
		Transport: transport,
	}

	respones, err := netClient.Get(baseURL)
	if err != nil {
		log.Fatal("cannot get baseUrl:", err)
	}
	defer respones.Body.Close()

	links, err := extractlinks.All(respones.Body)
	if err != nil {
		log.Fatal("cannot extract links:", err)
	}

	
	for i, link := range links {
		fmt.Printf("index %v -- link %+v \n", i, link)
	}

}
