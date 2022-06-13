package main

import (
	"context"
	"flag"
	pb "github.com/ohsourian/grpc-polyglot/go-grpc/config/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

const defaultName = "World"

var (
	addr = flag.String("addr", "localhost:30043", "gRPC server address")
	name = flag.String("name", defaultName, "guest")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Connection Fatal: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	log.Printf("Request Name: %s", *name)
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("Request Fail ::: %v", err)
	}
	log.Printf("Message: %s", r.GetMessage())
}
