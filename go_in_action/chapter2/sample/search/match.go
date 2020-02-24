package search

import (
	"fmt"
	"log"
)

// Result 保存搜索的数据
type Result struct {
	Field   string
	Content string
}

// Matcher 定义了要实现的新搜索类型的行为
type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

// Match 函数，为每个数据源单独启动goroutine来执行这个函数
// 并发地执行
func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result) {
	// 对特定的匹配器执行搜索
	searchResults, err := matcher.Search(feed, searchTerm)
	if err != nil {
		log.Println(err)
		return
	}

	// 将结果写进通道
	for _, searchResult := range searchResults {
		results <- searchResult
	}
}

// Display 从每个单独的goroutine接收到结果后
// 在终端窗口输出
func Display(results chan *Result) {
	// 通道会一直被阻塞，直到有结果写入
	// 一旦通道被关闭，for循环就会被终止
	for result := range results {
		fmt.Printf("%s:\n%s\n\n", result.Field, result.Content)
	}
}
