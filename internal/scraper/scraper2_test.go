package scrap

import (
	"context"
	"fmt"
	"testing"

	db "github.com/amirrmonfared/DiscountFinder/db/sqlc"
	"github.com/amirrmonfared/DiscountFinder/util"
	"github.com/gocolly/colly"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

type TestLinkProduct struct {
	Brand string `json:"brand"`
	Link string `json:"link"`
}

var TestLinkProducts = make([]LinkProduct, 0, 200)

func CreateRandomRow() db.First {
	return db.First{
		ID:    util.RandomInt(1, 5),
		Brand: util.RandomString(5),
		Link:  util.RandomLink(),
		Price: util.RandomPriceString(5),
	}
}

func TestCreateRow(t *testing.T) {
	CreateRandomRow()
}

func TestCollectorOnHTML2(t *testing.T) {
	ts := Ts
	defer ts.Close()

	c, err := Scraper2(testDB)
	if err != nil {
		fmt.Println(err)
	}

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

func TestGetInfo(t *testing.T) {

	for i := 0; i < 10; i++ {
		CreateRandomRow()
	}

	testLength, err := testQueries.GetLengthOfFirst(context.Background())
	require.NoError(t, err)
	require.NotZero(t, testLength)

	arg := db.ListFirstProductParams{
		Limit: int32(testLength),
		Offset: 0,
	}

	testListFirst, err := testQueries.ListFirstProduct(context.Background(), arg)
	require.NoError(t, err)
	require.NotZero(t, testListFirst)

	for _, a := range testListFirst {
		testLinkProducts := TestLinkProduct{
			Brand: a.Brand,
			Link: a.Link,
		}

		TestLinkProducts = append(TestLinkProducts, LinkProduct(testLinkProducts))
	}

	Info, err := getInfo(testDB)
	require.NoError(t, err)
	require.Equal(t, Info, TestLinkProducts)
	
}
