package gobot

import (
	"errors"
	"go-telegram/TGmodels"
)

const methodGetMe = "/getMe"

func (bot *Bot) GetMe() (TelegramResponse[TGmodels.User], error) {
	var user TelegramResponse[TGmodels.User]
	err := bot.makeRequest(methodGetMe, nil, nil, &user)
	if err != nil {
		return TelegramResponse[TGmodels.User]{}, err
	}
	if !user.Ok {
		return user, errors.New(user.Description)
	}
	return user, nil
}
