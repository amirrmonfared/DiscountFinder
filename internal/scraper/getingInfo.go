package scrap

import (
	"context"
	"fmt"

	db "github.com/amirrmonfared/DiscountFinder/db/sqlc"
)

func getInfoFromProduct(store db.Store) ([]ProductFromFirst, error) {

	length, err := store.GetLengthOfProducts(context.Background())
	if err != nil {
		fmt.Println(err)
	}



	arg := db.ListProductParams{
		Limit:  int32(length),
		Offset: 0,
	}

	listFirst, err := store.ListProduct(context.Background(), arg)
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
func getInfoFromOnSale(store db.Store) ([]ProductOnSale, error) {
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
			ID:       a.ID,
			Brand:    a.Brand,
			Link:     a.Link,
			Price:    a.Price,
			PrvPrice: a.PreviousPrice,
		}

		ProductsOnSale = append(ProductsOnSale, productsOnSale)
	}

	return ProductsOnSale, nil
}
