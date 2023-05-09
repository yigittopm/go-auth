package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type Client struct {
	client *redis.Client
	ctx    context.Context
}

func New() *Client {
	redis := redis.NewClient(&redis.Options{
		Addr:     "cache:6379",
		Password: "",
		DB:       0,
	})

	ctx := context.Background()

	return &Client{
		client: redis,
		ctx:    ctx,
	}
}

func (c *Client) Set(key string, value any, time time.Duration) error {
	err := c.client.Set(c.ctx, key, value, time).Err()
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) Get(key string) ([]byte, error) {
	value, err := c.client.Get(c.ctx, key).Result()
	if err != nil {
		return nil, err
	}
	return []byte(value), nil
}
