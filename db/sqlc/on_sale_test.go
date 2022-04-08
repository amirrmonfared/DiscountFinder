package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/amirrmonfared/WebCrawler/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomRowOnSale(t *testing.T) OnSale {
	arg := CreateOnSaleParams{
		Brand: util.RandomString(5),
		Link:  util.RandomLink(),
	//TODO: add onsale percentage	
		Price: util.RandomPriceString(4),
	}

	product, err := testQueries.CreateOnSale(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, product)

	require.Equal(t, arg.Brand, product.Brand)
	require.Equal(t, arg.Link, product.Link)
	require.Equal(t, arg.Price, product.Price)

	require.NotZero(t, product.ID)
	require.NotZero(t, product.CreatedAt)

	return product
}

func TestCreateRowOnSale(t *testing.T) {
	CreateRandomRow(t)
}

func TestGetOnSale(t *testing.T) {
	row1 := CreateRandomRowOnSale(t)
	row2, err := testQueries.GetOnSale(context.Background(), row1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, row2)

	require.Equal(t, row1.ID, row2.ID)
	require.Equal(t, row1.Brand, row2.Brand)
	require.Equal(t, row1.Link, row2.Link)
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
