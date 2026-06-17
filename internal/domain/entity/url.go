package entity

import "time"

type URL struct {
	ID          string
	Code        string
	OriginalURL string
	ShortURL    string
	Clicks      int
	UsedAt      time.Time
	CreatedAt   time.Time
}
