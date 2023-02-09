package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type Config struct {
	Message string
}

var conf = &Config{Message: "Before hot reload"}

func router() {
	log.Println("starting up....")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(conf.Message))
	})

	go func() {
		log.Fatal(http.ListenAndServe(":9002", nil))
	}()
}

func main() {
	router()
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)

	for {
		multiSignalHandler(<-sigCh)
	}
}

func multiSignalHandler(signal os.Signal) {
	switch signal {
	case syscall.SIGHUP:
		log.Println("Signal:", signal.String())
		log.Println("After hot reload")
		conf.Message = "Hot reload has been finished."
	case syscall.SIGINT:
		log.Println("Signal:", signal.String())
		log.Println("Interrupt by Ctrl+C")
		os.Exit(0)
	case syscall.SIGTERM:
		log.Println("Signal:", signal.String())
		log.Println("Process is killed.")
		os.Exit(0)
	default:
		log.Println("Unhandled/unknown signal")
	}
}
