package main

import (
	"fmt"
	"log/syslog"
)

func main() {
	//对于syslog包，在windows上怎么编译都是会报undefined，在Linux上就不会，慎重使用
	g_log, err := syslog.NewLogger(syslog.LOG_INFO, 0)
	fmt.Print(g_log)
	fmt.Print(err)
}
