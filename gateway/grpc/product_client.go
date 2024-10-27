// grpc/product_client.go
package grpc

import (
	"log"

	"github.com/superdev/ecommerce/gateway/proto" // Replace with actual import path

	"google.golang.org/grpc"
)

func NewProductClient() proto.ProductServiceClient {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to Product Service: %v", err)
	}
	return proto.NewProductServiceClient(conn)
}
