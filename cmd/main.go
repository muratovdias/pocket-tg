package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/muratovdias/pocket-tg/pkg/config"
	"github.com/muratovdias/pocket-tg/pkg/repository/boltdb"
	"github.com/muratovdias/pocket-tg/pkg/telegram"
	"github.com/zhashkevych/go-pocket-sdk"
	"log"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(cfg)

	bot, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		log.Panic(err.Error())
	}

	bot.Debug = true

	// Create pocket client
	pocketClient, err := pocket.NewClient(cfg.PocketConsumerKey)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Initialize BoltDB
	db, err := boltdb.InitDB()
	if err != nil {
		log.Fatal(err.Error())
	}
	boltRepo := boltdb.NewRepo(db)

	telegramBot := telegram.NewBot(bot, pocketClient, boltRepo, cfg.RedirectURL)
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err.Error())
	}
}
