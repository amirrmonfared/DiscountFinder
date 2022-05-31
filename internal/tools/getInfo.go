package tools

import (
	"context"
	"fmt"

	db "github.com/amirrmonfared/DiscountFinder/db/sqlc"
)

func GetInfoFromProduct(store db.Store) ([]db.Product, error) {

	length, err := store.GetLengthOfProducts(context.Background())
	if err != nil {
		fmt.Println(err)
	}

	listFirst, err := store.ListProduct(context.Background(), db.ListProductParams{
		Limit:  int32(length),
		Offset: 0,
	})
	if err != nil {
		fmt.Println(err)
	}

	return listFirst, nil
}

func GetInfoFromOnSale(store db.Store) ([]db.OnSale, error) {
	length, err := store.GetLengthOnSale(context.Background())
	if err != nil {
		fmt.Println(err)
	}

	listOnSale, err := store.ListOnSale(context.Background(), db.ListOnSaleParams{
		Limit:  int32(length),
		Offset: 0,
	})
	if err != nil {
		fmt.Println(err)
	}

	return listOnSale, nil
}
