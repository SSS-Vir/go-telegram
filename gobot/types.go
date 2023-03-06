package gobot

import (
	"bytes"
	"encoding/json"
	"errors"
	"go-telegram/TGmodels"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
)

const (
	telegramApiUrl = "https://api.telegram.org/bot"
	letterBytes    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

type HandlerFunc func(bot IBot, update *TGmodels.Update)
type CommandHandlerFunc func(bot IBot, update *TGmodels.Update, arguments []string)

type IBot interface {
	RemoveHandler(handlerId string) bool
	AddHandler(handler HandlerFunc) string
	AddCommandHandler(command string, handler CommandHandlerFunc) string
	RemoveCommandHandler(command string, handlerId string) bool

	// Telegram Methods
	GetMe() (TGmodels.User, error)
	SendMessage(params SendMessageParams) (TGmodels.Message, error)
}

type TelegramResponse[T TGmodels.TelegramModel] struct {
	Ok          bool   `json:"ok"`
	ErrorCode   int    `json:"error_code,omitempty"`
	Description string `json:"description,omitempty"`
	Result      T      `json:"result,omitempty"`
}

func randString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}

func makeRequest[T TGmodels.TelegramModel](token string, methodName string, body []byte, params url.Values, returnValue *T) error {
	paramsQuery := params.Encode()
	response, err := http.Post(telegramApiUrl+token+methodName+paramsQuery, "application/json", bytes.NewReader(body))
	if err != nil {
		return err
	}
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	var js TelegramResponse[T]

	err = json.Unmarshal(bodyBytes, &js)

	if err != nil {
		return err
	}
	if !js.Ok {
		return errors.New(strconv.FormatInt(int64(js.ErrorCode), 10) + ": " + js.Description)
	}
	*returnValue = js.Result
	return nil
}
