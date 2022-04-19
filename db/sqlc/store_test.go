package db

import (
	"context"
	"testing"

	"github.com/amirrmonfared/DiscountFinder/util"
	"github.com/stretchr/testify/require"
)

func TestCreateProduct(t *testing.T) {
	store := NewStore(testDB)

	errs := make(chan error)
	results := make(chan CreateProductResult)

	n := 5

	for i := 0; i < n; i++ {
		go func() {
			result, err := store.CreateProduct(context.Background(), CreateFirstProductParams{
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

func TestReviewProduct(t *testing.T) {
	store := NewStore(testDB)

	errs := make(chan error)
	results := make(chan CreateSecondProductResult)

	n := 5

	for i := 0; i < n; i++ {
		go func() {
			result, err := store.ReviewProduct(context.Background(), CreateSecondParams{
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

func TestLengthOfFirst(t *testing.T) {
	store := NewStore(testDB)

	result, err := store.LengthOfFirst(context.Background())
	require.NoError(t, err)
	require.NotZero(t, result)
}

func TestLengthOfSecond(t *testing.T) {
	store := NewStore(testDB)

	result, err := store.LengthOfSecond(context.Background())
	require.NoError(t, err)
	require.NotZero(t, result)
}

func TestLengthOfOnSale(t *testing.T) {
	store := NewStore(testDB)

	result, err := store.LengthOfOnSale(context.Background())
	require.NoError(t, err)
	require.NotZero(t, result)
}
