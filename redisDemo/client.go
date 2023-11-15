package main

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	ctx context.Context
	*redis.Client
}

func getClient() *RedisClient {
	return &RedisClient{
		ctx: context.Background(),
		Client: redis.NewClient(&redis.Options{
			Network:            "",
			Addr:               "127.0.0.1:6379",
			Dialer:             nil,
			OnConnect:          nil,
			Username:           "",
			Password:           "",
			DB:                 0,
			MaxRetries:         0,
			MinRetryBackoff:    0,
			MaxRetryBackoff:    0,
			DialTimeout:        0,
			ReadTimeout:        0,
			WriteTimeout:       0,
			PoolFIFO:           false,
			PoolSize:           0,
			MinIdleConns:       0,
			MaxConnAge:         0,
			PoolTimeout:        0,
			IdleTimeout:        0,
			IdleCheckFrequency: 0,
			TLSConfig:          nil,
			Limiter:            nil,
		}),
	}
}
