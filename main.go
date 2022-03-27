package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/amirrmonfared/WebCrawler/internal/crawler"
	"github.com/amirrmonfared/WebCrawler/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	fmt.Println("connected to database:")

	defer conn.Close()

	crawler.Scraper()

}
