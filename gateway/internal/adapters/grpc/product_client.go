// grpc/product_client.go
package grpc

import (
	"fmt"

	"github.com/superdev/ecommerce/gateway/internal/adapters/grpc/proto"
	"github.com/superdev/ecommerce/gateway/internal/config"
	"go.uber.org/zap"

	"google.golang.org/grpc"
)

func NewProductGrpcClient(
	conf *config.Config,
	logger zap.Logger,
) proto.ProductServiceClient {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		logger.Error(fmt.Sprintf("Could not connect to Product Service: %v", err))
	}
	return proto.NewProductServiceClient(conn)
}
