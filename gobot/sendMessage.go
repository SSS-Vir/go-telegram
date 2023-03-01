package gobot

import (
	"encoding/json"
	"go-telegram/TGmodels"
)

type SendMessageParams struct {
	ChatId  int    `json:"chat_id"`
	Text    string `json:"text"`
	ReplyTo int    `json:"reply_to_message_id,omitempty"`
	// etc
}

const methodSendMessage = "/sendMessage"

func (bot *Bot) SendMessage(params SendMessageParams) (TelegramResponse[TGmodels.Message], error) {

	body, err := json.Marshal(params)
	if err != nil {
		return TelegramResponse[TGmodels.Message]{}, err
	}

	var message TelegramResponse[TGmodels.Message]
	err = bot.makeRequest(methodSendMessage, body, nil, &message)
	if err != nil {
		return TelegramResponse[TGmodels.Message]{}, err
	}
	return message, nil
}
