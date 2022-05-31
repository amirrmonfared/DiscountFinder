package tools

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProductRemover(t *testing.T) {
	err := ProductRemover(testStore)
	require.NoError(t, err)
}

func TestOnSaleRemover(t *testing.T) {
	err := OnSaleRemover(testStore)
	require.NoError(t, err)
}
