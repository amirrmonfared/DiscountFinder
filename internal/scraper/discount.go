package scrap

import (
	"context"
	"fmt"
	"log"

	db "github.com/amirrmonfared/DiscountFinder/db/sqlc"
	"github.com/amirrmonfared/DiscountFinder/internal/tools"
	"github.com/gocolly/colly"
)

var (
	reviewSectionTag string = ".container-right-content"
	reviewPriceTag   string = ".prc-dsc"
	ProductsForReview = make([]db.Product, 0, 200)
	ProductsOnSale    = make([]db.OnSale, 0, 200)
)

// reviwer review product from product table and storing products into slice
func Reviwer(store db.Store) ([]db.Product, *colly.Collector, error) {
	firstProducts, err := tools.GetInfoFromProduct(store)
	if err != nil {
		fmt.Println("cannot get first Products", err)
	}

	collector := colly.NewCollector(
		colly.AllowedDomains(allowedURL1, allowedURL2),
		colly.MaxDepth(0),
	)

	for _, product := range firstProducts {
		collector.OnHTML(reviewSectionTag, func(e *colly.HTMLElement) {

			productForReview := db.Product{
				ID:    product.ID,
				Brand: product.Brand,
				Link:  product.Link,
				Price: e.ChildText(reviewPriceTag),
			}
			ProductsForReview = append(ProductsForReview, productForReview)
			fmt.Println(productForReview)

		})

		collector.Visit(product.Link)
	}

	return ProductsForReview, collector, nil
}

func DiscountFinder(fromFirst []db.Product, fromSecond []db.Product) ([]db.OnSale, error) {

	for i := 0; i < len(fromFirst) && i < len(fromSecond); i++ {
		if fromFirst[i].Price > fromSecond[i].Price {
			productsOnSale := db.OnSale{
				ID:    fromFirst[i].ID,
				Brand: fromFirst[i].Brand,
				Link:  fromFirst[i].Link,
				Price: fromSecond[i].Price,
			}
			ProductsOnSale = append(ProductsOnSale, productsOnSale)
		} else {
			//TODO: add update first product price
			continue
		}
	}

	return ProductsOnSale, nil
}

func uniquedForStore(store db.Store) ([]db.OnSale, error) {
	review, _, err := Reviwer(store)
	if err != nil {
		log.Println(err)
	}
	fromSecond, err := tools.UniqueReview(review)
	if err != nil {
		log.Println(err)
	}

	fromFirst, err := tools.GetInfoFromProduct(store)
	if err != nil {
		log.Println(err)
	}

	onSale, err := DiscountFinder(fromFirst, fromSecond)
	if err != nil {
		log.Println(err)
	}
	fromOnSale, err := tools.UniqueOnSale(onSale)
	if err != nil {
		log.Println(err)
	}

	return fromOnSale, nil
}

func Discounter(store db.Store) error {
	fromOnSale, err := uniquedForStore(store)
	if err != nil {
		log.Println(err)
	}

	fromFirst, err := tools.GetInfoFromProduct(store)
	if err != nil {
		log.Println(err)
	}

	for i := 0; i < len(fromOnSale); i++ {
		store.StoreOnSale(context.Background(), db.CreateOnSaleParams{
			Brand:         fromOnSale[i].Brand,
			Link:          fromOnSale[i].Link,
			Price:         fromOnSale[i].Price,
			PreviousPrice: fromFirst[i].Price,
		})
	}

	return nil

}
