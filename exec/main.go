package main

import (
	"fmt"
	"log"
	"os/exec"
	"time"
)

const (
	originPath       = "/Users/tim/Downloads/普通测试文件.html"
	targetPath       = "/Users/tim/Downloads/test.docx"
	referenceDocPath = "/Users/tim/Downloads/custom-reference.docx"
)

const (
	TimeOutSecond = 55
)

func main() {
	cmd := exec.Command("pandoc", "-f", "HTML", originPath, "-o", targetPath, "--reference-doc", referenceDocPath)
	go listenPandoc(cmd, "234567")
	err := cmd.Run()
	if err != nil {
		if err.Error() == "signal: killed" {
			log.Println("over time!")
		}
		fmt.Println("here")
		log.Fatalln(err.Error())
	}
	log.Println("complete!")
}

func listenPandoc(cmd *exec.Cmd, pageUUID string) {
	<-time.After(time.Second * TimeOutSecond)
	if cmd.Process == nil {
		return
	} else {
		log.Printf("kill pandoc, pageUUID: %s\n", pageUUID)
		_ = cmd.Process.Kill()
		return
	}
}
