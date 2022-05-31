package tools

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUniqueReview(t *testing.T) {
	CreateRandomProduct(t)
	fromFirst, _ := GetInfoFromProduct(testStore)
	copy(fromFirst, fromFirst)
	product, err := UniqueReview(fromFirst)
	require.NoError(t, err)
	require.NotEmpty(t, product, fromFirst)
}

func TestUniqueOnSale(t *testing.T) {
	CreateRandomRowOnSale(t)
	onSale, _ := GetInfoFromOnSale(testStore)
	copy(onSale, onSale)
	product, err := UniqueOnSale(onSale)
	require.NoError(t, err)
	require.NotEmpty(t, product, onSale)
}
