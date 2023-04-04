package main

// kafka producer

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Shopify/sarama"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	config.Producer.Timeout = 5 * time.Second

	p, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatalln(err)
	}
	defer p.Close()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		sig := <-signals
		fmt.Printf("Caught signal %s: terminating\n", sig)
		cancel()
	}()

	topic := "test"
	var i = 0
	for {
		select {
		case <-ctx.Done():
			fmt.Println("terminating: context cancelled")
			return
		default:
			msg := &sarama.ProducerMessage{
				Topic: topic,
				Value: sarama.StringEncoder(fmt.Sprintf("send msg, value: %d", i)),
			}
			partition, offset, err := p.SendMessage(msg)
			if err != nil {
				log.Printf("FAILED to send message: %s\n", err)
			} else {
				log.Printf("> message sent to partition %d at offset %d\n", partition, offset)
			}
			time.Sleep(time.Second)
			i++
		}
	}
}
