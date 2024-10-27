// grpc/order_client.go
package grpc

import (
	"fmt"

	"github.com/superdev/ecommerce/gateway/internal/adapters/grpc/proto"
	"github.com/superdev/ecommerce/gateway/internal/config"
	"go.uber.org/zap"

	"google.golang.org/grpc"
)

func NewOrderGrpcClient(
	conf *config.Config,
	logger zap.Logger,
) proto.OrderServiceClient {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		logger.Error(fmt.Sprintf("Could not connect to Order Service: %v", err))
	}
	return proto.NewOrderServiceClient(conn)
}
