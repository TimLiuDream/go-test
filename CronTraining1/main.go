package main

import (
	"fmt"
	"github.com/robfig/cron"
	"log"
)

type TestJob struct {
}

func (this TestJob) Run() {
	fmt.Println("testJob1")
}

type Test2Job struct {
}

func (this Test2Job) Run() {
	fmt.Println("testJob2")
}

func main() {
	i := 0
	c := cron.New()

	spec := "*/5 * * * * ?"

	c.AddFunc(spec, func() {
		i++
		log.Println("cron running :", i)
	})
	c.AddJob(spec, TestJob{})
	c.AddJob(spec, Test2Job{})

	c.Start()

	defer c.Stop()

	select {

	}
}
