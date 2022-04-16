package scrap

import (
	"context"
	"database/sql"
	"fmt"

	db "github.com/amirrmonfared/DiscountFinder/db/sqlc"
	"github.com/gocolly/colly"
	_ "github.com/lib/pq"
)

type LinkProduct struct {
	Brand string `json:"brand"`
	Link string `json:"link"`
}

type Product2 struct {
	Brand string `json:"brand"`
	Link  string `json:"link"`
	Price string `json:"price"`
}

var LinkProducts = make([]LinkProduct, 0, 200)
var Products2 = make([]Product2, 0, 200)

func Scraper2(conn *sql.DB) (*colly.Collector, error) {
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

	Collector := colly.NewCollector(			
		colly.AllowedDomains("trendyol.com", "www.trendyol.com"),
		colly.MaxDepth(0),
	)

	for _, b := range LinkProducts {
		collector := Collector

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
