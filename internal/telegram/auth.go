package telegram

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	scrap "github.com/amirrmonfared/DiscountFinder/internal/scraper"
	"github.com/amirrmonfared/DiscountFinder/util"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Bot(conn *sql.DB) {

	config, err := util.LoadConfig(".")
	if err != nil {
		fmt.Println("cannot get config:", err)
	}

	bot, err := tgbotapi.NewBotAPI(config.TelegramBotToken)
	if err != nil {
		log.Panic("cannot get bot API", err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	x, _ := scrap.GetInfoFromOnSaleInString(conn)
	x2 := strings.Join(x, " ")
	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages
			continue
		}

		// Create a new MessageConfig. We don't have text yet,
		// so we leave it empty.
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		// Extract the command from the Message.
		switch update.Message.Command() {
		case "help":
			msg.Text = "for see discounted /discount"
		case "discount":
			msg.Text = x2
		default:
			msg.Text = "I don't know that command"
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
