package main

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	openai "github.com/sashabaranov/go-openai"
)

const (
	ApiKey        = "aster"
	ProxyHost     = "http://127.0.0.1:10809"
	systemContent = `
你是一位非常资深的 golang 架构师。非常精通以下能力：
- 通用架构
- golang 语言
- 算法
- 设计模式
- 计算机网络
- linux
- mysql
- redis
- 项目管理
但不局限于以上能力。
我会提供一些问题的具体信息，而你的任务就是基于你的认知为我解答。`
)

var client *openai.Client

func init() {
	proxyUri, err := url.Parse(ProxyHost)
	if err != nil {
		return
	}
	hClient := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyUri),
		},
		Timeout: 60 * time.Second,
	}
	defaultConfig := openai.DefaultConfig(ApiKey)
	defaultConfig.HTTPClient = hClient
	client = openai.NewClientWithConfig(defaultConfig)
}

func main() {
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: systemContent,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Hello! Gopher! 让我们一起使用 ChatGPT 吧！",
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	fmt.Println(resp.Choices[0].Message.Content)
}
