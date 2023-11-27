package main

import (
	"github.com/timliudream/go-test/wireDemo/injector"
	"github.com/timliudream/go-test/wireDemo/provider"
)

func main() {
	// injector
	e, err := injector.InitializeEvent("Hello World!")
	if err != nil {
		panic(err)
	}
	e.Start()

	// provider
	event, err := provider.InitializeEvent("Provider!")
	if err != nil {
		panic(err)
	}
	event.Start()
}
