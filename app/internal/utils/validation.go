package utils

import (
	"errors"
	"net/url"
	"regexp"
)

var ErrInvalidURL = errors.New("invalid URL format")

var hostPattern = regexp.MustCompile(`^[a-zA-Z0-9.-]+$`)

func ValidateAndFixURL(urlString string) (string, error) {
	// first, check if url starts wth http or https
	// if it doesn't, add http:// to the beginning
	// then, check if url is validly formatted in general
	// throw an error if it is not validly formatted

	// any url without at least 3 characters is invalid
	if len(urlString) < 3 {
		return "", ErrInvalidURL
	}

	if !(len(urlString) > 4 && urlString[:7] == "http://" || len(urlString) > 5 && urlString[:8] == "https://") {
		urlString = "http://" + urlString
	}

	parsed, err := url.Parse(urlString)

	if err != nil || parsed.Scheme == "" || parsed.Host == "" {
		return "", ErrInvalidURL
	}

	if parsed.Scheme != "http" && parsed.Scheme != "https" {
		return "", ErrInvalidURL
	}

	// Ensure host contains only allowed characters
	if !hostPattern.MatchString(parsed.Hostname()) {
		return "", ErrInvalidURL
	}

	return urlString, nil
}
