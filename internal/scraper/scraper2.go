package scrap

import (
	"context"
	"database/sql"
	"fmt"

	db "github.com/amirrmonfared/DiscountFinder/db/sqlc"
	"github.com/gocolly/colly"
)

var SecondProducts = make([]Product2, 0, 200)

type Product2 struct {
	Brand string `json:"brand"`
	Link  string `json:"link"`
	Price string `json:"price"`
}

//Scraper2 intends to check the price of each product in the first table
//and re-store the products with new prices in the second table
func Scraper2(conn *sql.DB) (*colly.Collector, error) {
	store := db.NewStore(conn)

	info, _ := getInfoFromFirst(conn)

	// Instantiate default collector
	Collector := colly.NewCollector(
		colly.AllowedDomains("trendyol.com", "www.trendyol.com"),
		colly.MaxDepth(0),
	)

	for _, b := range info {
		collector := Collector

		//On every an element that has an attribute call callback and stores elements in the database
		collector.OnHTML(".container-right-content", func(e *colly.HTMLElement) {
			products2 := Product2{
				Brand: b.Brand,
				Link:  b.Link,
				Price: e.ChildText(".prc-dsc"),
			}

			SecondProducts = append(SecondProducts, products2)

			for _, i := range SecondProducts {
				store.ReviewProduct(context.Background(), db.CreateSecondParams{
					Brand: i.Brand,
					Link:  i.Link,
					Price: i.Price,
				})

				fmt.Println("Product reviewed and saved")
			}
		})

		collector.Visit(b.Link)
	}

	return Collector, nil
}
