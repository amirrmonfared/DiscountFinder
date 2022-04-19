package scrap

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	db "github.com/amirrmonfared/DiscountFinder/db/sqlc"
)

func DiscountFinder(conn *sql.DB) ([]ProductOnSale, error) {
	store := db.NewStore(conn)

	fromFirst, err := getInfoFromFirst(conn)
	if err != nil {
		fmt.Println(err)
	}

	fromSecond, err := getInfoFromSecond(conn)
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < len(fromFirst) && i < len(fromSecond); i++ {

		//in case of first price is less than second price
		if fromFirst[i].Price < fromSecond[i].Price {
			store.UpdateFirstProduct(context.Background(), db.UpdateFirstProductParams{
				ID:    fromFirst[i].ID,
				Price: fromSecond[i].Price,
			})
			fmt.Println("The price has risen")
		}

		//in case of second price is less than first price, first we update price of product and then add it to on sales
		if fromFirst[i].Price > fromSecond[i].Price {
			priceFirstInInt, _ := strconv.Atoi(fromFirst[i].Price)
			priceSecondInInt, _ := strconv.Atoi(fromSecond[i].Price)

			salePer := (float64(priceSecondInInt) - float64(priceFirstInInt)) / float64(priceFirstInInt) * 100.00
			discount, _ := store.CreateOnSale(context.Background(), db.CreateOnSaleParams{
				Brand:   fromFirst[i].Brand,
				Link:    fromFirst[i].Link,
				Price:   fromSecond[i].Price,
				Saleper: int64(salePer),
			})
			store.UpdateFirstProduct(context.Background(), db.UpdateFirstProductParams{
				ID:    fromFirst[i].ID,
				Price: fromSecond[i].Price,
			})
			productsOnSale := ProductOnSale{
				ID:    discount.ID,
				Brand: discount.Brand,
				Link:  discount.Link,
				Price: discount.Price,
			}
			ProductsOnSale = append(ProductsOnSale, productsOnSale)

			fmt.Println("The product is at discount")
		}

		fmt.Println("nothing")

	}

	return ProductsOnSale, nil
}
