package cache

import (
	"context"
	"errors"
	"fmt"
	"github.com/cweiser22/urls-ac/internal/models"
	"github.com/redis/go-redis/v9"
	"time"
)

type URLMappingCache struct {
	RedisClient *redis.Client
}

func NewURLMappingCache(rdb *redis.Client) *URLMappingCache {
	return &URLMappingCache{
		RedisClient: rdb,
	}
}

func (s *URLMappingCache) GetCachedMapping(shortCode string) (string, error) {
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

func (s *URLMappingCache) CacheMapping(mapping models.URLMapping) error {
	// cache key code:shortCode to value longURL in redis
	cacheKey := fmt.Sprintf("code:%s", mapping.ShortCode)
	err := s.RedisClient.Set(context.Background(), cacheKey, mapping.LongURL, time.Second*60).Err()
	if err != nil {
		return fmt.Errorf("error caching mapping: %w", err)
	}
	return nil // Return nil if caching was successful
}
