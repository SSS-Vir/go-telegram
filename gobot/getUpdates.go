package gobot

import (
	"errors"
	"go-telegram/TGmodels"
	"net/url"
	"strconv"
)

const methodGetUpdates = "/getUpdates?"

type GetUpdatesParams struct {
	Offset int
	// etc
}

func getUpdates(bot *Bot, params GetUpdatesParams) (TelegramResponse[[]TGmodels.Update], error) {
	queryParams := url.Values{
		"offset": {strconv.FormatInt(int64(params.Offset), 10)},
	}
	var response TelegramResponse[[]TGmodels.Update]
	err := bot.makeRequest(methodGetUpdates, nil, queryParams, &response)
	if err != nil {
		return TelegramResponse[[]TGmodels.Update]{}, err
	}
	if !response.Ok {
		return TelegramResponse[[]TGmodels.Update]{}, errors.New(response.Description)
	}
	return response, nil
}
