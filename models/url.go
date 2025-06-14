package models

import (
	"time"
)

type URL struct {
	ID        int       `db:"id" json:"id"`
	LongURL   string    `db:"long_url" json:"longUrl"`
	ShortCode string    `db:"short_code" json:"shortCode"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
}

func NewURL(ID int, longURL, shortCode string) *URL {
	return &URL{
		ID:        ID,
		LongURL:   longURL,
		ShortCode: shortCode,
	}
}

func (url *URL) Equals(other *URL) bool {
	return url.LongURL == other.LongURL
}
