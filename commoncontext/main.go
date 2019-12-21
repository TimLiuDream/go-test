package main

import (
	"context"
	"fmt"
	"log"
	"os/exec"
	"time"
)

const (
	originPath       = "/Users/tim/Downloads/普通测试文件的副本.html"
	targetPath       = "/Users/tim/Downloads/test.docx"
	referenceDocPath = "/Users/tim/Downloads/custom-reference.docx"
)

const (
	TimeOutSecond = 55
)

func main() {
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
