package scrap

import db "github.com/amirrmonfared/DiscountFinder/db/sqlc"

// uniqueReview find duplicate products and remove them
func uniqueReview(productSlice []db.Product) ([]db.Product, error) {
	keys := make(map[db.Product]bool)
	list := []db.Product{}
	for _, entry := range productSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list, nil
}

// uniqueOnSale find duplicate products and remove them
func uniqueOnSale(productSlice []db.OnSale) ([]db.OnSale, error) {
	keys := make(map[db.OnSale]bool)
	list := []db.OnSale{}
	for _, entry := range productSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list, nil
}
