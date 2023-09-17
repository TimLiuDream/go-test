package main

import "errors"

type cache interface {
	Set(key string, value interface{})
	Get(key string) interface{}
}

func DefaultCache() cache {
	return NewMemoryCache()
}

func NewCache(tp string) (cache, error) {
	switch tp {
	case "redis":
		return NewRedisCache(), nil
	default:
		return DefaultCache(), nil
	}
	return nil, errors.New("can not found target cache")
}
