package validation

type UseCaseResult struct {
	HasError   bool       `json:"has_error"`
	HasWarning bool       `json:"has_warning"`
	Messages   []*Message `json:"messages"`
}

func NewUseCaseResult() *UseCaseResult {
	return &UseCaseResult{
		Messages:   make([]*Message, 0),
		HasError:   false,
		HasWarning: false,
	}
}

func NewFailUseCaseResult(code string, message string) *UseCaseResult {
	return &UseCaseResult{
		Messages: append(make([]*Message, 0), NewMessage(code, Error, message)),
		HasError: true,
	}
}

func NewWarninglUseCaseResult(code string, message string) *UseCaseResult {
	return &UseCaseResult{
		Messages:   append(make([]*Message, 0), NewMessage(code, Warning, message)),
		HasWarning: true,
	}
}

func (vr *UseCaseResult) AddMessage(message *Message) {
	vr.Messages = append(vr.Messages, message)

	switch message.Type {
	case Error:
		vr.HasError = true
	case Warning:
		vr.HasWarning = true
	}
}

func (vr *UseCaseResult) AddError(code string, message string) {
	vr.Messages = append(vr.Messages, NewMessage(code, Error, message))
}

func (vr *UseCaseResult) AddWarning(code string, message string) {
	vr.Messages = append(vr.Messages, NewMessage(code, Warning, message))
}

func (vr *UseCaseResult) HasMessage() bool {

	return len(vr.Messages) > 0
}
