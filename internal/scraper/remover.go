package scrap

import (
	"context"
	"database/sql"
	"log"

	db "github.com/amirrmonfared/DiscountFinder/db/sqlc"
)

func ProductRemover(conn *sql.DB) {
	store := db.NewStore(conn)
	first, err := getInfoFromFirst(conn)
	if err != nil {
		log.Println(err)
	}

	for _, j := range first {
		store.DeleteFirstProduct(context.Background(), j.ID)
	}
}

func OnSaleRemover(conn *sql.DB) {
	store := db.NewStore(conn)
	onSale, err := getInfoFromOnSale(conn)
	if err != nil {
		log.Println(err)
	}

	for _, j := range onSale {
		store.DeleteOnSale(context.Background(), j.ID)
	}
}
