package main

import (
	"log"
	"os/exec"
	"time"
)

const (
	originPath       = "/Users/tim/Documents/KAk8aGVf_8DtmcCcR_1575012771030810.html"
	targetPath       = "/Users/tim/Downloads/test.docx"
	referenceDocPath = "/Users/tim/Downloads/custom-reference.docx"
)

func main() {
	cmd := exec.Command("pandoc", "-f", "HTML", originPath, "-o", targetPath, "--reference-doc", referenceDocPath)
	go listenPandoc(cmd)
	err := cmd.Run()
	if err != nil {
		if err.Error() == "signal: killed" {
			log.Println("over time!")
		}
		log.Fatalln(err.Error())
	}
	log.Println("complete!")
}

func listenPandoc(cmd *exec.Cmd) {
	for cmd.Process == nil {
		continue
	}
	i := 1
	for cmd.Process != nil {
		if i == 55 {
			log.Println("kill pandoc?")
			err := cmd.Process.Kill()
			if err != nil {
				log.Println(err.Error())
			}
			return
		}
		time.Sleep(1 * time.Second)
		i++
		if cmd.ProcessState == nil {
			continue
		}
		if cmd.ProcessState.String() == "signal: killed" {
			return
		}
		log.Println("Exited:", cmd.ProcessState.Exited())
		if cmd.ProcessState.Exited() {
			return
		}
		log.Println("cmd.ProcessState.Success():", cmd.ProcessState.Success())
		if cmd.ProcessState.Success() {
			return
		}
	}
}
