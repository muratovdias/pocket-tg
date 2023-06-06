package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/muratovdias/pocket-tg/pkg/repository"
	"github.com/zhashkevych/go-pocket-sdk"
	"log"
)

type Bot struct {
	bot          *tgbotapi.BotAPI
	pocketClient *pocket.Client
	tokenRepo    repository.TokenRepo
	redirectURL  string
}

func NewBot(bot *tgbotapi.BotAPI, pocketClient *pocket.Client, tr repository.TokenRepo, redirectURL string) *Bot {
	return &Bot{
		bot:          bot,
		pocketClient: pocketClient,
		tokenRepo:    tr,
		redirectURL:  redirectURL,
	}
}
func (b *Bot) Start() error {
	log.Printf("Authorized on account %s", b.bot.Self.UserName)

	updates := b.initUpdatesChannel()

	if err := b.handleUpdates(updates); err != nil {
		return err
	}
	return nil
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) error {
	for update := range updates {

		if update.Message == nil { // If we got a message
			continue
		}

		if update.Message.IsCommand() {
			if err := b.handleCommand(update.Message); err != nil {
				log.Println("handleCommand(): ", err.Error())
				return err
			}
			continue
		}

		if err := b.handleMessage(update.Message); err != nil {
			return err
		}
	}
	return nil
}

func (b *Bot) initUpdatesChannel() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.bot.GetUpdatesChan(u)

	return updates
}
