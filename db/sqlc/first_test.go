package db

import (
	"context"
	"testing"

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


