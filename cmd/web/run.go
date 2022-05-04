package main

import (
	"database/sql"
	"fmt"
	"log"

	scrap "github.com/amirrmonfared/DiscountFinder/internal/scraper"
	"github.com/amirrmonfared/DiscountFinder/internal/telegram"
	"github.com/gocolly/colly"
)

func RunScrap(webPage string, conn *sql.DB) (*colly.Collector, error) {
	log.Println("Scrapper started")
	scrap, err := scrap.Scraper(webPage, conn)
	if err != nil {
		fmt.Println(err)
	}
	return scrap, nil
}

func RunDiscountFinder(conn *sql.DB) error {
	log.Println("DiscountFinder started")
	err := scrap.DiscountFinder(conn)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func RunBot(conn *sql.DB) {
	log.Println("Bot started")
	telegram.Bot(conn)
}

func RunRemoveFirst(conn *sql.DB) {
	log.Println("First product remover started")
	scrap.ProductRemover(conn)
}

func RunRemoveOnSale(conn *sql.DB) {
	log.Println("OnSale remover started")
	scrap.OnSaleRemover(conn)
}
