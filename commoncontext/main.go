package main

import (
	"context"
	"fmt"
	"log"
	"os/exec"
	"time"
)

const (
	originPath       = "/Users/tim/go/src/github.com/timliudream/golangtraining/commoncontext/sample/普通测试文件.html"
	targetPath       = "/Users/tim/go/src/github.com/timliudream/golangtraining/commoncontext/sample/test.docx"
	referenceDocPath = "/Users/tim/go/src/github.com/timliudream/golangtraining/commoncontext/sample/custom-reference.docx"
)

const (
	Count         = 3
	TimeOutSecond = 55
)

var intChan chan int = make(chan int, Count)

func main() {
	// for i := 0; i < Count; i++ {
	// 	go func() {
	// 		func1()
	// 		intChan <- 1
	// 	}()
	// }
	// for i := 0; i < Count; i++ {
	// 	<-intChan
	// }

	func1()

	// func2()
}

func func1() {
	start := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), TimeOutSecond*time.Second)
	defer cancel()

	if err := exec.CommandContext(ctx, "pandoc", "-f", "HTML", originPath, "-o", targetPath, "--reference-doc", referenceDocPath).Run(); err != nil {
		if err.Error() == "signal: killed" {
			log.Println(1111111)
		} else {
			log.Println(222222222)
		}
		log.Println(err)
	}
	end := time.Now()
	fmt.Printf("spend time : %+v\n", end.Sub(start))
}

func func2() {
	start := time.Now()
	cmd := exec.Command("pandoc", "-f", "HTML", originPath, "-o", targetPath, "--reference-doc", referenceDocPath)
	err := cmd.Run()
	if err != nil {
		log.Fatalln(err)
	}
	end := time.Now()
	fmt.Printf("spend time : %+v\n", end.Sub(start))
}
