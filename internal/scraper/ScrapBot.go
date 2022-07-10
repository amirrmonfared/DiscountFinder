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
	imFalse                       = false
	imTrue                        = true
)

func scrapBotConfig() (*colly.Collector, error) {

	scrapBot := colly.NewCollector(
		colly.AllowedDomains(allowedURL1, allowedURL2),
		colly.MaxDepth(2),
	)

	time.AfterFunc(scrapTime, func() {
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
		product := db.Product{
			Brand: e.ChildAttr(productTitleTag, "title"),
			Link:  URL + e.ChildAttr(linkTag, "href"),
			Price: e.ChildText(priceTag),
		}
		storeProduct(store, product)
	})

	collector.Visit(webPage)

	return nil
}

func storeProduct(store db.Store, p db.Product) (db.StoreProductResult, error) {
	result, _ := store.StoreProduct(context.Background(), db.StoreProductParams{
		Brand: p.Brand,
		Link:  p.Link,
		Price: p.Price,
	})

	return result, nil
}
