package main

import (
	"database/sql"
	"fmt"
	"log"

	db "github.com/amirrmonfared/WebCrawler/db/sqlc"
	"github.com/amirrmonfared/WebCrawler/util"
	_ "github.com/lib/pq"
)

const webPage = "https://www.trendyol.com/erkek-t-shirt-x-g2-c73?pi=2"

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

	db.Scraper(webPage, conn)

	//	store := db.NewStore(conn)
	// server := api.NewServer(store)

	// err = server.Start(config.ServerAddress)
	// if err != nil {
	// 	log.Fatal("cannot start server:", err)
	// }

	defer conn.Close()

}
