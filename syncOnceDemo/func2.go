package main

import (
	"fmt"
	"sync"
	"time"
)

type Config struct {
	Server string
	Port   int64
}

var (
	once2  sync.Once
	config *Config
)

func ReadConfig() *Config {
	once2.Do(func() {
		config = &Config{
			Server: "test",
			Port:   10009,
		}
		fmt.Println("init config....")
	})
	return config
}

func func2() {
	for i := 0; i < 10; i++ {
		go func() {
			_ = ReadConfig()
		}()
	}
	time.Sleep(2 * time.Second)
}
