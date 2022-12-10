package main

import (
	"github.com/timliudream/go-test/DI/services/cache"
	"github.com/timliudream/go-test/DI/services/logger"
)

func main() {
	// 注册服务
	sentryService := new(logger.NormalSentryClient)
	loggerService := logger.NewLogger(sentryService)
	cacheService := cache.NewCache(loggerService)

	cacheService.Set("test", []byte("test"))
	cacheService.Get("test1")
}
