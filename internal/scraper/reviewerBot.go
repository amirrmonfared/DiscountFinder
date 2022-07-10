package scrap

import (
	"context"
	"log"

	db "github.com/amirrmonfared/DiscountFinder/db/sqlc"
	"github.com/amirrmonfared/DiscountFinder/internal/tools"
	"github.com/gocolly/colly"
)

var (
	reviewSectionTag string = ".container-right-content"
	reviewPriceTag   string = ".prc-dsc"
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
		toCheck := toCheckPrice{reviewedProduct: reviewedProduct, productForReview: productForReview}
		err := discountFinder(toCheck, store)
		if err != nil {
			log.Println(err)
		}
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

func discountFinder(t toCheckPrice, store db.Store) error {
	firstPrice := t.productForReview.Price
	secondPrice := t.reviewedProduct.Price

	if isOnSale(firstPrice, secondPrice) {
		product := db.OnSale{
			Brand:         t.productForReview.Brand,
			Link:          t.productForReview.Link,
			Price:         t.reviewedProduct.Price,
			PreviousPrice: t.productForReview.Price,
		}
		err := storeOnSale(store, product)
		if err != nil {
			return err
		}
		return nil
	}
	if priceUpdated(firstPrice, secondPrice) {

	}

	return nil
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

func storeOnSale(store db.Store, pr db.OnSale) error {
	_, err := store.StoreOnSale(context.Background(), db.CreateOnSaleParams{
		Brand:         pr.Brand,
		Link:          pr.Link,
		Price:         pr.Price,
		PreviousPrice: pr.PreviousPrice,
	})
	if err != nil {
		return err
	}

	return nil
}
