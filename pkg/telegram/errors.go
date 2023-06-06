package telegram

import "errors"

var (
	errInvalidLink  = errors.New("link is invalid")
	errUnauthorized = errors.New("you are not authorized")
	errUnableToSave = errors.New("unable to save")
)

func (b *Bot) handleError(chatID int64, err error) {
	switch err {

	}
}
