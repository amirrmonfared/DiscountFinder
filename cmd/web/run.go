package main

import (
	"fmt"
	"log"

	db "github.com/amirrmonfared/DiscountFinder/db/sqlc"
	scrap "github.com/amirrmonfared/DiscountFinder/internal/scraper"
	"github.com/amirrmonfared/DiscountFinder/internal/telegram"
	"github.com/gocolly/colly"
	"github.com/jasonlvhit/gocron"
)

func RunScrap(webPage string, store db.Store) (*colly.Collector, error) {
	log.Println("Scrapper started")
	scrap, err := scrap.Scraper(webPage, store)
	if err != nil {
		fmt.Println(err)
	}
	return scrap, nil
}

func RunDiscountFinder(store db.Store) error {
	log.Println("DiscountFinder started")
	err := scrap.DiscountFinder(store)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func RunBot(store db.Store) {
	log.Println("Bot started")
	telegram.Bot(store)
}

func RunRemoveFirst(store db.Store) {
	log.Println("First product remover started")
	scrap.ProductRemover(store)
}

func RunRemoveOnSale(store db.Store) {
	log.Println("OnSale remover started")
	scrap.OnSaleRemover(store)
}

func cronJob(store db.Store) {
	
	gocron.Every(2).Minutes().Do(RunScrap, webPage, store)
	gocron.Every(1).Minutes().Do(RunDiscountFinder, store)
	gocron.Every(10).Minutes().Do(RunRemoveFirst, store)
	gocron.Every(10).Minutes().Do(RunRemoveOnSale, store)

	<-gocron.Start()
}
