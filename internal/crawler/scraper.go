package crawler

import (
	"fmt"

	"github.com/gocolly/colly"
)

type Product struct {
	Brand string
	Price string
	URL   string
}

func Scraper() {
	// Instantiate default collector
	collector := colly.NewCollector(
		colly.AllowedDomains("trendyol.com", "www.trendyol.com"),
		colly.MaxDepth(2),
	)

	// On every a element which has href attribute call callback
	collector.OnHTML(".p-card-wrppr", func(e *colly.HTMLElement) {
		product := Product{}
		product.Brand = e.ChildAttr(".prdct-desc-cntnr-wrppr span", "title")
		product.Price = e.ChildText(".prc-box-dscntd")
		product.URL = "https://trendyol.com" + e.ChildAttr(".p-card-chldrn-cntnr a", "href")
		fmt.Println(product.Brand, product.Price, product.URL)
	})

	// Visit next page
	collector.OnHTML("a[href]", func(e *colly.HTMLElement) {
		//fmt.Println("Next page link found:", e.Attr("href"))
		e.Request.Visit(e.Attr("href"))
	})

	collector.OnResponse(func(r *colly.Response) {
		fmt.Println("Status Code:", r.StatusCode)
	})

	collector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	collector.Visit("https://www.trendyol.com/erkek-t-shirt-x-g2-c73?pi=2")
}
