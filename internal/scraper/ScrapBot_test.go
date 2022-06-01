package scrap

import (
	"net/http"
	"net/http/httptest"
	"testing"

	db "github.com/amirrmonfared/DiscountFinder/db/sqlc"
	"github.com/amirrmonfared/DiscountFinder/util"
	"github.com/gocolly/colly"
	"github.com/stretchr/testify/require"
)

var (
	serverIndexResponse = []byte("hello world\n")
	Ts                  = newTestServer()
)

func newTestServer() *httptest.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write(serverIndexResponse)
	})

	mux.HandleFunc("/html", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`<!DOCTYPE html>
			<html>
			<head>
			<title>Test Page</title>
			</head>
			<body>
			<h1>Hello World</h1>
			<p class="description">This is a test page</p>
			<p class="description">This is a test paragraph</p>
			</body>
			</html>
		`))
	})

	return httptest.NewServer(mux)
}
func TestScrapBotConfig(t *testing.T) {
	c, err := scrapBotConfig()
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

func TestScrapBot(t *testing.T) {
	ts := Ts
	defer ts.Close()

	err := ScrapBot(ts.URL+"/html", testStore)
	require.NoError(t, err)
}

func TestStoreProduct(t *testing.T) {
	p := db.Product{
		Brand: util.RandomString(5),
		Link:  util.RandomLink(),
		Price: util.RandomPriceString(4),
	}

	result, err := storeProduct(testStore, p)
	require.NoError(t, err)
	require.Equal(t, p.Brand, result.Product.Brand)
	require.Equal(t, p.Link, result.Product.Link)
	require.Equal(t, p.Price, result.Product.Price)
}

func TestRemoveProductFromSlice(t *testing.T) {
	p := db.Product{
		Brand: util.RandomString(5),
		Link:  util.RandomLink(),
		Price: util.RandomPriceString(4),
	}
	Products = append(Products, p)
	product, err := removeProductFromSlice()
	require.NoError(t, err)
	require.Equal(t, product, p)
	require.Empty(t, Products)
}
