package scrap

func uniqueReview(productSlice []ProductForReview) ([]ProductForReview, error) {
	keys := make(map[ProductForReview]bool)
	list := []ProductForReview{}
	for _, entry := range productSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list, nil
}

func uniqueOnSale(productSlice []ProductOnSale) ([]ProductOnSale, error) {
	keys := make(map[ProductOnSale]bool)
	list := []ProductOnSale{}
	for _, entry := range productSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list, nil
}
