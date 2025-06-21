package repository

import (
	"github.com/cweiser22/urls-ac/internal/models"
	"github.com/cweiser22/urls-ac/internal/testsupport"
	"testing"
)

func TestURLMappingsRepository_InsertAndGetByShortCode(t *testing.T) {
	// set up test postgres
	pg, err := testsupport.GetTestPostgres()
	if err != nil {
		t.Fatalf("failed to get postgres: %v", err)
	}

	repo := &URLMappingsRepository{
		DB: pg.DB,
	}

	mapping := CreateURLMapping{
		LongURL:   "http://example.com",
		ShortCode: "exmpl",
	}

	createdMapping, err := repo.Insert(&mapping)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if createdMapping.Equals(&models.URLMapping{}) {
		t.Fatal("expected created mapping to not be nil")
	}
	if createdMapping.LongURL != mapping.LongURL || createdMapping.ShortCode != mapping.ShortCode {
		t.Errorf("expected created mapping to match input, got %v", createdMapping)
	}

	// Test inserting the same mapping again
	_, err = repo.Insert(&mapping)
	if err == nil {
		t.Error("expected error when inserting duplicate mapping, got nil")
	}

	// Test getting by short code
	retrievedMapping, err := repo.GetByShortCode(mapping.ShortCode)
	if err != nil {
		t.Fatalf("unexpected error when getting by short code: %v", err)
	}
	if retrievedMapping.Equals(&models.URLMapping{}) {
		t.Fatal("expected retrieved mapping to not be nil")
	}
	if retrievedMapping.LongURL != mapping.LongURL || retrievedMapping.ShortCode != mapping.ShortCode {
		t.Errorf("expected retrieved mapping to match input, got %v", retrievedMapping)
	}
}
