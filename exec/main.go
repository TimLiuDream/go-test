package main

import (
	"log"
	"os/exec"
	"time"
)

const (
	originPath       = "/Users/tim/Downloads/test的副本.html"
	targetPath       = "/Users/tim/Downloads/test.docx"
	referenceDocPath = "/Users/tim/Downloads/custom-reference.docx"
)

func main() {
	cmd := exec.Command("pandoc", "-f", "HTML", originPath, "-o", targetPath, "--reference-doc", referenceDocPath)
	go getPid(cmd)
	err := cmd.Run()
	if err != nil {
		if err.Error() == "signal: killed" {
			log.Println("over time!")
		}
		log.Fatalln(err.Error())
	}
	log.Println("complete!")
}

func getPid(cmd *exec.Cmd) {
	for cmd.Process == nil {
		continue
	}
	i := 1
	for cmd.Process != nil {
		if i == 60 {
			_ = cmd.Process.Kill()
		}
		log.Println(cmd.Process.Pid)
		time.Sleep(1 * time.Second)
		i++
		if cmd.ProcessState == nil {
			continue
		}
		log.Println(cmd.ProcessState.Success())
		if cmd.ProcessState.Success() {
			return
		}
	}
}
