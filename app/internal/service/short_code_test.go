package service

import "testing"

// NOTE: test has an extremely tiny chance of failing due to a collision
func TestGenerateShortcodeUniqueness(t *testing.T) {
	s := &ShortCodeService{}
	urls := []string{"youtube.com", "google.com"}

	for length := 5; length <= 9; length++ {
		code1 := s.GenerateShortcode(urls[0], length)
		code2 := s.GenerateShortcode(urls[1], length)
		if code1 == code2 {
			t.Errorf("shortcodes should differ for %s and %s at length %d, but both got %s, possible collision (extremely rare).", urls[0], urls[1], length, code1)
		}
	}
}
