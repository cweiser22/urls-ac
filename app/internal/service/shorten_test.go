package service

import (
	"github.com/cweiser22/urls-ac/internal/testsupport"
	"testing"
)

// make a test for GetOrCreateMapping that uses the test database and covers all cases
func TestShortenService_GetLongURL(t *testing.T) {
	pg, err := testsupport.GetTestPostgres()
	if err != nil {
		t.Fatalf("failed to get postgres: %v", err)
	}

	s := &ShortenService{
		DB: pg.DB,
	}

	longURL := "http://example2.com"
	result, err := s.GetLongURL(longURL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result != longURL {
		t.Errorf("expected %s, got %s", longURL, result)
	}
}
