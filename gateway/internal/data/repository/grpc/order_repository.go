package grpc

import (
	"context"
	"errors"
	"strconv"

	"github.com/superdev/ecommerce/gateway/internal/adapters/grpc/proto"
	"github.com/superdev/ecommerce/gateway/internal/domain/entity"
	"github.com/superdev/ecommerce/gateway/internal/domain/repository"
	"go.uber.org/zap"
)

type orderRepository struct {
	client proto.OrderServiceClient
	logger zap.Logger
}

func NewGrpcOrderRepository(
	client proto.OrderServiceClient,
	logger zap.Logger,
) repository.OrderRepository {
	return &orderRepository{
		client: client,
		logger: logger,
	}
}

func (s *orderRepository) Create(ctx *context.Context, entity entity.Order) error {
	var req proto.CreateOrderRequest
	for _, item := range entity.GetItems() {
		var itReq proto.OrderItem
		productId, err := strconv.ParseUint(item.GetProductId(), 10, 32)
		if err != nil {
			panic(err)
		}
		itReq.ProductId = int32(productId)
		itReq.Quantity = itReq.Quantity
		req.Items = append(req.Items, &itReq)
	}

	_, err := s.client.CreateOrder(context.Background(), &req)
	return err
}

func (s *orderRepository) Delete(ctx *context.Context, entity entity.Order) error {
	return errors.ErrUnsupported
}
func (s *orderRepository) FindAll(ctx *context.Context) ([]entity.Order, error) {
	var req proto.Empty
	res, err := s.client.ListOrders(context.Background(), &req)
	if err != nil {
		return nil, err
	}
	var orders []entity.Order
	for _, ord := range res.Orders {
		order := fromProtoToOrderEntity(ord)
		orders = append(orders, order)
	}
	return orders, err

}
func (s *orderRepository) FindById(ctx *context.Context, id string) (entity.Order, error) {
	return nil, errors.ErrUnsupported
}

func (s *orderRepository) Update(ctx *context.Context, entity entity.Order) error {
	return errors.ErrUnsupported
}
