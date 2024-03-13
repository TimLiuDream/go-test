package main

import (
	"fmt"
	"sync"
	"time"
)

func func4() {
	var wg sync.WaitGroup

	// 设置并发任务数量
	concurrency := 3

	// 创建一个用于控制并发的通道
	semaphore := make(chan struct{}, concurrency)

	// 假设有一组任务需要并发执行
	tasks := []string{"task1", "task2", "task3", "task4", "task5"}

	// 遍历任务列表
	for _, task := range tasks {
		// 增加 WaitGroup 的计数器
		wg.Add(1)

		// 启动一个 goroutine 来执行任务
		go func(t string) {
			// 在 goroutine 开始前向通道发送一个信号
			semaphore <- struct{}{}

			// 执行任务
			fmt.Println("Executing", t)
			time.Sleep(time.Second)

			// 模拟任务执行时间
			// 这里可以是任何实际的任务逻辑
			// ...

			// 任务完成后从通道释放一个信号
			<-semaphore

			// 减少 WaitGroup 的计数器
			wg.Done()
		}(task)
	}

	// 等待所有任务完成
	wg.Wait()

	fmt.Println("All tasks completed")
}
