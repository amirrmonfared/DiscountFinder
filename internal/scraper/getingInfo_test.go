package scrap

import (
	"context"
	"testing"

	db "github.com/amirrmonfared/DiscountFinder/db/sqlc"
	"github.com/amirrmonfared/DiscountFinder/util"
	"github.com/stretchr/testify/require"
)

func TestGetInfoFromFirst(t *testing.T) {
	info, err := getInfoFromProduct(TestStore)
	require.NoError(t, err)
	require.NotEmpty(t, info)
}

func TestGetInfoFromOnSale(t *testing.T) {
	CreateRandomRowOnSale(t)
	info, err := getInfoFromOnSale(TestStore)
	require.NoError(t, err)
	require.NotEmpty(t, info)
}

func CreateRandomRowOnSale(t *testing.T) db.OnSale {
	arg := db.CreateOnSaleParams{
		Brand:    util.RandomString(5),
		Link:     util.RandomLink(),
		Price:    util.RandomPriceString(4),
		PreviousPrice: util.RandomPriceString(1),
	}

	product, err := testQueries.CreateOnSale(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, product)

	require.Equal(t, arg.Brand, product.Brand)
	require.Equal(t, arg.Link, product.Link)
	require.Equal(t, arg.Price, product.Price)
	require.Equal(t, arg.PreviousPrice, product.PreviousPrice)
	require.NotZero(t, product.ID)
	require.NotZero(t, product.CreatedAt)

	return product
}

func TestCreateRowOnSale(t *testing.T) {
	CreateRandomRowOnSale(t)
}
