package gobot

import (
	"strings"
	"time"
)

type Bot struct {
	IBot
	token           string
	lastUpdateId    int
	handlers        map[string]HandlerFunc
	commandHandlers map[string]map[string]CommandHandlerFunc // command -> id -> handler
}

func (bot *Bot) RemoveHandler(handlerId string) bool {
	_, ok := bot.handlers[handlerId]
	delete(bot.handlers, handlerId)
	return ok
}

func (bot *Bot) AddHandler(handler HandlerFunc) string {
	handlerId := randString(32)
	bot.handlers[handlerId] = handler
	return handlerId
}

func (bot *Bot) AddCommandHandler(command string, handler CommandHandlerFunc) string {
	handlerId := randString(32)
	_, ok := bot.commandHandlers[command]
	if !ok {
		bot.commandHandlers[command] = map[string]CommandHandlerFunc{}
	}
	bot.commandHandlers[command][handlerId] = handler
	return handlerId
}

func (bot *Bot) RemoveCommandHandler(command string, handlerId string) bool {
	_, ok := bot.commandHandlers[command][handlerId]
	delete(bot.commandHandlers[command], handlerId)
	return ok
}

func New(token string) *Bot {
	return &Bot{
		token:           token,
		lastUpdateId:    0,
		handlers:        map[string]HandlerFunc{},
		commandHandlers: map[string]map[string]CommandHandlerFunc{},
	}
}

func (bot *Bot) Run() {

	for {
		updates, err := getUpdates(bot, GetUpdatesParams{Offset: bot.lastUpdateId})
		if err != nil {
			println(err.Error())
			goto SLEEP
		}

		for _, update := range updates {
			if len(update.Message.Text) != 0 && strings.HasPrefix(update.Message.Text, "/") {
				for command, handlers := range bot.commandHandlers {
					split := strings.Split(update.Message.Text, " ")
					if strings.HasPrefix(split[0], command) {
						for _, handler := range handlers {
							handler(bot, &update, split[1:])
						}
					}
				}
				goto SLEEP
			}

			for _, handler := range bot.handlers {
				handler(bot, &update)
			}
		}
	SLEEP:
		if len(updates) != 0 {
			bot.lastUpdateId = updates[len(updates)-1].UpdateId + 1
		}
		time.Sleep(time.Millisecond * 1000)
	}

}
