package scrap

import (
	"fmt"
	"testing"

	"github.com/gocolly/colly"
	"github.com/stretchr/testify/require"
)

func TestDiscountFinder(t *testing.T) {
	err := DiscountFinder(TestStore)
	require.NoError(t, err)
}

func TestCollectorOnHTML2(t *testing.T) {
	ts := Ts
	defer ts.Close()

	_, _, c, err := collector(TestStore)
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
