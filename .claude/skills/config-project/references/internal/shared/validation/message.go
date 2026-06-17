package validation

type MessageType int

const (
	ValidationError = "VALIDATION_ERROR"
	UrlCodeNotFound = "URL_CODE_NOT_FOUND"
)

const (
	Error   MessageType = iota // 0
	Warning                    // 1
)

type Message struct {
	Code        string
	Type        MessageType
	Description string
}

func NewMessage(code string, t MessageType, description string) *Message {
	return &Message{
		Code:        code,
		Type:        t,
		Description: description,
	}
}
