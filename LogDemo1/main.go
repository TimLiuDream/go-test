package main

import "GolangTraining/LogLib"

func main() {
	LogLib.Trace.Println("i have something standard to say")
	LogLib.Info.Println("special information")
	LogLib.Warning.Println("There is something you need to know about")
	LogLib.Error.Println("something has failed")
}
