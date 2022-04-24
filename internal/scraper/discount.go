package scrap

import (
	"context"
	"database/sql"
	"fmt"

	db "github.com/amirrmonfared/DiscountFinder/db/sqlc"
	"github.com/gocolly/colly"
)

//DiscountFinder store OnSale products
func DiscountFinder(conn *sql.DB) ([]ProductOnSale, error) {
	store := db.NewStore(conn)
	//collecting data from the first table and reviewed slice
	fromFirst, review, _, _ := collector(conn)
	fromSecond, _ := uniqueReview(review)
	onSale, _ := differences(fromFirst, fromSecond)
	fromOnSale, _ := uniqueOnSale(onSale)

	//iterating over fromOnSale slice to storing elements
	// in table on_sale
	for i := 0; i < len(fromOnSale); i++ {
		// TODO: add discount percentage
		// priceFirstInInt, _ := strconv.Atoi(fromFirst[i].Price)
		// priceSecondInInt, _ := strconv.Atoi(fromSecond[i].Price)

		// salePer := (float64(priceSecondInInt) - float64(priceFirstInInt)) / float64(priceFirstInInt) * 100.00
		store.CreateOnSale(context.Background(), db.CreateOnSaleParams{
			Brand:    fromFirst[i].Brand,
			Link:     fromFirst[i].Link,
			Price:    fromSecond[i].Price,
			PrvPrice: fromFirst[i].Price,
		})
		fmt.Println("The product is at discount")
	}

	return ProductsOnSale, nil

}

//collector trying to collect product from first table
//and storing products into slice
func collector(conn *sql.DB) ([]ProductFromFirst, []ProductForReview, *colly.Collector, error) {
	firstProducts, err := getInfoFromFirst(conn)
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

func differences(fromFirst []ProductFromFirst, fromSecond []ProductForReview) ([]ProductOnSale, error) {

	for i := 0; i < len(fromFirst) && i < len(fromSecond); i++ {

		//in case of second price is less than first price
		// store product into on_sale table
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
