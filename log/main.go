package main

import "log"

// "GolangTraining/LogLib"

// "github.com/sirupsen/logrus"

func init() {
	log.SetPrefix("TRACE:")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func main() {

	log.Println("message")

	log.Fatalln("fatal message")

	log.Panicln("panic message")
	// log.Info("HelloWorld!")
	// d := 24 * time.Hour
	// d1 := 12 * time.Hour
	// LogLib.ConfigLocalFilesystemLogger("./log/", "test.log", d, d1)

	// log.Info("Hello World!")
}
