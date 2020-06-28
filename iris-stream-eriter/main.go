package main

import (
	"bufio"
	"fmt"  // just an optional helper
	"time" // showcase the delay

	"github.com/iris-contrib/graceful"
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()

	app.Get("/", func(ctx *iris.Context) {
		ctx.SetContentType("text/html")
		// ctx.SetHeader("Transfer-Encoding", "chunked")
		// i := 0
		ints := []int{1, 2, 3, 5, 7, 9, 11, 13, 15, 17, 23, 29}
		// Send the response in chunks and wait for half a second between each chunk.
		ctx.StreamWriter(func(w *bufio.Writer) {
			for _, i := range ints {
				s := fmt.Sprintf("Message number %d<br>", i)
				w.WriteString(s)
				w.Flush()
				// fmt.Fprintf(w, "Message number %d<br>", i)
				time.Sleep(500 * time.Millisecond) // simulate delay.
			}
		})
	})

	type messageNumber struct {
		Number int `json:"number"`
	}

	app.Get("/alternative", func(ctx *iris.Context) {
		ctx.SetContentType("text/html")
		ctx.SetHeader("Transfer-Encoding", "chunked")
		ints := []int{1, 2, 3, 5, 7, 9, 11, 13, 15, 17, 23, 29}
		// Send the response in chunks and wait for half a second between each chunk.
		ctx.StreamWriter(func(w *bufio.Writer) {
			for _, i := range ints {
				s := fmt.Sprintf("Message number %v<br>", messageNumber{Number: i})
				w.WriteString(s)
				w.Flush()
				time.Sleep(500 * time.Millisecond) // simulate delay
			}
		})
	})

	graceful.Run(":8080", time.Duration(1000)*time.Millisecond, app, func() {
		fmt.Println("all queues stop!")
	})
}
