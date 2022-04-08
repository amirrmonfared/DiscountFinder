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

	// // check results
	// existed := make(map[int]bool)

	// for i := 0; i < n; i++ {
	// 	err := <-errs
	// 	require.NoError(t, err)

	// 	result := <-results
	// 	require.NotEmpty(t, result)

	// 	// check transfer
	// 	transfer := result.Transfer
	// 	require.NotEmpty(t, transfer)
	// 	require.Equal(t, account1.ID, transfer.FromAccountID)
	// 	require.Equal(t, account2.ID, transfer.ToAccountID)
	// 	require.Equal(t, amount, transfer.Amount)
	// 	require.NotZero(t, transfer.ID)
	// 	require.NotZero(t, transfer.CreatedAt)
	// }
}
