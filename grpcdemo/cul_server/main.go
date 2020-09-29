package main

import (
	"context"
	"log"
	"net"

	pb "github.com/timliudream/go-test/grpcdemo/cul"
	"google.golang.org/grpc"
)

const (
	port = ":50001"
)

type Server struct {
	pb.UnimplementedCulServer
}

func (s *Server) CulInt(ctx context.Context, in *pb.CulRequest) (*pb.CulReply, error) {
	log.Printf("received: A: %v, B: %v", in.GetA(), in.GetB())
	result := in.GetA() + in.GetB()
	return &pb.CulReply{Result: result}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln(err)
	}
	s := grpc.NewServer()
	log.Println("start server")
	pb.RegisterCulServer(s, &Server{})
	if err = s.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}
