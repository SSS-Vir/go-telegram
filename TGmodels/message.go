package TGmodels

type Message struct {
	TelegramModel
	MessageId       int    `json:"message_id"`
	MessageThreadId int    `json:"message_thread_id,omitempty"`
	From            User   `json:"from,omitempty"`
	SenderChat      Chat   `json:"sender_chat,omitempty"`
	Date            int    `json:"date"`
	Chat            Chat   `json:"chat"`
	Text            string `json:"text,omitempty"`
	// etc
}
