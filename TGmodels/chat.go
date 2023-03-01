package TGmodels

type ChatType string

const (
	ChatTypePrivate    = "private"
	ChatTypeGroup      = "group"
	ChatTypeSupergroup = "supergroup"
	ChatTypeChannel    = "channel"
)

type Chat struct {
	TelegramModel
	Id        int      `json:"id"`
	Type      ChatType `json:"type"`
	Title     string   `json:"title,omitempty"`
	Username  string   `json:"username,omitempty"`
	FirstName string   `json:"first_name,omitempty"`
	LastName  string   `json:"last_name,omitempty"`
	// etc
}
