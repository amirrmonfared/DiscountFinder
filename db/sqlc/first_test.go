package db

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/amirrmonfared/WebCrawler/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomRow(t *testing.T) First {
	arg := CreateFirstParams{
		Link:  util.RandomLink(),
		Price: util.RandomMoney(),
	}

	link, err := testQueries.CreateFirst(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, link)

	require.Equal(t, arg.Link, link.Link)
	require.Equal(t, arg.Price, link.Price)

	require.NotZero(t, link.ID)
	require.NotZero(t, link.CreatedAt)

	return link
}

func TestCreateRow(t *testing.T) {
	CreateRandomRow(t)
}

func TestGetRow(t *testing.T) {
	row1 := CreateRandomRow(t)
	row2, err := testQueries.GetFirst(context.Background(), row1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, row2)

	require.Equal(t, row1.ID, row2.ID)
	require.Equal(t, row1.Link, row2.Link)
	require.Equal(t, row1.Price, row2.Price)
	require.WithinDuration(t, row1.CreatedAt, row2.CreatedAt, time.Second)
}

func TestUpdateRow(t *testing.T) {
	row1 := CreateRandomRow(t)

	fmt.Println("row1:", row1.Price)
	arg := UpdateFirstParams{
		ID:    row1.ID,
		Price: util.RandomMoney(),
	}

	row2, err := testQueries.UpdateFirst(context.Background(), arg)
	fmt.Println("row2:", row2.Price)
	fmt.Println("row1 after:", row1.Price)
	require.NoError(t, err)
	require.NotEmpty(t, row2)

	require.Equal(t, row1.ID, row2.ID)
	require.Equal(t, row1.Link, row2.Link)
	require.Equal(t, row1.Price, row2.Price)
	require.WithinDuration(t, row1.CreatedAt, row2.CreatedAt, time.Second)
}

func TestDeleteRow(t *testing.T) {
	row1 := CreateRandomRow(t)
	err := testQueries.DeleteFirst(context.Background(), row1.ID)
	require.NoError(t, err)

	account2, err := testQueries.GetFirst(context.Background(), row1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account2)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomRow(t)
	}

	arg := ListFirstParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListFirst(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
