package main

import (
	"database/sql"
	"fmt"

	scrap "github.com/amirrmonfared/DiscountFinder/internal/scraper"
	"github.com/gocolly/colly"
)


func RunScrap(webPage string, conn *sql.DB) (*colly.Collector, error) {
	scrap, err := scrap.Scraper(webPage, conn)
	if err != nil {
		fmt.Println(err)
	}
	return scrap, nil
}

func RunDiscountFinder(conn *sql.DB) ([]scrap.ProductOnSale, error) {
	discount, err := scrap.DiscountFinder(conn)
	if err != nil {
		fmt.Println(err)
	}
	return discount, nil
}
