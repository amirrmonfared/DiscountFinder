package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/amirrmonfared/DiscountFinder/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomRow(t *testing.T) First {
	arg := CreateFirstProductParams{
		Brand: util.RandomString(5),
		Link:  util.RandomLink(),
		Price: util.RandomPriceString(4),
	}

	product, err := testQueries.CreateFirstProduct(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, product)

	require.Equal(t, arg.Brand, product.Brand)
	require.Equal(t, arg.Link, product.Link)
	require.Equal(t, arg.Price, product.Price)

	require.NotZero(t, product.ID)
	require.NotZero(t, product.CreatedAt)

	return product
}

func TestCreateRow(t *testing.T) {
	CreateRandomRow(t)
}

func TestGetFirstProduct(t *testing.T) {
	row1 := CreateRandomRow(t)
	row2, err := testQueries.GetFirstProduct(context.Background(), row1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, row2)

	require.Equal(t, row1.ID, row2.ID)
	require.Equal(t, row1.Brand, row2.Brand)
	require.Equal(t, row1.Link, row2.Link)
	require.Equal(t, row1.Price, row2.Price)
	require.WithinDuration(t, row1.CreatedAt, row2.CreatedAt, time.Second)
}

func TestUpdateFirstProduct(t *testing.T) {
	row1 := CreateRandomRow(t)

	arg := UpdateFirstProductParams{
		ID:    row1.ID,
		Price: util.RandomPriceString(4),
	}

	row2, err := testQueries.UpdateFirstProduct(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, row2)

	require.Equal(t, row1.ID, row2.ID)
	require.Equal(t, row1.Brand, row2.Brand)
	require.Equal(t, row1.Link, row2.Link)
	require.Equal(t, arg.Price, row2.Price)
	require.WithinDuration(t, row1.CreatedAt, row2.CreatedAt, time.Second)
}

func TestDeleteFirstProduct(t *testing.T) {
	row1 := CreateRandomRow(t)
	err := testQueries.DeleteFirstProduct(context.Background(), row1.ID)
	require.NoError(t, err)

	account2, err := testQueries.GetFirstProduct(context.Background(), row1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account2)
}

func TestListFirst(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomRow(t)
	}

	arg := ListFirstProductParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListFirstProduct(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}

func TestGetLengthOfFirst(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomRow(t)
	}

	length, err := testQueries.GetLengthOfFirst(context.Background())
	require.NoError(t, err)
	require.NotZero(t, length)

}

func TestGetFirstProductForUpdate(t *testing.T) {
	product := CreateRandomRow(t)
	product2, err := testQueries.GetFirstProductForUpdate(context.Background(), product.ID)
	require.NoError(t, err)
	require.NotEmpty(t, product2)

	require.Equal(t, product.ID, product2.ID)
	require.Equal(t, product.Brand, product2.Brand)
	require.Equal(t, product.Link, product2.Link)
	require.Equal(t, product.Price, product2.Price)
	require.WithinDuration(t, product.CreatedAt, product2.CreatedAt, time.Second)
}
