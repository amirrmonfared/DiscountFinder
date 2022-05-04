package telegram

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	db "github.com/amirrmonfared/DiscountFinder/db/sqlc"
	"github.com/amirrmonfared/DiscountFinder/util"
	tele "gopkg.in/telebot.v3"
)

const IMTRUE = true

func Bot(conn *sql.DB) {
	store := db.NewStore(conn)

	config, err := util.LoadConfig(".")
	if err != nil {
		fmt.Println("cannot get config:", err)
	}

	pref := tele.Settings{
		Token:  config.TelegramBotToken,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Println(err)
		return
	}

	length, _ := store.GetLengthOnSale(context.Background())

	b.Handle("/hello", func(ctx tele.Context) error {
		return ctx.Send("Hello!")
	})

	b.Handle("/discount", func(ctx tele.Context) error {
		if IMTRUE == true {
			for i := 0; i < int(length); i++ {
				list, _ := getListOfOnSale(conn, int32(i))

				ctx.Send(list)
			} 
		}
		return ctx.Send("Finish")
	})

	b.Start()
}

func getListOfOnSale(conn *sql.DB, offset int32) (string, error) {
	var result []string
	store := db.NewStore(conn)

	list, _ := store.ListOnSale(context.Background(), db.ListOnSaleParams{
		Limit: 10,
		Offset: offset,
	})

	for _, j := range list{
		result = append(result,"next", j.Brand, j.Link, j.Price)
	}

	strList := strings.Join(result, "  ")

	return strList, nil
} 
