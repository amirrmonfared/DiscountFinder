package scrap

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	db "github.com/amirrmonfared/DiscountFinder/db/sqlc"
	"github.com/gocolly/colly"
)

type Product struct {
	Brand string `json:"brand"`
	Link  string `json:"link"`
	Price string `json:"price"`
}

var Products = make([]Product, 0, 200)

//Scraper starts scraping on webpage and stores products on first product table.
func Scraper(webPage string, conn *sql.DB) (*colly.Collector, error) {
	store := db.NewStore(conn)

	// Instantiate default collector
	collector := colly.NewCollector(
		colly.AllowedDomains("trendyol.com", "www.trendyol.com"),
		colly.MaxDepth(2),
	)

	// On every a element which has attribute call callback and store elemnts in database
	collector.OnHTML(".p-card-wrppr", func(e *colly.HTMLElement) {
		log.Println("product found", e.Request.URL)
		products := Product{
			Brand: e.ChildAttr(".prdct-desc-cntnr-wrppr span", "title"),
			Link:  "https://trendyol.com" + e.ChildAttr(".p-card-chldrn-cntnr a", "href"),
			Price: e.ChildText(".prc-box-dscntd"),
		}
		Products = append(Products, products)

		for _, i := range Products {
			store.CreateProduct(context.Background(), db.CreateProductParams{
				Brand: i.Brand,
				Link:  i.Link,
				Price: i.Price,
			})
		}
	})

	// Visit next page
	collector.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	collector.OnResponse(func(r *colly.Response) {
		fmt.Println("Status Code:", r.StatusCode)
	})

	collector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on Trendyol.com
	collector.Visit(webPage)

	return collector, nil
}
