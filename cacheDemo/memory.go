package main

import "sync"

type InMemoryCache struct {
	m     sync.Mutex
	store map[string]interface{}
}

func NewMemoryCache() *InMemoryCache {
	return &InMemoryCache{
		m:     sync.Mutex{},
		store: make(map[string]interface{}),
	}
}

func (c *InMemoryCache) Get(key string) interface{} {
	return c.store[key]
}

func (c *InMemoryCache) Set(key string, value interface{}) {
	c.m.Lock()
	defer c.m.Unlock()
	c.store[key] = value
}
