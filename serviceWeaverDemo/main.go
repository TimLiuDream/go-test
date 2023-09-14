package main

import (
	"context"
	"github.com/ServiceWeaver/weaver"
	"github.com/timliudream/go-test/serviceWeaverDemo/study"
	"log"
)

func main() {
	if err := weaver.Run(context.Background(), study.Serve); err != nil {
		log.Fatal(err)
	}
}
