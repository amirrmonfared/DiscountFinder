package main

import (
	"database/sql"
	"fmt"

	scrap "github.com/amirrmonfared/DiscountFinder/internal/scraper"
	"github.com/amirrmonfared/DiscountFinder/internal/telegram"
	"github.com/gocolly/colly"
)


func RunScrap(webPage string, conn *sql.DB) (*colly.Collector, error) {
	scrap, err := scrap.Scraper(webPage, conn)
	if err != nil {
		fmt.Println(err)
	}
	return scrap, nil
}

func RunDiscountFinder(conn *sql.DB)  error {
	err := scrap.DiscountFinder(conn)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func RunBot(conn *sql.DB) {
	telegram.Bot(conn)
}
