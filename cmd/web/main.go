package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/amirrmonfared/DiscountFinder/api"
	db "github.com/amirrmonfared/DiscountFinder/db/sqlc"
	"github.com/amirrmonfared/DiscountFinder/util"
	"github.com/jasonlvhit/gocron"
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

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		fmt.Println("cannot connect to server", err)
	}

	go RunScrap(webPage, conn)
	go RunBot(conn)

	gocron.Every(2).Minutes().Do(RunScrap, webPage, conn)
	gocron.Every(1).Minutes().Do(RunDiscountFinder, conn)
	gocron.Every(1).Minutes().Do(RunRemoveFirst, conn)
	gocron.Every(1).Minutes().Do(RunRemoveOnSale, conn)

	<- gocron.Start()

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

	defer conn.Close()

}
