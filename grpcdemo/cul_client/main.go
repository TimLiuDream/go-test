package main

import (
	"context"
	"log"
	"os"
	"strconv"

	pb "github.com/timliudream/go-test/grpcdemo/cul"
	"google.golang.org/grpc"
)

const (
	addr = "localhost:50001"
)

func main() {
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	c := pb.NewCulClient(conn)

	var A int64
	var B int64
	if len(os.Args) > 2 {
		A, _ = strconv.ParseInt(os.Args[1], 10, 64)
		B, _ = strconv.ParseInt(os.Args[2], 10, 64)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	r, err := c.CulInt(ctx, &pb.CulRequest{A: A, B: B})
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("the result is : %d", r.GetResult())
}
