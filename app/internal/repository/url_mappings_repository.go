package repository

import (
	"github.com/cweiser22/urls-ac/internal/models"
	"github.com/jmoiron/sqlx"
	"log"
)

type URLMappingsRepository struct {
	DB *sqlx.DB
}

// NewURLMappingsRepository creates a new instance of URLMappingsRepository
func NewURLMappingsRepository(db *sqlx.DB) *URLMappingsRepository {
	return &URLMappingsRepository{
		DB: db,
	}
}

type CreateURLMapping struct {
	LongURL   string `db:"long_url"`
	ShortCode string `db:"short_code"`
}

func (r *URLMappingsRepository) Insert(mapping *CreateURLMapping) (models.URLMapping, error) {
	var createdMapping models.URLMapping
	query := `INSERT INTO url_mappings (long_url, short_code) VALUES (:long_url, :short_code) RETURNING id, long_url, short_code, created_at`
	rows, err := r.DB.NamedQuery(query, mapping)
	if err != nil {
		return models.URLMapping{}, err
	}
	if rows.Next() {
		err = rows.StructScan(&createdMapping)
		if err != nil {
			log.Fatal(err)
			return models.URLMapping{}, err
		}
	}
	return createdMapping, nil
}

func (r *URLMappingsRepository) GetByShortCode(shortCode string) (models.URLMapping, error) {
	var mapping models.URLMapping
	query := `SELECT id, long_url, short_code, created_at FROM url_mappings WHERE short_code = $1 LIMIT 1`
	err := r.DB.Get(&mapping, query, shortCode)
	if err != nil {
		return models.URLMapping{}, err
	}
	return mapping, nil
}
