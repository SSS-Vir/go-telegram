package gobot

import (
	"go-telegram/TGmodels"
	"net/url"
	"strconv"
)

const methodGetUpdates = "/getUpdates?"

type GetUpdatesParams struct {
	Offset int
	// etc
}

func getUpdates(bot *Bot, params GetUpdatesParams) ([]TGmodels.Update, error) {
	queryParams := url.Values{
		"offset": {strconv.FormatInt(int64(params.Offset), 10)},
	}
	var response []TGmodels.Update
	err := makeRequest(bot.token, methodGetUpdates, nil, queryParams, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
