package logger

import "fmt"

type sentryClient interface {
	sendMessage(interface{})
}

type NormalSentryClient struct{}

func (s *NormalSentryClient) sendMessage(i interface{}) {
	fmt.Printf("send message %+v to sentry\n", i)
}

type Logger struct {
	sentry sentryClient
}

func NewLogger(sentryClient sentryClient) *Logger {
	return &Logger{
		sentry: sentryClient,
	}
}

// 我们看到这里有四个方法，但是 cache 对象只能看到两个：Error 以及 Info
func (l *Logger) Error(e error) {
	fmt.Printf("[ERROR] %+v\n", e)
}

func (l *Logger) Info(msg string) {
	fmt.Printf("[INFO] %s\n", msg)
}

func (l *Logger) Debug(msg string) {
	fmt.Printf("[DEBUG] %s\n", msg)
}

func (l *Logger) Log(msg string) {
	fmt.Printf("[LOG] %s\n", msg)
}
