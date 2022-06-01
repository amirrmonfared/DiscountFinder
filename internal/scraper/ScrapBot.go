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
	Products                      = make([]db.Product, 0, 200)
	imFalse                       = false
	imTrue                        = true
)

func scrapBotConfig() (*colly.Collector, error) {

	scrapBot := colly.NewCollector(
		colly.AllowedDomains(allowedURL1, allowedURL2),
		colly.MaxDepth(2),
	)

	time.AfterFunc(10*time.Second, func() {
		imFalse = true
	})

	scrapBot.OnRequest(func(r *colly.Request) {
		if imFalse == imTrue {
			panic("exit")
		}
	})

	// Visit next page
	scrapBot.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	return scrapBot, nil
}

func ScrapBot(webPage string, store db.Store) error {
	defer func() {
		if r := recover(); r != nil {
			log.Println("exit crawl")
		}
	}()
	collector, err := scrapBotConfig()
	if err != nil {
		log.Println("cannot recive scrap bot config", err)
	}

	collector.OnHTML(sectionTag, func(e *colly.HTMLElement) {
		products, err := addProductToSlice(e)
		if err != nil {
			log.Println("cannot add products to slice")
		}

		for _, p := range products {
			storeProduct(store, p)
			removeProductFromSlice()
		}
	})

	collector.Visit(webPage)

	return nil
}

func addProductToSlice(e *colly.HTMLElement) ([]db.Product, error) {
	products := db.Product{
		Brand: e.ChildAttr(productTitleTag, "title"),
		Link:  URL + e.ChildAttr(linkTag, "href"),
		Price: e.ChildText(priceTag),
	}

	Products = append(Products, products)

	return Products, nil
}

func removeProductFromSlice() (db.Product, error) {
	l := len(Products) - 1
	toRemove := Products[l]
	Products = Products[:l]
	return toRemove, nil
}

func storeProduct(store db.Store, p db.Product) (db.StoreProductResult, error) {
	result, err := store.StoreProduct(context.Background(), db.StoreProductParams{
		Brand: p.Brand,
		Link:  p.Link,
		Price: p.Price,
	})
	if err != nil {
		log.Println("Cannot store on product", err)
	}

	return result, nil
}
