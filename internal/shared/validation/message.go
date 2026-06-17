package validation

type MessageType int

const (
	ValidationErrorCode = "VALIDATION_ERROR"
	UrlCodeNotFoundCode = "URL_CODE_NOT_FOUND"
)

const (
	Error MessageType = iota
	Warning
)

type Message struct {
	Code        string      `json:"code"`
	Type        MessageType `json:"type"`
	Description string      `json:"description"`
}

func NewMessage(code string, t MessageType, description string) *Message {
	return &Message{
		Code:        code,
		Type:        t,
		Description: description,
	}
}
