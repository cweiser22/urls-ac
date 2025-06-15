package models

import (
	"time"
)

type URLMapping struct {
	ID        int       `db:"id" json:"id"`
	LongURL   string    `db:"long_url" json:"longUrl"`
	ShortCode string    `db:"short_code" json:"shortCode"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
}

func NewURLMapping(ID int, longURL, shortCode string) *URLMapping {
	return &URLMapping{
		ID:        ID,
		LongURL:   longURL,
		ShortCode: shortCode,
	}
}

func (url *URLMapping) Equals(other *URLMapping) bool {
	return url.LongURL == other.LongURL
}
