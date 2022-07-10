package telegram

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	db "github.com/amirrmonfared/DiscountFinder/db/sqlc"
	"github.com/amirrmonfared/DiscountFinder/util"
	tele "gopkg.in/telebot.v3"
)

const IMTRUE = true

func Bot(store db.Store) {

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

	length, err := store.GetLengthOnSale(context.Background())
	if err != nil {
		log.Println(err)
	}

	b.Handle("/hello", func(ctx tele.Context) error {
		return ctx.Send("Hello!")
	})

	b.Handle("/discount", func(ctx tele.Context) error {
		if IMTRUE == true {
			for i := 0; i < int(length); i++ {
				list, err := getListOfOnSale(store, int32(i))
				if err != nil {
					log.Println(err)
					ctx.Send("error")
				}

				ctx.Send(list)
			}
		}
		return ctx.Send("Finish")
	})

	b.Start()
}

func getListOfOnSale(store db.Store, offset int32) (string, error) {
	var result []string

	list, err := store.ListOnSale(context.Background(), db.ListOnSaleParams{
		Limit:  10,
		Offset: offset,
	})
	if err != nil {
		log.Println(err)
	}

	for _, j := range list {
		result = append(result, "next", j.Brand, j.Link, j.Price)
	}

	strList := strings.Join(result, "  ")

	return strList, nil
}
