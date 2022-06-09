package scrap

import (
	"fmt"
	"log"

	db "github.com/amirrmonfared/DiscountFinder/db/sqlc"
	"github.com/amirrmonfared/DiscountFinder/internal/tools"
	"github.com/gocolly/colly"
)

type toCheckPrice struct {
	productForReview db.Product
	reviewedProduct  db.Product
}

func reviewerBotConfig() (*colly.Collector, error) {

	collector := colly.NewCollector(
		colly.AllowedDomains(allowedURL1, allowedURL2),
		colly.MaxDepth(0),
	)

	return collector, nil
}

func ReviewerBot(store db.Store) error {
	productsForReview, err := tools.GetInfoFromProduct(store)
	if err != nil {
		log.Println("cannot get first Products", err)
	}

	collector, err := reviewerBotConfig()
	if err != nil {
		log.Println("cannot recive reviewer bot config", err)
	}

	for _, productForReview := range productsForReview {
		reviewedProduct, _ := htmlCollector(collector, productForReview)
		products := toCheckPrice{reviewedProduct: reviewedProduct, productForReview: productForReview}
		discountFinder(products, store)
	}

	return nil
}

func htmlCollector(collector *colly.Collector, product db.Product) (db.Product, error) {
	var productForReview db.Product
	collector.OnHTML(reviewSectionTag, func(e *colly.HTMLElement) {
		productForReview = db.Product{
			Brand: product.Brand,
			Link:  product.Link,
			Price: e.ChildText(reviewPriceTag),
		}
	})
	collector.Visit(product.Link)
	return productForReview, nil
}

func discountFinder(t toCheckPrice, store db.Store) (db.Product, error) {
	firstPrice := t.productForReview.Price
	secondPrice := t.reviewedProduct.Price

	if isOnSale(firstPrice, secondPrice) {
		fmt.Println("first", firstPrice, secondPrice)
		return t.productForReview, nil
	}
	if priceUpdated(firstPrice, secondPrice) {
		fmt.Println("second", firstPrice, secondPrice)
		return t.productForReview, nil
	}

	return t.productForReview, nil
}

func isOnSale(firstPrice, secondPrice string) bool {
	if firstPrice > secondPrice {
		return true
	}
	return false
}

func priceUpdated(firstPrice, secondPrice string) bool {
	if firstPrice < secondPrice {
		return true
	}
	return false
}


