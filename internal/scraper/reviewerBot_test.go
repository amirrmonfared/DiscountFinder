package scrap

import (
	"testing"

	mockdb "github.com/amirrmonfared/DiscountFinder/db/mock"
	db "github.com/amirrmonfared/DiscountFinder/db/sqlc"
	"github.com/amirrmonfared/DiscountFinder/util"
	"github.com/gocolly/colly"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestReviewerBotConfig(t *testing.T) {
	c, err := reviewerBotConfig()
	require.NoError(t, err)

	ts := Ts
	defer ts.Close()

	c.OnHTML("title", func(e *colly.HTMLElement) {
		if e.Text != "Test Page" {
			t.Error("Title element text does not match, got", e.Text)
		}
	})

	c.OnHTML("p", func(e *colly.HTMLElement) {
		if e.Attr("class") != "description" {
			t.Error("Failed to get paragraph's class attribute")
		}
	})

	c.OnHTML("body", func(e *colly.HTMLElement) {
		if e.ChildAttr("p", "class") != "description" {
			t.Error("Invalid class value")
		}
		classes := e.ChildAttrs("p", "class")
		if len(classes) != 2 {
			t.Error("Invalid class values")
		}
	})

	c.Visit(ts.URL + "/html")
}

func TestReviewerBot(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdb.NewMockStore(ctrl)
	store.EXPECT().GetLengthOfProducts(gomock.Any())
	store.EXPECT().ListProduct(gomock.Any(), gomock.Any())
	err := ReviewerBot(store)
	require.NoError(t, err)
}

func TestHtmlCollector(t *testing.T) {
	//TODO improve test and not connect to internet
	product := db.Product{
		Link: "https://www.trendyol.com/us-polo-assn/misu-2fx-lacivert-erkek-terlik-p-201096857?boutiqueId=595230&merchantId=107040",
	}
	collector, err := reviewerBotConfig()
	require.NoError(t, err)
	products, err := htmlCollector(collector, product)
	require.NoError(t, err)
	require.NotEmpty(t, products)

}

func TestDiscountFinder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdb.NewMockStore(ctrl)

	product1 := db.Product{
		Brand: util.RandomString(3),
		Link:  util.RandomLink(),
		Price: util.RandomPriceString(3),
	}
	product2 := db.Product{
		Brand: util.RandomString(3),
		Link:  util.RandomLink(),
		Price: product1.Price + util.RandomPriceString(1),
	}

	products := toCheckPrice{reviewedProduct: product2, productForReview: product1}

	err := discountFinder(products, store)
	require.NoError(t, err)
}

func TestStoreOnSale(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdb.NewMockStore(ctrl)
	store.EXPECT().StoreOnSale(gomock.Any(), gomock.Any())

	onSale := db.OnSale{
		Brand: util.RandomString(3),
		Link:  util.RandomLink(),
		Price: util.RandomPriceString(3),
	}
	err := storeOnSale(store, onSale)
	require.NoError(t, err)
}
