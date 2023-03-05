package gobot

import (
	"go-telegram/TGmodels"
)

const methodGetMe = "/getMe"

func (bot *Bot) GetMe() (TGmodels.User, error) {
	var user TGmodels.User
	err := makeRequest(bot.token, methodGetMe, nil, nil, &user)
	if err != nil {
		return TGmodels.User{}, err
	}
	return user, nil
}
