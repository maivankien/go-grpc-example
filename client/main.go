package main

import (
	"context"
	"log"
	"time"

	proto "github.com/maivankien/go-grpc-example/grpc/order"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	address := "localhost:8000"

	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal("Could not connect to", address, err)
	}
	defer conn.Close()

	client := proto.NewOrderServiceClient(conn)

	ticket := time.NewTicker(2 * time.Second)

	defer ticket.Stop()

	for range ticket.C {
		orderId := "123"

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)

		r, err := client.NewOrder(ctx, &proto.OrderRequest{OrderRequest: orderId})

		if err != nil {
			log.Fatal("Could not place order", err)
		}
		log.Printf("Order confirmation: %s", r.GetOrderResponse())
		cancel()
	}
}
