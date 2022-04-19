package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/amirrmonfared/DiscountFinder/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomRowSecond(t *testing.T) Second {
	arg := CreateSecondParams{
		Brand: util.RandomString(5),
		Link:  util.RandomLink(),
		Price: util.RandomPriceString(4),
	}

	product, err := testQueries.CreateSecond(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, product)

	require.Equal(t, arg.Link, product.Link)
	require.Equal(t, arg.Price, product.Price)

	require.NotZero(t, product.ID)
	require.NotZero(t, product.CreatedAt)

	return product
}

func TestCreateRowSecond(t *testing.T) {
	CreateRandomRowSecond(t)
}

func TestGetRowSecond(t *testing.T) {
	row1 := CreateRandomRowSecond(t)
	row2, err := testQueries.GetSecond(context.Background(), row1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, row2)

	require.Equal(t, row1.ID, row2.ID)
	require.Equal(t, row1.Brand, row2.Brand)
	require.Equal(t, row1.Link, row2.Link)
	require.Equal(t, row1.Price, row2.Price)
	require.WithinDuration(t, row1.CreatedAt, row2.CreatedAt, time.Second)
}

func TestUpdateRowSecond(t *testing.T) {
	row1 := CreateRandomRowSecond(t)

	arg := UpdateSecondParams{
		ID:    row1.ID,
		Price: util.RandomPriceString(4),
	}

	row2, err := testQueries.UpdateSecond(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, row2)

	require.Equal(t, row1.ID, row2.ID)
	require.Equal(t, row1.Brand, row2.Brand)
	require.Equal(t, row1.Link, row2.Link)
	require.Equal(t, arg.Price, row2.Price)
	require.WithinDuration(t, row1.CreatedAt, row2.CreatedAt, time.Second)
}

func TestDeleteRowSecond(t *testing.T) {
	row1 := CreateRandomRowSecond(t)
	err := testQueries.DeleteSecond(context.Background(), row1.ID)
	require.NoError(t, err)

	account2, err := testQueries.GetSecond(context.Background(), row1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account2)
}

func TestListAccountsSecond(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomRowSecond(t)
	}

	arg := ListSecondParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListSecond(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}

func TestGetLengthOfSecond(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomRowSecond(t)
	}

	length, err := testQueries.GetLengthOfSecond(context.Background())
	require.NoError(t, err)
	require.NotZero(t, length)

}

func TestGetSecondForUpdate(t *testing.T) {
	product := CreateRandomRowSecond(t)
	product2, err := testQueries.GetSecondForUpdate(context.Background(), product.ID)
	require.NoError(t, err)
	require.NotEmpty(t, product2)

	require.Equal(t, product.ID, product2.ID)
	require.Equal(t, product.Brand, product2.Brand)
	require.Equal(t, product.Link, product2.Link)
	require.Equal(t, product.Price, product2.Price)
	require.WithinDuration(t, product.CreatedAt, product2.CreatedAt, time.Second)
}
