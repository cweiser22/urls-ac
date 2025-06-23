package repository

import (
	"fmt"

	"github.com/cweiser22/urls-ac/internal/models"
	"github.com/jmoiron/sqlx"
)

type FiftyFiftyLinkRepository struct {
	DB *sqlx.DB
}

// NewFiftyFiftyLinkRepository creates a new instance of FiftyFiftyLinkRepository.
func NewFiftyFiftyLinkRepository(db *sqlx.DB) *FiftyFiftyLinkRepository {
	return &FiftyFiftyLinkRepository{
		DB: db,
	}
}

// CreateFiftyFiftyLinkDTO is used to insert a new FiftyFiftyLink.
type CreateFiftyFiftyLinkDTO struct {
	Probability float64 `db:"probability_a"`
	URLa        string  `db:"url_a"`
	URLb        string  `db:"url_b"`
	ShortCode   string  `db:"short_code"`
}

// Insert adds a new FiftyFiftyLink and returns the complete saved model.
func (r *FiftyFiftyLinkRepository) Insert(dto *CreateFiftyFiftyLinkDTO) (*models.FiftyFiftyLink, error) {
	query := `
		INSERT INTO fifty_fifty_links (probability_a, url_a, url_b, short_code)
		VALUES (:probability_a, :url_a, :url_b, :short_code)
		RETURNING id
	`
	rows, err := r.DB.NamedQuery(query, dto)
	if err != nil {
		return nil, fmt.Errorf("insert link: %w", err)
	}
	defer rows.Close()

	var id int
	if rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return nil, fmt.Errorf("scan id: %w", err)
		}
	}

	return &models.FiftyFiftyLink{
		ID:          id,
		Probability: dto.Probability,
		URLa:        dto.URLa,
		URLb:        dto.URLb,
		ShortCode:   dto.ShortCode,
	}, nil
}

// Delete removes a link by its shortcode.
func (r *FiftyFiftyLinkRepository) Delete(shortCode string) error {
	query := `DELETE FROM fifty_fifty_links WHERE short_code = $1`
	_, err := r.DB.Exec(query, shortCode)
	if err != nil {
		return fmt.Errorf("delete link: %w", err)
	}
	return nil
}

// GetByShortCode retrieves a link by its shortcode.
func (r *FiftyFiftyLinkRepository) GetByShortCode(shortCode string) (*models.FiftyFiftyLink, error) {
	var link models.FiftyFiftyLink
	query := `SELECT id, probability_a, url_a, url_b, short_code FROM fifty_fifty_links WHERE short_code = $1`
	if err := r.DB.Get(&link, query, shortCode); err != nil {
		return nil, fmt.Errorf("get link by shortcode: %w", err)
	}
	return &link, nil
}
