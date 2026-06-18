package models

type GetShortURLByCodeInput struct {
	Code string
}

type GetShortURLByCodeResponse struct {
	ID          string `json:"id"`
	Code        string `json:"code"`
	OriginalURL string `json:"original_url"`
	ShortURL    string `json:"short_url"`
	Clicks      int    `json:"clicks"`
	UsedAt      string `json:"used_at"`
	CreatedAt   string `json:"created_at"`
}
