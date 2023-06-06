package telegram

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/zhashkevych/go-pocket-sdk"
	"log"
	url2 "net/url"
)

const (
	AlreadyAuthorized = "You are already authorized"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Sorry, I don't know that command :(")
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
	msg := tgbotapi.NewMessage(message.Chat.ID, "Link saved successfully ✅")

	_, err := url2.ParseRequestURI(message.Text)
	if err != nil {
		msg.Text = "Invalid link"
		if _, err := b.bot.Send(msg); err != nil {
			return err
		}
		return nil
	}

	accessToken, err := b.getAccessToken(msg.ChatID)
	if err != nil {
		msg.Text = "You are unautorized, please use command /start"
		if _, err := b.bot.Send(msg); err != nil {
			return err
		}
		return nil
	}

	if err := b.pocketClient.Add(context.Background(), pocket.AddInput{
		AccessToken: accessToken,
		URL:         message.Text,
	}); err != nil {
		msg.Text = "Sorry, couldn't save the link, try again later"
		if _, err := b.bot.Send(msg); err != nil {
			return err
		}
		return nil
	}
	// Send the message.
	if _, err := b.bot.Send(msg); err != nil {
		return err
	}
	return nil
}

func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, AlreadyAuthorized)
	if _, err := b.getAccessToken(message.Chat.ID); err != nil {
		fmt.Println("authorization")
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
