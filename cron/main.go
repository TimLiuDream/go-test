package main

import (
	"fmt"
	"github.com/revel/cron"
	"time"
)

func main() {
	c := cron.New()
	spec := "0 0 8 * * *"
	c.AddFunc(spec, func() {
		fmt.Println(time.Now())
	})
	c.Start()
	select {}
}
