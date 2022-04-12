package db

import (
	"context"
	"fmt"
	"testing"
)

func TestCreateProduct(t *testing.T) {
	store := NewStore(testDB)

	row1 := CreateRandomRow(t)

	n := 5
	//amount := int64(10)

	errs := make(chan error)
	results := make(chan CreateProductResult)

	// run n concurrent transfer transaction
	for i := 0; i < n; i++ {
		go func() {
			result, err := store.CreateProduct(context.Background(), CreateProductParams{
				Brand: row1.Brand,
				Link:  row1.Link,
				Price: row1.Price,
			})
			fmt.Println("a")

			errs <- err
			results <- result
		}()
	}
}
