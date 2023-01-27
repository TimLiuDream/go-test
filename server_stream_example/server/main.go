package main

import (
	"log"
	"math/rand"
	"net"

	pb "github.com/timliudream/go-test/server_stream_example/protofiles/data_streaming"
	"google.golang.org/grpc"
)

type server struct{}

// *DataRequest, StreamingService_GetDataStreamingServer) error
func (s server) GetDataStreaming(req *pb.DataRequest, srv pb.StreamingService_GetDataStreamingServer) error {
	log.Println("Fetch data streaming")

	for i := 0; i < 10; i++ {
		value := randStringBytes(500)

		resp := pb.DataResponse{
			Part:   int32(i),
			Buffer: value,
		}

		if err := srv.Send(&resp); err != nil {
			log.Println("error generating response")
			return err
		}
	}

	return nil
}

func randStringBytes(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func main() {
	// create listener
	listener, err := net.Listen("tcp", "localhost:10080")

	if err != nil {
		panic("error building server: " + err.Error())
	}

	// create gRPC server
	s := grpc.NewServer()
	pb.RegisterStreamingServiceServer(s, server{})

	log.Println("start server")

	if err := s.Serve(listener); err != nil {
		panic("error building server: " + err.Error())
	}

}
