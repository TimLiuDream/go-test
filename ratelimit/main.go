package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

func main() {
	//mux := http.NewServeMux()
	//mux.HandleFunc("/", okHandler)
	//
	//log.Println("Listening on :4000")
	//http.ListenAndServe(":4000", limit(mux))
	adaptiveRateLimiting()
}

func okHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func fixedWindowRateLimiting() {
	limiter := rate.NewLimiter(rate.Limit(100), 1) // 允许每秒100次

	for i := 0; i < 200; i++ {
		if !limiter.Allow() {
			fmt.Println("Rate limit exceeded. Request rejected.")
			continue
		}
		go process()
	}
}

func tokenBucketRateLimiting() {
	limiter := rate.NewLimiter(rate.Limit(10), 5)
	ctx, _ := context.WithTimeout(context.TODO(), time.Millisecond)
	for i := 0; i < 200; i++ {
		if err := limiter.Wait(ctx); err != nil {
			fmt.Println("Rate limit exceeded. Request rejected.")
			continue
		}
		go process()
	}
}

func dynamicRateLimiting() {
	limiter := rate.NewLimiter(rate.Limit(10), 1) // 允许每秒100次

	// Dynamic rate adjustment
	go func() {
		time.Sleep(time.Second * 10) // 每10秒调整 limiter
		fmt.Println("---adjust limiter---")
		limiter.SetLimit(rate.Limit(200)) // 将 limiter 提升到每秒 200
	}()

	for i := 0; i < 3000; i++ {
		if !limiter.Allow() {
			fmt.Println("Rate limit exceeded. Request rejected.")
			time.Sleep(time.Millisecond * 100)
			continue
		}
		process()
	}
}

func adaptiveRateLimiting() {
	limiter := rate.NewLimiter(rate.Limit(10), 1) // 允许每秒10次

	// 自适应调整
	go func() {
		for {
			time.Sleep(time.Second * 10)
			responseTime := measureResponseTime() // 测量之前请求的响应时间
			if responseTime > 500*time.Millisecond {
				fmt.Println("---adjust limiter 50---")
				limiter.SetLimit(rate.Limit(50))
			} else {
				fmt.Println("---adjust limiter 100---")
				limiter.SetLimit(rate.Limit(100))
			}
		}
	}()

	for i := 0; i < 3000; i++ {
		if !limiter.Allow() {
			fmt.Println("Rate limit exceeded. Request rejected.")
			time.Sleep(time.Millisecond * 100)
			continue
		}
		process()
	}
}

// 测量以前请求的响应时间
// 执行自己的逻辑来测量响应时间
func measureResponseTime() time.Duration {
	return time.Millisecond * 100
}

// 处理请求
func process() {
	fmt.Println("Request processed successfully.")
	time.Sleep(time.Millisecond * 10) // 模拟请求处理时间
}
