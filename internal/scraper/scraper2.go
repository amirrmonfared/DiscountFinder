package scrap

import (
	"context"
	"database/sql"
	"fmt"

	db "github.com/amirrmonfared/DiscountFinder/db/sqlc"
	"github.com/gocolly/colly"
)

var LinkProducts = make([]LinkProduct, 0, 200)
var Products2 = make([]Product2, 0, 200)

type LinkProduct struct {
	Brand string `json:"brand"`
	Link string `json:"link"`
}

type Product2 struct {
	Brand string `json:"brand"`
	Link  string `json:"link"`
	Price string `json:"price"`
}

//Scraper2 intends to check the price of each product in the first table
//and re-store the products with new prices in the second table
func Scraper2(conn *sql.DB) (*colly.Collector, error) {
	store := db.NewStore(conn)

	//getInfo(conn)

	// Instantiate default collector
	Collector := colly.NewCollector(			
		colly.AllowedDomains("trendyol.com", "www.trendyol.com"),
		colly.MaxDepth(0),
	)


	for _, b := range LinkProducts {
		collector := Collector

		//On every an element that has an attribute call callback and stores elements in the database
		collector.OnHTML(".container-right-content", func(e *colly.HTMLElement) {
			products2 := Product2{
				Brand: b.Brand,
				Link: b.Link,
				Price:  e.ChildText(".prc-dsc"),
			}

			Products2 = append(Products2, products2)

			for _, i := range Products2 {
				store.ReviewProduct(context.Background(), db.CreateSecondParams{
					Brand: i.Brand,
					Link: i.Link,
					Price: i.Price,
				})

			}
		})

		collector.Visit(b.Link)
	}

	return Collector, nil
}

// getInfo tries to get nesseccasry information for itterating over table first
func getInfo(conn *sql.DB) ([]LinkProduct, error) {
	store := db.NewStore(conn)

	length, err := store.GetLengthOfFirst(context.Background())
	if err != nil {
		fmt.Println(err)
	}

	arg := db.ListFirstProductParams{
		Limit:  int32(length),
		Offset: 0,
	}

	listFirst, err := store.ListFirstProduct(context.Background(), arg)
	if err != nil {
		fmt.Println(err)
	}

	for _, a := range listFirst {
		linkProducts := LinkProduct{
			Brand: a.Brand,
			Link: a.Link,
		}

		LinkProducts = append(LinkProducts, linkProducts)
	}

	return LinkProducts, nil
}