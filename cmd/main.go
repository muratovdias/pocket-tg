package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/muratovdias/pocket-tg/pkg/telegram"
	"log"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("6250725689:AAFD5w19F5mE3PE6EZUrD5leeeg8qTnVF_I")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	telegramBot := telegram.NewBot(bot)
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}
}
