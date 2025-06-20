package models

type FiftyFiftyLink struct {
	ID          int     `db:"id" json:"id"`
	Probability float64 `db:"probability" json:"probability"`
	URLa        string  `db:"url_a" json:"urlA"`
	URLb        string  `db:"url_b" json:"urlB"`
	ShortCode   string  `db:"short_code" json:"shortCode"`
}
