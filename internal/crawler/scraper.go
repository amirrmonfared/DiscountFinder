package crawler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func Scraper(href string) float64 {
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

	fix, err := GetPrice(title)
	if err != nil {
		fmt.Println("cannot get price", err)
	}
	
	fmt.Println(fix)

	return fix
}

// GetPrice giving a htmltag from scrapper and turn it into a float64
func GetPrice(htmlTag string) (float64, error) {

	trimmed := strings.Trim(htmlTag, "TL")

	truePrice := strings.Split(trimmed, ",")

	priceSlice := make([]float64, len(truePrice))

	for i, s := range truePrice {
		priceSlice[i], _ = strconv.ParseFloat(s, 64)
	}

	price := priceSlice[0]

	return price, nil
}
