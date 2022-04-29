package scrap

import (
	"context"
	"database/sql"
	"fmt"

	db "github.com/amirrmonfared/DiscountFinder/db/sqlc"
)

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

//getInfoFromOnSale tries to get nesseccasry information
func getInfoFromOnSale(conn *sql.DB) ([]ProductOnSale, error) {
	store := db.NewStore(conn)

	length, err := store.GetLengthOnSale(context.Background())
	if err != nil {
		fmt.Println(err)
	}

	arg := db.ListOnSaleParams{
		Limit:  int32(length),
		Offset: 0,
	}

	listOnSale, err := store.ListOnSale(context.Background(), arg)
	if err != nil {
		fmt.Println(err)
	}

	for _, a := range listOnSale {
		productsOnSale := ProductOnSale{
			ID: a.ID,
			Brand: a.Brand,
			Link:  a.Link,
			Price: a.Price,
			PrvPrice: a.PrvPrice,
		} 

		ProductsOnSale = append(ProductsOnSale, productsOnSale)
	}

	return ProductsOnSale, nil
}

var ProductsOnSaleInString = make([]string, 0, 200)
//getInfoFromOnSale tries to get nesseccasry information
func GetInfoFromOnSaleInString(conn *sql.DB) ([]string, error) {
	store := db.NewStore(conn)

	length, err := store.GetLengthOnSale(context.Background())
	if err != nil {
		fmt.Println(err)
	}

	arg := db.ListOnSaleParams{
		Limit:  int32(length),
		Offset: 0,
	}

	listOnSale, err := store.ListOnSale(context.Background(), arg)
	if err != nil {
		fmt.Println(err)
	}

	for _, a := range listOnSale {

		ProductsOnSaleInString = append(ProductsOnSaleInString, a.Brand, a.Link, a.PrvPrice, a.Price)
	}

	return ProductsOnSaleInString, nil
}
