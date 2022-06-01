package main

import (
	"fmt"
	"log"

	db "github.com/amirrmonfared/DiscountFinder/db/sqlc"
	scrap "github.com/amirrmonfared/DiscountFinder/internal/scraper"
	"github.com/amirrmonfared/DiscountFinder/internal/telegram"
	"github.com/amirrmonfared/DiscountFinder/internal/tools"
	"github.com/jasonlvhit/gocron"
)

func RunScrap(webPage string, store db.Store) error {
	log.Println("Scrapper started")
	err := scrap.ScrapBot(webPage, store)
	if err != nil {
		fmt.Println(err)
	}
	return nil
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

func RunTelegramBot(store db.Store) {
	log.Println("Telegram Bot started")
	telegram.Bot(store)
}

func RunRemoveFirst(store db.Store) {
	log.Println("First product remover started")
	tools.ProductRemover(store)
	log.Println("RemoveFirst Done!")
}

func RunRemoveOnSale(store db.Store) {
	log.Println("OnSale remover started")
	tools.OnSaleRemover(store)
	log.Println("RemoveOnSale Done!")
}

func cronJob(store db.Store) {

	gocron.Every(1).Minutes().Do(RunScrap, webPage, store)
	gocron.Every(1).Minutes().Do(RunDiscounter, store)
	gocron.Every(1).Minutes().Do(RunRemoveFirst, store)
	gocron.Every(1).Minutes().Do(RunRemoveOnSale, store)

	<-gocron.Start()
}
