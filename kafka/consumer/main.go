package main

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// kafka consumer by kafka-go

//
//func main() {
//	config := kafka.ReaderConfig{
//		Brokers: []string{"localhost:9092"},
//		GroupID: "my-group",
//		Topic:   "test",
//		// MinBytes: 10e3, // 10KB
//		// MaxBytes: 10e6, // 10MB
//		// MaxWait:  1 * time.Second,
//	}
//
//	r := kafka.NewReader(config)
//	defer r.Close()
//	offsetReq := kafka.TimeOffsetOf(0, time.Now().Add(-1*time.Hour))
//	r.SetOffsetAt()
//
//	signals := make(chan os.Signal, 1)
//	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
//
//	ctx, cancel := context.WithCancel(context.Background())
//	defer cancel()
//
//	go func() {
//		sig := <-signals
//		fmt.Printf("Caught signal %s: terminating\n", sig)
//		cancel()
//	}()
//
//	for {
//		select {
//		case <-ctx.Done():
//			fmt.Println("terminating: context cancelled")
//			return
//		default:
//			m, err := r.ReadMessage(ctx)
//			if err != nil {
//				log.Println(err)
//				break
//			}
//			log.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
//		}
//	}
//}

// Compare this snippet from kafka/consumer/main.go:
// package main
//
// // kafka consumer by sarama
//
// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"os"
// 	"os/signal"
// 	"syscall"
//
// 	"github.com/Shopify/sarama"
// )
//

func main() {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.AutoCommit.Enable = false
	config.Version = sarama.V2_0_0_0

	client, err := sarama.NewClient([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	consumer, err := sarama.NewConsumerFromClient(client)
	if err != nil {
		log.Fatalln(err)
	}
	defer consumer.Close()

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
	partitionList, err := consumer.Partitions(topic)
	if err != nil {
		log.Fatalln(err)
	}

	for partition := range partitionList {
		// 只能根据时间来获取 offset
		// how to use GetOffset
		//offset, err := client.GetOffset(topic, int32(partition), sarama.OffsetNewest)
		//if err != nil {
		//	panic(err)
		//}
		//log.Printf("offset: %d", offset)
		pc, err := consumer.ConsumePartition(topic, int32(partition), 20)
		if err != nil {
			log.Fatalln(err)
		}
		defer pc.AsyncClose()

		go func(sarama.PartitionConsumer) {
			for {
				select {
				case <-ctx.Done():
					fmt.Println("terminating: context cancelled")
					return
				case msg := <-pc.Messages():
					log.Printf("Partition:%d Offset:%d Key:%s Value:%s", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
				case err := <-pc.Errors():
					log.Printf("PartitionConsumer error: %s", err)
				}
			}
		}(pc)
	}

	<-ctx.Done()
}
