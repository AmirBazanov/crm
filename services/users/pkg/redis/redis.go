package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log/slog"
)

type Client struct {
	client *redis.Client
	logger *slog.Logger
}

func NewClient(logger *slog.Logger, addr string, password string, db int, ctx context.Context) (*Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		logger.Error("redis ping error:", err)
		return nil, err
	}
	logger.Info("redis ping success")
	return &Client{rdb, logger}, nil
}

func (c *Client) Close() {
	err := c.client.Close()
	if err != nil {
		c.logger.Error("redis close error:", err)
		return
	}
}

func (c *Client) Set(key string, value interface{}, ctx context.Context) error {
	return c.client.Set(ctx, key, value, 0).Err()
}

func (c *Client) Get(key string, ctx context.Context) (string, error) {
	return c.client.Get(ctx, key).Result()
}

func (c *Client) Del(key string, ctx context.Context) error {
	return c.client.Del(ctx, key).Err()

}
