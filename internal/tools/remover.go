package tools

import (
	"context"
	"log"

	db "github.com/amirrmonfared/DiscountFinder/db/sqlc"
)

func ProductRemover(store db.Store) error {
	first, err := GetInfoFromProduct(store)
	if err != nil {
		log.Println(err)
		return err
	}

	for _, j := range first {
		store.DeleteProduct(context.Background(), j.ID)
	}

	return nil
}

func OnSaleRemover(store db.Store) error {
	onSale, err := GetInfoFromOnSale(store)
	if err != nil {
		log.Println(err)
		return err
	}

	for _, j := range onSale {
		store.DeleteOnSale(context.Background(), j.ID)
	}

	return nil
}
