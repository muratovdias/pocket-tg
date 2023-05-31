package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Извините, я не знаю такой команды :(")
	switch message.Command() {
	case "start":
		if err := b.handleStartCommand(&msg); err != nil {
			return err
		}
	}
	if _, err := b.bot.Send(msg); err != nil {
		return err
	}

	return nil
}

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	// Send the message.
	if _, err := b.bot.Send(msg); err != nil {
		return err
	}
	return nil
}

func (b *Bot) handleStartCommand(msg *tgbotapi.MessageConfig) error {
	msg.Text = "Welcome 🤗\nI will be helping you to save links into your Pocket account.\nFor this, follow the link and give me permission."
	keyboard, err := b.createInlineKeyboard(msg)
	if err != nil {
		return err
	}
	msg.ReplyMarkup = keyboard
	return nil
}

func (b *Bot) createInlineKeyboard(msg *tgbotapi.MessageConfig) (tgbotapi.InlineKeyboardMarkup, error) {
	link, err := b.generateAuthorizationLink(msg.ChatID)
	if err != nil {
		return tgbotapi.InlineKeyboardMarkup{}, fmt.Errorf("createInlineKeyboard(): %w", err)
	}
	clientKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL("Permission", link),
		),
	)
	return clientKeyboard, nil
}
