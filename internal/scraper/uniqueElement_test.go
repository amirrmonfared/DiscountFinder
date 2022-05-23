package scrap

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUniqueReview(t *testing.T) {
	fromFirst, _ := getInfoFromProduct(TestStore)
	for _, b := range fromFirst {

		productForReview := ProductForReview{
			ID:    b.ID,
			Brand: b.Brand,
			Link:  b.Link,
			Price: b.Price,
		}
		ProductsForReview = append(ProductsForReview, productForReview)
	}

	copy := copy(ProductsForReview, ProductsForReview)
	product, err := uniqueReview(ProductsForReview)
	require.NoError(t, err)
	require.NotEqual(t, product, copy)
}

func TestUniqueOnSale(t *testing.T) {
	onSale, _ := getInfoFromOnSale(TestStore)
	copy(onSale, onSale)
	product, err := uniqueOnSale(onSale)
	require.NoError(t, err)
	require.NotEmpty(t, product, onSale)
}
