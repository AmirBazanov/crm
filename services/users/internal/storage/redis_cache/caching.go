package redis_cache

import (
	"context"
	"crm/services/users/pkg/redis"
	"encoding/json"
	"log/slog"
)

type Cache struct {
	redis  *redis.Client
	logger *slog.Logger
}

type Caching interface {
	Set(key string, value interface{}, ctx context.Context) error
	Get(key string, ctx context.Context) (string, error)
	Del(key string, ctx context.Context)
}

func New(logger *slog.Logger, redis *redis.Client) *Cache {
	return &Cache{logger: logger, redis: redis}
}

func (c *Cache) SaveToCache(ctx context.Context, key string, data interface{}) {
	op := "storage.RedisCache.SaveToCache"
	if bytes, err := json.Marshal(&data); err == nil {
		err := c.redis.Set(key, string(bytes), ctx)
		if err != nil {
			c.logger.Error("error caching", op, err)
		}
	}
	c.logger.Info("cache saved", op, key)
}

func (c *Cache) GetFromCache(ctx context.Context, key string) (data interface{}, error error) {
	op := "storage.RedisCache.GetFromCache"
	cached, err := c.redis.Get(key, ctx)
	if err == nil && cached != "" {
		if err := json.Unmarshal([]byte(cached), &data); err == nil {
			return data, nil
		}
	} else {
		c.logger.Info("no such key in cache", op, key)
	}
	return nil, err
}

func (c *Cache) DelFromCache(ctx context.Context, key string) (error error) {
	op := "storage.RedisCache.DelFromCache"
	err := c.redis.Del(key, ctx)
	if err != nil {
		c.logger.Error("error caching", op, err)
		return err
	}
	c.logger.Info("cached del", op, key)
	return nil
}
