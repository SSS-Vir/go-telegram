package gobot

import (
	"encoding/json"
	"fmt"
	"go-telegram/TGmodels"
)

type SendMessageParams struct {
	ChatId  int    `json:"chat_id"`
	Text    string `json:"text"`
	ReplyTo int    `json:"reply_to_message_id,omitempty"`
	// etc
}

const methodSendMessage = "/sendMessage"

func (bot *Bot) SendMessage(params SendMessageParams) (TGmodels.Message, error) {

	body, err := json.Marshal(params)
	if err != nil {
		return TGmodels.Message{}, err
	}

	var message TGmodels.Message
	err = makeRequest(bot.token, methodSendMessage, body, nil, &message)
	fmt.Printf("message: %+v\n", message)
	if err != nil {
		return TGmodels.Message{}, err
	}
	return message, nil
}
