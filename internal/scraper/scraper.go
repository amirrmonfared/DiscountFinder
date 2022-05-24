package scrap

import (
	"context"
	"log"
	"time"

	db "github.com/amirrmonfared/DiscountFinder/db/sqlc"
	"github.com/gocolly/colly"
)

//Scraper starts scraping on webpage and stores products on first product table.
func Scraper(webPage string, store db.Store) (*colly.Collector, error) {
	imFalse := false

	// Instantiate default collector
	collector := colly.NewCollector(
		colly.AllowedDomains("trendyol.com", "www.trendyol.com"),
		colly.MaxDepth(2),
	)

	// On every a element which has attribute call callback and store elemnts in database
	collector.OnHTML(".p-card-wrppr", func(e *colly.HTMLElement) {
		products := Product{
			Brand: e.ChildAttr(".prdct-desc-cntnr-wrppr span", "title"),
			Link:  "https://trendyol.com" + e.ChildAttr(".p-card-chldrn-cntnr a", "href"),
			Price: e.ChildText(".prc-box-dscntd"),
		}
		Products = append(Products, products)

		for _, i := range Products {
			store.StoreProduct(context.Background(), db.StoreProductParams{
				Brand: i.Brand,
				Link:  i.Link,
				Price: i.Price,
			})
			//to remove last element in slice
			removeProduct()
		}
	})

	time.AfterFunc(30*time.Second, func() {
		imFalse = true
	})

	// Visit next page
	collector.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	collector.OnResponse(func(r *colly.Response) {
	})

	collector.OnRequest(func(r *colly.Request) {
		imTrue := true
		if imFalse == imTrue {
			panic("exit")
		}
	})

	defer func() {
		if r := recover(); r != nil {
			log.Println("exit crawl")
		}
	}()

	// Start scraping on URL
	collector.Visit(webPage)

	return collector, nil
}

func removeProduct() Product {
	l := len(Products) - 1
	toRemove := Products[l]
	Products = Products[:l]
	return toRemove
}
