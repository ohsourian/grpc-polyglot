package main

import (
	"context"
	"flag"
	"fmt"
	pb "github.com/ohsourian/grpc-polyglot/go-grpc/config/grpc"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	port = flag.Int("port", 30043, "gRPC Server Port")
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Name: %v", in.GetName())
	return &pb.HelloReply{
		Message: "Hello, " + in.GetName(),
	}, nil
}
func main() {
	log.Println("Go Init")
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Failed to Start Server Port: %d, %v", *port, err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("Server At %v", lis.Addr())
	lisErr := s.Serve(lis)
	if lisErr != nil {
		log.Fatalf("Faild to Start Server: %v", err)
	}
}
