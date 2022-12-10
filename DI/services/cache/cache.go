package cache

import (
	"fmt"
	"sync"
)

// 将被注入的记录器服务的接口，用小写的方式来隐藏出口，是一个好的方式。
type logger interface {
	Error(error)
	Info(string)
}

type Cache struct {
	logger logger
	m      *sync.Map
}

// NewCache 我们服务的构造函数，接收将被注入的服务的接口（我们可以注入几个服务），并返回结构（缓存的实例）。
func NewCache(logger logger) *Cache {
	return &Cache{
		logger: logger,
		m:      new(sync.Map),
	}
}

func (r *Cache) Get(key string) (string, error) {
	rawValue, ok := r.m.Load(key)
	if !ok {
		err := fmt.Errorf("not found: %s", key)
		r.logger.Error(err)
		return "", err
	}
	return rawValue.(string), nil
}

func (r *Cache) Set(key string, data []byte) {
	r.m.Store(key, data)
	r.logger.Info(fmt.Sprintf("store key: %s", key))
}
