package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/amirrmonfared/WebCrawler/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomRowSecond(t *testing.T) Second {
	arg := CreateSecondParams{
		Link:  util.RandomLink(),
		Price: util.RandomMoney(),
	}

	link, err := testQueries.CreateSecond(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, link)

	require.Equal(t, arg.Link, link.Link)
	require.Equal(t, arg.Price, link.Price)

	require.NotZero(t, link.ID)
	require.NotZero(t, link.CreatedAt)

	return link
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
	require.Equal(t, row1.Link, row2.Link)
	require.Equal(t, row1.Price, row2.Price)
	require.WithinDuration(t, row1.CreatedAt, row2.CreatedAt, time.Second)
}

func TestUpdateRowSecond(t *testing.T) {
	row1 := CreateRandomRowSecond(t)

	arg := UpdateSecondParams{
		ID:    row1.ID,
		Price: util.RandomMoney(),
	}

	row2, err := testQueries.UpdateSecond(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, row2)

	require.Equal(t, row1.ID, row2.ID)
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
