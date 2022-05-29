package scrap

import (
	"context"
	"log"
	"time"

	db "github.com/amirrmonfared/DiscountFinder/db/sqlc"
	"github.com/gocolly/colly"
)

var (
	allowedURL1     string        = "trendyol.com"
	allowedURL2     string        = "www.trendyol.com"
	URL             string        = "https://trendyol.com"
	sectionTag      string        = ".p-card-wrppr"
	productTitleTag string        = ".prdct-desc-cntnr-wrppr span"
	linkTag         string        = ".p-card-chldrn-cntnr a"
	priceTag        string        = ".prc-box-dscntd"
	scrapTime       time.Duration = 50 * time.Second
)

var (
	Products          = make([]db.Product, 0, 200)
	ProductsForReview = make([]db.Product, 0, 200)
	ProductsOnSale    = make([]db.OnSale, 0, 200)
)

func Scraper(webPage string, store db.Store) (*colly.Collector, error) {
	imFalse := false
	imTrue := true

	collector := colly.NewCollector(
		colly.AllowedDomains(allowedURL1, allowedURL2),
		colly.MaxDepth(2),
	)

	collector.OnHTML(sectionTag, func(e *colly.HTMLElement) {
		products := addProductToSlice(e)

		for _, p := range products {
			storeProduct(store, p)
			removeProductFromSlice()
		}
	})

	time.AfterFunc(scrapTime, func() {
		imFalse = true
	})

	// Visit next page
	collector.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	defer func() {
		if r := recover(); r != nil {
			log.Println("exit crawl")
		}
	}()

	collector.OnRequest(func(r *colly.Request) {
		if imFalse == imTrue {
			panic("exit")
		}
	})

	// Start scraping on URL
	collector.Visit(webPage)

	return collector, nil
}

func storeProduct(store db.Store, p db.Product) {
	store.StoreProduct(context.Background(), db.StoreProductParams{
		Brand: p.Brand,
		Link:  p.Link,
		Price: p.Price,
	})
}

func addProductToSlice(e *colly.HTMLElement) []db.Product {
	products := db.Product{
		Brand: e.ChildAttr(productTitleTag, "title"),
		Link:  URL + e.ChildAttr(linkTag, "href"),
		Price: e.ChildText(priceTag),
	}

	Products = append(Products, products)

	return Products
}

func removeProductFromSlice() db.Product {
	l := len(Products) - 1
	toRemove := Products[l]
	Products = Products[:l]
	return toRemove
}
