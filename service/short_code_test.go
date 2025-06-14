package service

import (
	"github.com/cweiser22/urls-ac/testsupport"
	"testing"
)

func TestGenerateShortcodeConsistency(t *testing.T) {

	s := &ShortCodeService{}
	url := "example.com"

	// Check that the shortcode is consistent across 10 generations for lengths 5 to 9
	for length := 5; length <= 9; length++ {
		var code string
		for i := 0; i < 10; i++ {
			next := s.generateShortcode(url, length)
			if i == 0 {
				code = next
			} else if next != code {
				t.Errorf("Inconsistent shortcode for %s at length %d: got %s, expected %s", url, length, next, code)
			}
		}
	}
}

func TestGenerateShortcodeUniqueness(t *testing.T) {
	s := &ShortCodeService{}
	urls := []string{"youtube.com", "google.com"}

	for length := 5; length <= 9; length++ {
		code1 := s.generateShortcode(urls[0], length)
		code2 := s.generateShortcode(urls[1], length)
		if code1 == code2 {
			t.Errorf("Shortcodes should differ for %s and %s at length %d, but both got %s", urls[0], urls[1], length, code1)
		}
	}
}

func TestGenerateShortcodeAmazonDifferentLengths(t *testing.T) {
	s := &ShortCodeService{}
	url := "amazon.com"
	seen := make(map[string]bool)

	for length := 5; length <= 10; length++ {
		code := s.generateShortcode(url, length)
		if seen[code] {
			t.Errorf("Duplicate shortcode for amazon.com at length %d: %s", length, code)
		}
		seen[code] = true
	}
}

// create a test for findMatchOrCollision that uses the test database and covers all cases:
func TestFindMatchOrCollision(t *testing.T) {
	pg, err := testsupport.GetTestPostgres()
	if err != nil {
		t.Fatalf("failed to get postgres: %v", err)
	}

	s := &ShortCodeService{
		DB: pg.DB,
	}

	longURL := "http://example1.com"
	shortCode := s.generateShortcode(longURL, 6)

	// Insert the shortcode into the database
	if err := s.insertShortcode(shortCode, longURL); err != nil {
		t.Fatalf("failed to insert shortcode: %v", err)
	}

	// Test for a match
	collision, existingURL, err := s.findMatchOrCollision(shortCode, longURL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if collision {
		t.Errorf("should not have gotten collision")
	}
	if existingURL == nil {
		t.Fatal("expected existing URL to be non-nil")
	}

	// Test for a collision with a different URL
	collisionLongURL := "http://another-example1.com"
	collisionMatch, _, err := s.findMatchOrCollision(shortCode, collisionLongURL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !collisionMatch {
		t.Error("expected collision with different URL")
	}

	// Test for a completely new shortcode and url
	newShortCode := s.generateShortcode("http://new-example1.com", 6)
	collisionNew, existingNewURL, err := s.findMatchOrCollision(newShortCode, "http://new-example1.com")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if collisionNew {
		t.Errorf("should not have gotten collision for new shortcode %s", newShortCode)
	}
	if existingNewURL != nil {
		t.Errorf("expected existing URL to be nil for new shortcode %s, got %v", newShortCode, existingNewURL)
	}
}

// make a test for GetOrCreateMapping that uses the test database and covers all cases
func TestGetOrCreateMapping(t *testing.T) {
	pg, err := testsupport.GetTestPostgres()
	if err != nil {
		t.Fatalf("failed to get postgres: %v", err)
	}

	s := &ShortCodeService{
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

	// Try to get the same mapping again
	mapping2, err := s.GetOrCreateMapping(longURL)
	if err != nil {
		t.Fatalf("unexpected error on second call: %v", err)
	}

	if mapping2.ShortCode != mapping.ShortCode {
		t.Errorf("expected same shortcode on second call, got %s and %s", mapping.ShortCode, mapping2.ShortCode)
	}
}
