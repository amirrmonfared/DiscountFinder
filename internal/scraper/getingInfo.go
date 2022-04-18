package scrap

import (
	"context"
	"database/sql"
	"fmt"

	db "github.com/amirrmonfared/DiscountFinder/db/sqlc"
)

var ProductsFromFirst = make([]ProductFromFirst, 0, 200)
var ProductsFromSecond = make([]ProductFromSecond, 0, 200)

type ProductFromFirst struct {
	ID    int64  `json:"id"`
	Brand string `json:"brand"`
	Link  string `json:"link"`
	Price string `json:"price"`
}

type ProductFromSecond struct {
	ID    int64  `json:"id"`
	Brand string `json:"brand"`
	Link  string `json:"link"`
	Price string `json:"price"`
}

//getInfoFromFirst tries to get nesseccasry information
func getInfoFromFirst(conn *sql.DB) ([]ProductFromFirst, error) {
	store := db.NewStore(conn)

	length, err := store.GetLengthOfFirst(context.Background())
	if err != nil {
		fmt.Println(err)
	}

	arg := db.ListFirstProductParams{
		Limit:  int32(length),
		Offset: 0,
	}

	listFirst, err := store.ListFirstProduct(context.Background(), arg)
	if err != nil {
		fmt.Println(err)
	}

	for _, a := range listFirst {
		productsFromFirst := ProductFromFirst{
			ID:    a.ID,
			Brand: a.Brand,
			Link:  a.Link,
			Price: a.Price,
		}

		ProductsFromFirst = append(ProductsFromFirst, productsFromFirst)
	}

	return ProductsFromFirst, nil
}

//getInfoFromSecond tries to get nesseccasry information
func getInfoFromSecond(conn *sql.DB) ([]ProductFromSecond, error) {
	store := db.NewStore(conn)

	length, err := store.GetLengthOfSecond(context.Background())
	if err != nil {
		fmt.Println(err)
	}

	arg := db.ListSecondParams{
		Limit:  int32(length),
		Offset: 0,
	}

	listSecond, err := store.ListSecond(context.Background(), arg)
	if err != nil {
		fmt.Println(err)
	}

	for _, a := range listSecond {
		productsFromSecond := ProductFromSecond{
			ID: a.ID,
			Brand: a.Brand,
			Link:  a.Link,
			Price: a.Price,
		}

		ProductsFromSecond = append(ProductsFromSecond, productsFromSecond)
	}

	return ProductsFromSecond, nil
}
