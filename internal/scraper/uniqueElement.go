package scrap

func uniqueReview(productSlice []ProductForReview) []ProductForReview {
	keys := make(map[ProductForReview]bool)
	list := []ProductForReview{}
	for _, entry := range productSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func uniqueOnSale(productSlice []ProductOnSale) []ProductOnSale {
	keys := make(map[ProductOnSale]bool)
	list := []ProductOnSale{}
	for _, entry := range productSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
