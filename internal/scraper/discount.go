package scrap

import (
	"context"
	"fmt"
	"log"

	db "github.com/amirrmonfared/DiscountFinder/db/sqlc"
	"github.com/gocolly/colly"
)

//DiscountFinder store OnSale products
func DiscountFinder(store db.Store) error {
	//collecting data from the first table and reviewed slice
	fromFirst, review, _, err := collector(store)
	if err != nil {
		log.Println(err)
	}
	fromSecond, err := uniqueReview(review)
	if err != nil {
		log.Println(err)
	}
	onSale, err := differences(fromFirst, fromSecond)
	if err != nil {
		log.Println(err)
	}
	fromOnSale, err := uniqueOnSale(onSale)
	if err != nil {
		log.Println(err)
	}

	// iterating over fromOnSale slice to storing elements in table on_sale
	for i := 0; i < len(fromOnSale); i++ {
		store.CreateOnSale(context.Background(), db.CreateOnSaleParams{
			Brand:    fromFirst[i].Brand,
			Link:     fromFirst[i].Link,
			Price:    fromSecond[i].Price,
			PreviousPrice: fromFirst[i].Price,
		})

		fmt.Println("The product is at discount")
	}

	return nil

}

// collector trying to collect product from first table and storing products into slice
func collector(store db.Store) ([]ProductFromFirst, []ProductForReview, *colly.Collector, error) {
	firstProducts, err := getInfoFromProduct(store)
	if err != nil {
		fmt.Println("cannot get first Products", err)
	}

	Collector := colly.NewCollector(
		colly.AllowedDomains("trendyol.com", "www.trendyol.com"),
		colly.MaxDepth(0),
	)

	Collector.Limit(&colly.LimitRule{Parallelism: 1})

	for _, b := range firstProducts {
		collector := Collector
		collector.OnHTML(".container-right-content", func(e *colly.HTMLElement) {
			productForReview := ProductForReview{
				ID:    b.ID,
				Brand: b.Brand,
				Link:  b.Link,
				Price: e.ChildText(".prc-dsc"),
			}
			ProductsForReview = append(ProductsForReview, productForReview)
		})

		collector.Visit(b.Link)
	}

	return firstProducts, ProductsForReview, Collector, nil
}

// differences check the price difference between first price and reviewed price
func differences(fromFirst []ProductFromFirst, fromSecond []ProductForReview) ([]ProductOnSale, error) {

	for i := 0; i < len(fromFirst) && i < len(fromSecond); i++ {
		//in case of second price is less than first price store product into on_sale table
		if fromFirst[i].Price > fromSecond[i].Price {
			productsOnSale := ProductOnSale{
				ID:    fromFirst[i].ID,
				Brand: fromFirst[i].Brand,
				Link:  fromFirst[i].Link,
				Price: fromSecond[i].Price,
			}
			ProductsOnSale = append(ProductsOnSale, productsOnSale)
		}
	}

	return ProductsOnSale, nil
}
