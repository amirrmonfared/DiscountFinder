package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/amirrmonfared/DiscountFinder/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomRowOnSale(t *testing.T) OnSale {
	arg := CreateOnSaleParams{
		Brand: util.RandomString(5),
		Link:  util.RandomLink(),
		PrvPrice: util.RandomPriceString(1),
		Price: util.RandomPriceString(4),
	}

	product, err := testQueries.CreateOnSale(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, product)

	require.Equal(t, arg.Brand, product.Brand)
	require.Equal(t, arg.Link, product.Link)
	require.Equal(t, arg.Price, product.Price)
	require.Equal(t, arg.PrvPrice, product.PrvPrice)
	require.NotZero(t, product.ID)
	require.NotZero(t, product.CreatedAt)

	return product
}

func TestCreateRowOnSale(t *testing.T) {
	CreateRandomRowOnSale(t)
}

func TestGetOnSale(t *testing.T) {
	row1 := CreateRandomRowOnSale(t)
	row2, err := testQueries.GetOnSale(context.Background(), row1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, row2)

	require.Equal(t, row1.ID, row2.ID)
	require.Equal(t, row1.Brand, row2.Brand)
	require.Equal(t, row1.Link, row2.Link)
	require.Equal(t, row1.PrvPrice, row2.PrvPrice)
	require.Equal(t, row1.Price, row2.Price)
	require.WithinDuration(t, row1.CreatedAt, row2.CreatedAt, time.Second)
}

func TestUpdateOnSale(t *testing.T) {
	row1 := CreateRandomRowOnSale(t)

	arg := UpdateOnSaleParams{
		ID:    row1.ID,
		Price: util.RandomPriceString(4),
	}

	row2, err := testQueries.UpdateOnSale(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, row2)

	require.Equal(t, row1.ID, row2.ID)
	require.Equal(t, row1.Brand, row2.Brand)
	require.Equal(t, row1.Link, row2.Link)
	require.Equal(t, row1.PrvPrice, row2.PrvPrice)
	require.Equal(t, arg.Price, row2.Price)
	require.WithinDuration(t, row1.CreatedAt, row2.CreatedAt, time.Second)
}

func TestDeleteOnSale(t *testing.T) {
	row1 := CreateRandomRowOnSale(t)
	err := testQueries.DeleteOnSale(context.Background(), row1.ID)
	require.NoError(t, err)

	account2, err := testQueries.GetOnSale(context.Background(), row1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account2)
}

func TestListOnSale(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomRowOnSale(t)
	}

	arg := ListOnSaleParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListOnSale(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}

func TestGetLengthOnSale(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomRowOnSale(t)
	}

	length, err := testQueries.GetLengthOnSale(context.Background())
	require.NoError(t, err)
	require.NotZero(t, length)

}

func TestGetOnSaleForUpdate(t *testing.T) {
	product := CreateRandomRowOnSale(t)
	product2, err := testQueries.GetOnSaleForUpdate(context.Background(), product.ID)
	require.NoError(t, err)
	require.NotEmpty(t, product2)

	require.Equal(t, product.ID, product2.ID)
	require.Equal(t, product.Brand, product2.Brand)
	require.Equal(t, product.Link, product2.Link)
	require.Equal(t, product.Price, product2.Price)
	require.WithinDuration(t, product.CreatedAt, product2.CreatedAt, time.Second)
}
