package main

import (
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/miniprogram"
	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"
)

func getMiniProgram() *miniprogram.MiniProgram {
	wc := wechat.NewWechat()

	//redisOpts := &cache.RedisOpts{
	//	Host:        cfg.Redis.Host,
	//	Password:    cfg.Redis.Password,
	//	Database:    cfg.Redis.Database,
	//	MaxActive:   cfg.Redis.MaxActive,
	//	MaxIdle:     cfg.Redis.MaxIdle,
	//	IdleTimeout: cfg.Redis.IdleTimeout,
	//}
	//redisCache := cache.NewRedis(redisOpts)
	//wc.SetCache(redisCache)

	memory := cache.NewMemory()
	cfg := &miniConfig.Config{
		AppID:     appID,
		AppSecret: appSecret,
		Cache:     memory,
	}
	//cfg.Cache = redisCache
	return wc.GetMiniProgram(cfg)
}
