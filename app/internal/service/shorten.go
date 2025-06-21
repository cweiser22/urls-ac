package service

import (
	"github.com/cweiser22/urls-ac/internal/cache"
	"github.com/cweiser22/urls-ac/internal/models"
	"github.com/cweiser22/urls-ac/internal/repository"
	"log"

	"fmt"
	"github.com/prometheus/client_golang/prometheus"
)

type ShortenService struct {
	CacheMetrics         *prometheus.CounterVec
	ShortCodeService     *ShortCodeService
	URLMappingCache      *cache.URLMappingCache
	urlMappingRepository *repository.URLMappingsRepository
}

func NewShortenService(urlMappingCache *cache.URLMappingCache, cacheMetrics *prometheus.CounterVec, urlMappingRepository *repository.URLMappingsRepository) *ShortenService {
	return &ShortenService{
		CacheMetrics:         cacheMetrics,
		URLMappingCache:      urlMappingCache,
		urlMappingRepository: urlMappingRepository,
	}
}

// GetMapping retrieves a URLMapping mapping by its short code.
func (s *ShortenService) GetLongURL(shortCode string) (string, error) {
	cachedURL, err := s.URLMappingCache.GetCachedMapping(shortCode)
	if err != nil {
		log.Printf("Failed to get cached URL mapping: %v", err)
	}

	// handle a cache hit
	if cachedURL != "" {
		s.CacheMetrics.With(prometheus.Labels{"result": "hit"}).Inc()
		return cachedURL, nil // Return cached URLMapping if available
	}

	// handle a cache miss (database lookup)
	s.CacheMetrics.With(prometheus.Labels{"result": "miss"}).Inc()

	urlMapping, err := s.urlMappingRepository.GetByShortCode(shortCode)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return "", nil // No mapping found
		} else {
			return "", fmt.Errorf("error retrieving mapping: %w", err)
		}
	}

	// cache the found URLMapping mapping for future requests
	if err := s.URLMappingCache.CacheMapping(urlMapping); err != nil {
		return "", fmt.Errorf("error caching mapping: %w", err)
	}

	return urlMapping.LongURL, nil // Return the found URLMapping mapping
}

func (s *ShortenService) insertShortCode(longURL string) (models.URLMapping, error) {
	// Generate a short code for the long URL
	shortCode := s.ShortCodeService.GenerateShortcode(longURL, 6)

	// Create a new URLMapping mapping
	mapping := &repository.CreateURLMapping{
		LongURL:   longURL,
		ShortCode: shortCode,
	}

	// Insert the mapping into the database
	createdMapping, err := s.urlMappingRepository.Insert(mapping)
	if err != nil {
		return models.URLMapping{}, fmt.Errorf("error inserting URL mapping: %w", err)
	}

	return createdMapping, nil // Return the created short code
}

func (s *ShortenService) CreateURLMapping(longURL string) (models.URLMapping, error) {
	// allow up to 100 retries to avoid collisions
	// should rarely take more than 1-2 attempts
	for i := 0; i < 100; i++ {
		// try to generate a short code and insert it into the database
		shortCode, err := s.insertShortCode(longURL)
		if err != nil {
			log.Printf("Error inserting URL mapping: %v", err)
			continue // retry if there was an error
		}

		return shortCode, nil // return the created short code if successful
	}
	log.Println("Failed to create URL mapping after 100 attempts")
	return models.URLMapping{}, fmt.Errorf("failed to create URL mapping")
}
