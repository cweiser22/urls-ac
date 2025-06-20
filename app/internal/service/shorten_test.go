package service

import (
	"github.com/cweiser22/urls-ac/internal/testsupport"
	"testing"
)

// make a test for GetOrCreateMapping that uses the test database and covers all cases
func TestGetOrCreateMapping(t *testing.T) {
	pg, err := testsupport.GetTestPostgres()
	if err != nil {
		t.Fatalf("failed to get postgres: %v", err)
	}

	s := &ShortenService{
		DB: pg.DB,
	}

	longURL := "http://example2.com"
	mapping, err := s.GetOrCreateMapping(longURL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if mapping == nil || mapping.LongURL != longURL {
		t.Errorf("expected mapping for %s, got %v", longURL, mapping)
	}

}
