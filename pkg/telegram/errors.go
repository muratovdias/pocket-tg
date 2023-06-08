package telegram

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

var (
	errInvalidLink  = errors.New("link is invalid")
	errUnauthorized = errors.New("user is not authorized")
	errUnableToSave = errors.New("unable to save")
)

func (b *Bot) handleError(chatID int64, err error) {
	msg := tgbotapi.NewMessage(chatID, b.messages.Errors.Default)
	switch err {
	case errUnauthorized:
		msg.Text = b.messages.Errors.Unauthorized
	case errInvalidLink:
		msg.Text = b.messages.Errors.InvalidLink
	case errUnableToSave:
		msg.Text = b.messages.Errors.UnableToSave
	}
	if _, err := b.bot.Send(msg); err != nil {
		log.Println(err.Error())
		return
	}
}
