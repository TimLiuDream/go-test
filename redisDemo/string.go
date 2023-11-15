package main

import (
	"github.com/go-redis/redis/v8"
	"log"
)

func (c *RedisClient) Set(key string, value interface{}) error {
	_, err := c.Client.Set(c.ctx, key, value, -1).Result()
	if err != nil {
		return err
	}
	return nil
}

func (c *RedisClient) Get(key string) (value interface{}, err error) {
	val, err := c.Client.Get(c.ctx, key).Result()
	if err != nil && err != redis.Nil {
		return nil, err
	} else if err == redis.Nil {
		log.Printf("key: %s is not found.\n", key)
		return "", nil
	}
	return val, nil
}
