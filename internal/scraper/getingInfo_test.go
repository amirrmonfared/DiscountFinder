package scrap

import (
	"context"
	"testing"

	db "github.com/amirrmonfared/DiscountFinder/db/sqlc"
	"github.com/stretchr/testify/require"
)

func TestGetInfoFromFirst(t *testing.T) {

	testLength, err := testQueries.GetLengthOfFirst(context.Background())
	require.NoError(t, err)
	require.NotZero(t, testLength)

	arg := db.ListFirstProductParams{
		Limit:  int32(testLength),
		Offset: 0,
	}

	testListFirst, err := testQueries.ListFirstProduct(context.Background(), arg)
	require.NoError(t, err)
	require.NotZero(t, testListFirst)

	for _, a := range testListFirst {
		productsFromFirst := ProductFromFirst{
			ID: a.ID,
			Brand: a.Brand,
			Link:  a.Link,
			Price: a.Price,
		}

		ProductsFromFirst = append(ProductsFromFirst, productsFromFirst)
	}

	info, err := getInfoFromFirst(testDB)
	require.NoError(t, err)
	require.NotEmpty(t, info)
}

func TestGetInfoFromSecond(t *testing.T) {

	testLength, err := testQueries.GetLengthOfSecond(context.Background())
	require.NoError(t, err)
	require.NotZero(t, testLength)

	arg := db.ListSecondParams{
		Limit:  int32(testLength),
		Offset: 0,
	}

	testListSecond, err := testQueries.ListSecond(context.Background(), arg)
	require.NoError(t, err)
	require.NotZero(t, testListSecond)

	for _, a := range testListSecond {
		productsFromSecond := ProductFromSecond{
			ID: a.ID,
			Brand: a.Brand,
			Link:  a.Link,
			Price: a.Price,
		}

		ProductsFromSecond = append(ProductsFromSecond, productsFromSecond)
	}

	info, err := getInfoFromSecond(testDB)
	require.NoError(t, err)
	require.NotEmpty(t, info)
}
