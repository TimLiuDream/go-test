package main

import (
	"context"
	"fmt"
	"log"
	"os/exec"
	"time"
)

const (
	originPath       = "/Users/tim/go/src/github.com/timliudream/golangtraining/commandcontext/sample/图片 copy.html"
	targetPath       = "/Users/tim/go/src/github.com/timliudream/golangtraining/commandcontext/sample/test.docx"
	referenceDocPath = "/Users/tim/go/src/github.com/timliudream/golangtraining/commandcontext/sample/custom-reference.docx"
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
	fmt.Println(1111)
	cmd := exec.Command("pandoc", "-f", "HTML", originPath, "-o", targetPath, "--reference-doc", referenceDocPath)
	fmt.Println(2222)
	err := cmd.Run()
	fmt.Println(333)
	if err != nil {
		log.Fatalln(err)
	}
	end := time.Now()
	fmt.Printf("spend time : %+v\n", end.Sub(start))
}
