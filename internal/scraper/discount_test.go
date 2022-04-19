package scrap

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDiscountFinder(t *testing.T) {
	getOnSale, err := getInfoFromOnSale(testDB)
	require.NoError(t, err)

	discount, err := DiscountFinder(testDB)
	require.NoError(t, err)
	require.Equal(t, getOnSale, discount)
}
