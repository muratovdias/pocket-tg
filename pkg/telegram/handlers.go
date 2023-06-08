package telegram

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/zhashkevych/go-pocket-sdk"
	"log"
	url2 "net/url"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, b.messages.Responses.UnknownCommand)
	switch message.Command() {
	case "start":
		return b.handleStartCommand(message)
	}

	if _, err := b.bot.Send(msg); err != nil {
		return err
	}
	return nil
}

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	_, err := url2.ParseRequestURI(message.Text)
	if err != nil {
		return errInvalidLink
	}

	accessToken, err := b.getAccessToken(message.Chat.ID)
	if err != nil {
		return errUnauthorized
	}

	if err := b.pocketClient.Add(context.Background(), pocket.AddInput{
		AccessToken: accessToken,
		URL:         message.Text,
	}); err != nil {
		return errUnableToSave
	}
	// Send the message.
	msg := tgbotapi.NewMessage(message.Chat.ID, b.messages.Responses.SavedSuccessfully)
	if _, err := b.bot.Send(msg); err != nil {
		return err
	}
	return nil
}

func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, b.messages.Responses.AlreadyAuthorized)
	if _, err := b.getAccessToken(message.Chat.ID); err != nil {
		return b.initAuthorization(msg.ChatID)
	}

	if _, err := b.bot.Send(msg); err != nil {
		return err
	}
	return nil
}

func (b *Bot) createInlineKeyboard(chatID int64) (tgbotapi.InlineKeyboardMarkup, error) {
	link, err := b.generateAuthorizationLink(chatID)
	if err != nil {
		log.Println(err.Error())
		return tgbotapi.InlineKeyboardMarkup{}, err
	}
	clientKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL("Permission✔️", link),
		),
	)
	return clientKeyboard, nil
}
