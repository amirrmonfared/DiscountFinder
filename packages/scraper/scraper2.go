package scrap

import (
	"context"
	"database/sql"
	"fmt"

	db "github.com/amirrmonfared/DiscountFinder/db/sqlc"
	_ "github.com/lib/pq"
)

type LinkProduct struct {
	Link  string `json:"link"`
}

type Product2 struct {
	Brand string `json:"brand"`
	Link  string `json:"link"`
	Price string `json:"price"`
}

var LinkProducts = make([]LinkProduct, 0, 200)

func Scraper2(conn *sql.DB) {
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
		linkProducts := LinkProduct{
			Link: a.Link,
		}

		LinkProducts = append(LinkProducts, linkProducts)

		fmt.Println(LinkProducts)
		// firstProduct, err := store.GetFirstProduct(context.Background(), a.ID)
		// if err != nil {
		// 	fmt.Println(err)
		// }

	}

}
