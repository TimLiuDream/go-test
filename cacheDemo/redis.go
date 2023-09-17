package main

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache() *RedisCache {
	c := &RedisCache{client: redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})}
	return c
}

func (c *RedisCache) Get(key string) interface{} {
	ctx := context.Background()
	return c.client.Get(ctx, key)
}

func (c *RedisCache) Set(key string, value interface{}) {
	ctx := context.Background()
	c.client.Set(ctx, key, value, -1)
}
