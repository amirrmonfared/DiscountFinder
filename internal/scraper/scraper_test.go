package scrap

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gocolly/colly"
	_ "github.com/lib/pq"
)

var Ts = newTestServer()

var serverIndexResponse = []byte("hello world\n")

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

func TestCollectorOnHTML(t *testing.T) {
	ts := Ts
	defer ts.Close()

	c, err := Scraper(ts.URL, TestStore)
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
