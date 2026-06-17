package validation

type UseCaseResult struct {
	HasError   bool       `json:"has_error"`
	HasWarning bool       `json:"has_warning"`
	Messages   []*Message `json:"messages"`
	Output     any        `json:"-"`
}

func NewUseCaseResult() *UseCaseResult {
	return &UseCaseResult{
		Messages: make([]*Message, 0),
	}
}

func NewFailUseCaseResult(code string, message string) *UseCaseResult {
	result := NewUseCaseResult()
	result.AddMessage(NewMessage(code, Error, message))
	return result
}

func NewWarninglUseCaseResult(code string, message string) *UseCaseResult {
	result := NewUseCaseResult()
	result.AddMessage(NewMessage(code, Warning, message))
	return result
}

func (vr *UseCaseResult) AddMessage(message *Message) {
	if message == nil {
		return
	}

	vr.Messages = append(vr.Messages, message)

	switch message.Type {
	case Error:
		vr.HasError = true
	case Warning:
		vr.HasWarning = true
	}
}

func (vr *UseCaseResult) AddError(code string, message string) {
	vr.AddMessage(NewMessage(code, Error, message))
}

func (vr *UseCaseResult) AddWarning(code string, message string) {
	vr.AddMessage(NewMessage(code, Warning, message))
}

func (vr *UseCaseResult) HasMessage() bool {
	return len(vr.Messages) > 0
}
