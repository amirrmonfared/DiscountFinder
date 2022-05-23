package scrap

import (
	"context"
	"log"

	db "github.com/amirrmonfared/DiscountFinder/db/sqlc"
)

func ProductRemover(store db.Store) {
	first, err := getInfoFromProduct(store)
	if err != nil {
		log.Println(err)
	}

	for _, j := range first {
		store.DeleteFirstProduct(context.Background(), j.ID)
	}
}

func OnSaleRemover(store db.Store) {
	onSale, err := getInfoFromOnSale(store)
	if err != nil {
		log.Println(err)
	}

	for _, j := range onSale {
		store.DeleteOnSale(context.Background(), j.ID)
	}
}
