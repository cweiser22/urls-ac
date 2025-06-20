package service

import (
	"context"

	"errors"
	"fmt"
	"github.com/cweiser22/urls-ac/internal/models"
	"github.com/jmoiron/sqlx"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/redis/go-redis/v9"

	"log"

	"time"
)

type ShortenService struct {
	DB               *sqlx.DB
	RedisClient      *redis.Client
	CacheMetrics     *prometheus.CounterVec
	ShortCodeService *ShortCodeService
}

func NewShortCodeService(db *sqlx.DB, redisClient *redis.Client, cacheMetrics *prometheus.CounterVec) *ShortenService {
	return &ShortenService{
		DB:           db,
		RedisClient:  redisClient,
		CacheMetrics: cacheMetrics,
	}
}

// only used for testing to blindly insert data
func (s *ShortenService) insertShortcode(shortCode string, longURL string) error {
	// simply insert the shortcode with the data blindly
	_, err := s.DB.Exec("INSERT INTO url_mappings (long_url, short_code) VALUES ($1, $2)", longURL, shortCode)
	return err
}

// findMatchOrCollision checks if the generated short code matches an existing URLMapping or collides with another
// if a short code with the same URLMapping exists, it returns false and the existing URLMapping
// if a short code with a different URLMapping exists, it returns true to indicate a collision
// it returns an error if the database query fails
func (s *ShortenService) findMatchOrCollision(shortCode string, longURL string) (bool, *models.URLMapping, error) {
	var existing models.URLMapping

	err := s.DB.Get(&existing, `
		SELECT id, long_url, short_code, created_at
		FROM url_mappings
		WHERE short_code = $1
	`, shortCode)

	if err != nil {
		// No row means no match or collision
		if err.Error() == "sql: no rows in result set" {
			return false, nil, nil
		}
		return false, nil, err
	}

	if existing.LongURL == longURL {
		// Same long URLMapping: treat as match (not a collision)
		return false, &existing, nil
	}

	// Different long URLMapping: it's a collision
	return true, nil, nil
}

func (s *ShortenService) GetOrCreateMapping(longURL string) (*models.URLMapping, error) {
	// this function is the core logic that creates url mappings
	// it will iterate through numbers 6 to 15
	// for each number, it will attempt to generate a short code
	// then it will check for a collision or match
	// if there is a match, it will return the existing URLMapping
	// if there is a collision, it will continue to the next number
	// if there is no match or collision, it will insert the new mapping and return it
	validLengths := []int{6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	for _, length := range validLengths {
		shortCode := s.ShortCodeService.GenerateShortcode(longURL, length)

		collision, existingURL, err := s.findMatchOrCollision(shortCode, longURL)
		if err != nil {
			return nil, fmt.Errorf("error checking for match or collision: %w", err)
		}

		if collision {
			continue // Collision found, try next length
		}

		if existingURL != nil {
			return existingURL, nil // Match found, return existing URLMapping
		}

		// No match or collision, insert new mapping
		newURL := models.NewURLMapping(0, longURL, shortCode)
		if err := s.insertShortcode(newURL.ShortCode, newURL.LongURL); err != nil {
			return nil, fmt.Errorf("error inserting new shortcode: %w", err)
		}
		return newURL, nil // Return the newly created URLMapping mapping
	}
	// If we reach here, it means we couldn't find a valid shortcode
	return nil, fmt.Errorf("could not find a valid shortcode for %s (this is exceedingly rare)", longURL)
}

func (s *ShortenService) CacheMapping(mapping models.URLMapping) error {
	// cache key code:shortCode to value longURL in redis
	cacheKey := fmt.Sprintf("code:%s", mapping.ShortCode)
	err := s.RedisClient.Set(context.Background(), cacheKey, mapping.LongURL, time.Second*60).Err()
	if err != nil {
		return fmt.Errorf("error caching mapping: %w", err)
	}
	return nil // Return nil if caching was successful
}

func (s *ShortenService) GetCachedMapping(shortCode string) (string, error) {
	// retrieve the cached mapping from redis
	cacheKey := fmt.Sprintf("code:%s", shortCode)
	longURL, err := s.RedisClient.Get(context.Background(), cacheKey).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {

			return "", nil // No cached mapping found
		}
		return "", fmt.Errorf("error retrieving cached mapping: %w", err)
	}
	err = s.CacheMapping(models.URLMapping{
		LongURL:   longURL,
		ShortCode: shortCode,
	})
	if err != nil {
		return "", fmt.Errorf("error resetting ttl for cached mapping: %w", err)
	} // reset ttl
	return longURL, nil // Return the cached long URLMapping
}

// GetMapping retrieves a URLMapping mapping by its short code.
func (s *ShortenService) GetLongURL(shortCode string) (string, error) {

	cachedURL, err := s.GetCachedMapping(shortCode)
	if err != nil {
		log.Println("cache unavailable, falling back to database:", err)
	}

	if cachedURL != "" {
		s.CacheMetrics.With(prometheus.Labels{"result": "hit"}).Inc()
		return cachedURL, nil // Return cached URLMapping if available
	}
	s.CacheMetrics.With(prometheus.Labels{"result": "miss"}).Inc()

	var url models.URLMapping
	err = s.DB.Get(&url, `
		SELECT id, long_url, short_code, created_at
		FROM url_mappings
		WHERE short_code = $1
	`, shortCode)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return "", nil // No mapping found
		}
		return "", fmt.Errorf("error retrieving mapping: %w", err)
	}

	// Cache the found URLMapping mapping for future requests
	if err := s.CacheMapping(url); err != nil {
		return "", fmt.Errorf("error caching mapping: %w", err)
	}

	return url.LongURL, nil // Return the found URLMapping mapping
}
