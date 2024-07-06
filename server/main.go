package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	proto "github.com/maivankien/go-grpc-example/grpc/order"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 8000, "The server port")
)

type server struct {
	proto.UnimplementedOrderServiceServer
}

func (s *server) NewOrder(ctx context.Context, in *proto.OrderRequest) (*proto.OrderResponse, error) {
	log.Printf("Received order request for %s", in.GetOrderRequest())
	return &proto.OrderResponse{OrderResponse: "New order " + in.GetOrderRequest()}, nil
}

func main() {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))

	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	proto.RegisterOrderServiceServer(s, &server{})
	log.Println("Starting server on port", *port)

	if err := s.Serve(listen); err != nil {
		panic(err)
	}
}
