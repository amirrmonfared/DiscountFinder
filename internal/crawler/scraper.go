package crawler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func Scraper(href string) {
	resp, err := http.Get(href)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalf("failed to fetch data: %d %s", resp.StatusCode, resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	title := doc.Find("span.prc-dsc").Text()

	fix := fixPrice(title)

	fmt.Println(fix)
}

// fixPrice giving a htmltag from scrapper and turn it into a []float 64
func fixPrice(htmlTag string) []float64 {

	trimmed := strings.Trim(htmlTag, "TL")

	truePrice := strings.Split(trimmed, ",")

	ints := make([]float64, len(truePrice))

	for i, s := range truePrice {
		ints[i], _ = strconv.ParseFloat(s, 64)
	}

	price := ints[:1]

	return price
}
