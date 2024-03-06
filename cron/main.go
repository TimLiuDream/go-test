package main

import (
	"fmt"
	"github.com/revel/cron"
	"time"
)

func main() {
	c := cron.New()
	spec := "0 */10 * * * *"
	c.AddFunc(spec, func() {
		fmt.Println(time.Now())
	})

	spec := "0 */30 * * * *"
	c.Start()
	select {}
}
