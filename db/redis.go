package db

// write a function to create a redis client with a given url
import (
	"context"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func NewRedisClient(redisURL string) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr: redisURL,
	})

	// Test the connection
	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}
