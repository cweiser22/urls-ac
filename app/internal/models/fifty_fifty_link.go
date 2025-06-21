package models

type FiftyFiftyLink struct {
	ID int `db:"id" json:"id"`
	// Chance of A being chosen (inverse is chance of B being chosen)
	Probability float64 `db:"probability_a" json:"probability"`
	URLa        string  `db:"url_a" json:"urlA"`
	URLb        string  `db:"url_b" json:"urlB"`
	ShortCode   string  `db:"short_code" json:"shortCode"`
}
