package TGmodels

type Update struct {
	TelegramModel
	UpdateId      int     `json:"update_id"`
	Message       Message `json:"message"`
	EditedMessage Message `json:"edited_message"`
	// etc
}
