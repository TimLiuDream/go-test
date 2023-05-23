package main

import (
	"context"
	"fmt"
)

func main() {
	result, err := Call("test")
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func Call(prompt string) (string, error) {
	openAIKey := ""
	client := NewClient(openAIKey)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		ChatCompletionRequest{
			Model: GPT3Dot5Turbo,
			Messages: []ChatCompletionMessage{
				{
					Role:    ChatMessageRoleUser,
					Content: prompt + "将上述内容按照会议主题、会议时间、参与人、会议内容、待办事项五个点归纳为会议纪要",
				},
			},
		},
	)
	err = parseError(err)
	if err != nil {
		return "", err
	}
	return resp.Choices[0].Message.Content, nil
}
