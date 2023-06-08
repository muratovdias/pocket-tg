package telegram

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/muratovdias/pocket-tg/pkg/repository"
)

func (b *Bot) initAuthorization(chatID int64) error {
	msg := tgbotapi.NewMessage(chatID, b.messages.Responses.Start)
	keyboard, err := b.createInlineKeyboard(msg.ChatID)
	if err != nil {
		return err
	}
	msg.ReplyMarkup = keyboard

	if _, err := b.bot.Send(msg); err != nil {
		return err
	}
	return nil
}

func (b *Bot) getAccessToken(chatID int64) (string, error) {
	return b.tokenRepo.Get(chatID, repository.AccessToken)
}

func (b *Bot) generateAuthorizationLink(chatID int64) (string, error) {
	redirectURL := b.generateRedirectURL(chatID, b.redirectURL)
	requestToken, err := b.pocketClient.GetRequestToken(context.Background(), redirectURL)
	if err != nil {
		return "", fmt.Errorf("generateAuthorizationLink(): GetRequestToken(): %w", err)
	}

	err = b.tokenRepo.Save(chatID, requestToken, repository.RequestToken)
	if err != nil {
		return "", err
	}

	return b.pocketClient.GetAuthorizationURL(requestToken, redirectURL)
}

func (b *Bot) generateRedirectURL(chatID int64, redirectURL string) string {
	return fmt.Sprintf("%s?chat_id=%d", redirectURL, chatID)
}
