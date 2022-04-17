package scrap

import (
	"context"
	"testing"

	db "github.com/amirrmonfared/DiscountFinder/db/sqlc"
	"github.com/stretchr/testify/require"
)

var TestProductsFromFirst = make([]ProductFromFirst, 0, 200)
//var TestProductsFromSecond = make([]ProductFromSecond, 0, 200)

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
		productFromFirst := ProductFromFirst{
			Brand: a.Brand,
			Link:  a.Link,
			Price: a.Price,
		}

		TestProductsFromFirst = append(TestProductsFromFirst, productFromFirst)
	}

	info, err := getInfoFromFirst(testDB)
	require.NoError(t, err)
	require.Equal(t, info, TestProductsFromFirst)
}