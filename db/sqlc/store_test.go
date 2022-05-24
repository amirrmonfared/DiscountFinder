package db

import (
	"context"
	"testing"

	"github.com/amirrmonfared/DiscountFinder/util"
	"github.com/stretchr/testify/require"
)

func TestStoreProduct(t *testing.T) {
	store := NewStore(testDB)

	errs := make(chan error)
	results := make(chan CreateProductResult)

	n := 5

	for i := 0; i < n; i++ {
		go func() {
			result, err := store.StoreProduct(context.Background(), CreateProductParams{
				Brand: util.RandomString(4),
				Link:  util.RandomLink(),
				Price: util.RandomPriceString(3),
			})

			errs <- err
			results <- result
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)
	}
}

func TestStoreOnSale(t *testing.T) {
	store := NewStore(testDB)

	errs := make(chan error)
	results := make(chan CreateOnSaleResult)

	n := 5

	for i := 0; i < n; i++ {
		go func() {
			result, err := store.StoreOnSale(context.Background(), CreateOnSaleParams{
				Brand:         util.RandomString(4),
				Link:          util.RandomLink(),
				Price:         util.RandomPriceString(3),
				PreviousPrice: util.RandomPriceString(3),
			})

			errs <- err
			results <- result
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)
	}
}
