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

func RunDiscounter(store db.Store) error {
	log.Println("Discounter started")
	err := scrap.Discounter(store)
	if err != nil {
		fmt.Println(err)
	}
	log.Println("Discounter Done!")
	return nil
}

func RunBot(store db.Store) {
	log.Println("Bot started")
	telegram.Bot(store)
}

func RunRemoveFirst(store db.Store) {
	log.Println("First product remover started")
	scrap.ProductRemover(store)
	log.Println("RemoveFirst Done!")
}

func RunRemoveOnSale(store db.Store) {
	log.Println("OnSale remover started")
	scrap.OnSaleRemover(store)
	log.Println("RemoveOnSale Done!")
}

func cronJob(store db.Store) {

	gocron.Every(1).Minutes().Do(RunScrap, webPage, store)
	gocron.Every(1).Minutes().Do(RunDiscounter, store)
	gocron.Every(1).Minutes().Do(RunRemoveFirst, store)
	gocron.Every(1).Minutes().Do(RunRemoveOnSale, store)

	<-gocron.Start()
}
