package main

import "fmt"

func makeGreeter(greeting string) func(string) string {
	return func(name string) string {
		fmt.Printf("func greeting: %s, name: %s\n", greeting, name)
		return greeting + ", " + name
	}
}

func func3() {
	englishGreeter := makeGreeter("Hello")
	spanishGreeter := makeGreeter("Hola")

	fmt.Println(englishGreeter("John"))
	fmt.Println(englishGreeter("Tim"))
	fmt.Println(spanishGreeter("Juan"))
	fmt.Println(spanishGreeter("Taylor"))
}
