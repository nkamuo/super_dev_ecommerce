// grpc/order_client.go
package grpc

import (
	"log"

	"github.com/superdev/ecommerce/gateway/proto" // Replace with actual import path

	"google.golang.org/grpc"
)

func NewOrderClient() proto.OrderServiceClient {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to Order Service: %v", err)
	}
	return proto.NewOrderServiceClient(conn)
}
