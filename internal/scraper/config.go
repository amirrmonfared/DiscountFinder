package scrap

var(
	Products = make([]Product, 0, 200)
	ProductsFromFirst = make([]ProductFromFirst, 0, 200)
	ProductsForReview = make([]ProductForReview, 0, 200)
	ProductsOnSale = make([]ProductOnSale, 0, 200)
)

type Product struct {
	Brand string `json:"brand"`
	Link  string `json:"link"`
	Price string `json:"price"`
}

type ProductFromFirst struct {
	ID    int64  `json:"id"`
	Brand string `json:"brand"`
	Link  string `json:"link"`
	Price string `json:"price"`
}

type ProductForReview struct {
	ID    int64  `json:"id"`
	Brand string `json:"brand"`
	Link  string `json:"link"`
	Price string `json:"price"`
}

type ProductOnSale struct {
	ID    int64  `json:"id"`
	Brand string `json:"brand"`
	Link  string `json:"link"`
	Price string `json:"price"`
}