package main

import (
	"log"

	"github.com/robfig/cron"
)

func main() {
	i := 0
	c := cron.New()
	spec := "* * * * *"
	c.AddFunc(spec, func() {
		i++
		log.Println("execute per minute", i)
	})
	c.Start()
	select {}
}
