package scrap

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetInfoFromFirst(t *testing.T) {
	info, err := getInfoFromFirst(testDB)
	require.NoError(t, err)
	require.NotEmpty(t, info)
}

func TestGetInfoFromSecond(t *testing.T) {
	info, err := getInfoFromSecond(testDB)
	require.NoError(t, err)
	require.NotEmpty(t, info)
}
