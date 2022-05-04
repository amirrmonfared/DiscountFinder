package telegram

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/amirrmonfared/DiscountFinder/util"
	tele "gopkg.in/telebot.v3"
)

const IMTRUE = true

func Bot(conn *sql.DB) {
	//	store := db.NewStore(conn)

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

	b.Handle("/hello", func(ctx tele.Context) error {
		return ctx.Send("Hello!")
	})

	b.Handle("/discount", func(ctx tele.Context) error {
		return ctx.Send("Finish")
	})

	b.Start()
}
