package utils

import (
	"testing"
)

func TestValidateAndFixURL(t *testing.T) {
	type testCase struct {
		input    string
		expected string
		wantErr  bool
	}

	tests := []testCase{
		// Valid URLs with scheme
		{"http://example.com", "http://example.com", false},
		{"https://example.com", "https://example.com", false},
		{"https://www.google.com", "https://www.google.com", false},

		// Valid URLs without scheme (should default to http)
		{"example.com", "http://example.com", false},
		{"www.google.com", "http://www.google.com", false},
		{"test.org/path", "http://test.org/path", false},

		// Invalid URLs
		{"", "", true},
		{"ht!tp://bad", "", true},
		{"///", "", true},
		{"..", "", true},
		{"ab", "", true}, // < 3 characters
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			got, err := ValidateAndFixURL(tc.input)

			if tc.wantErr {
				if err == nil {
					t.Errorf("Expected error for input %q but got none", tc.input)
				}
			} else {
				if err != nil {
					t.Errorf("Did not expect error for input %q but got: %v", tc.input, err)
				}
				if got != tc.expected {
					t.Errorf("Expected %q but got %q", tc.expected, got)
				}
			}
		})
	}
}
