package db

// write a function to create a redis client with a given url
import (
	"context"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func NewRedisClient(redisURL string) (*redis.Client, error) {
	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		panic(err)
	}
	client := redis.NewClient(opt)

	// Test the connection
	_, err = client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}
