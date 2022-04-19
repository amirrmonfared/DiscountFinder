package main

import (
	"database/sql"
	"fmt"
	"log"

	scrap "github.com/amirrmonfared/DiscountFinder/internal/scraper"
	"github.com/amirrmonfared/DiscountFinder/util"
	"github.com/gocolly/colly"
	_ "github.com/lib/pq"
)

const webPage = "https://www.trendyol.com/elbise-x-c56"

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	fmt.Println("connected to server on port 8080")
	fmt.Println("--------------------------------------")

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	fmt.Println("connected to database")
	fmt.Println("--------------------------------------")

	//scrap.Scraper2(conn)
	//scrap.DiscountFinder(conn)
	// store := db.NewStore(conn)
	// server := api.NewServer(store)

	go run(webPage, conn)

	// err = server.Start(config.ServerAddress)
	// if err != nil {
	// 	log.Fatal("cannot start server:", err)
	// }

	defer conn.Close()

}

func run(webPage string, conn *sql.DB) (*colly.Collector, error) {
	scrap, err := scrap.Scraper(webPage, conn)
	if err != nil {
		fmt.Println(err)
	}

	return scrap, nil
}
