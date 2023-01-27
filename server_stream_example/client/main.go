package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"google.golang.org/grpc/credentials/insecure"

	"google.golang.org/grpc"

	pb "github.com/timliudream/go-test/server_stream_example/protofiles/data_streaming"
)

func main() {
	// dial to server
	conn, err := grpc.Dial("localhost:10080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("Error connecting to gRPC server: ", err.Error())
	}

	defer conn.Close()

	// create the stream
	client := pb.NewStreamingServiceClient(conn)

	req := pb.DataRequest{Id: "123"}
	stream, err := client.GetDataStreaming(context.Background(), &req)
	if err != nil {
		panic(err) // dont use panic in your real project
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			return
		} else if err == nil {
			valStr := fmt.Sprintf("Response\n Part: %d \n Val: %s", resp.Part, resp.Buffer)
			log.Println(valStr)
		}

		if err != nil {
			panic(err) // dont use panic in your real project
		}

	}
}
