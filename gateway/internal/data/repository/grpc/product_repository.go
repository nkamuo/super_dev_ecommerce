package grpc

import (
	"context"
	"errors"

	"github.com/superdev/ecommerce/gateway/internal/adapters/grpc/proto"
	"github.com/superdev/ecommerce/gateway/internal/domain/entity"
	"github.com/superdev/ecommerce/gateway/internal/domain/repository"
	"go.uber.org/zap"
)

type productRepository struct {
	client proto.ProductServiceClient
	logger zap.Logger
}

func NewGrpcProductRepository(
	client proto.ProductServiceClient,
	logger zap.Logger,
) repository.ProductRepository {
	return &productRepository{
		client: client,
		logger: logger,
	}
}

func (s *productRepository) Create(ctx *context.Context, entity entity.Product) error {
	var req proto.CreateProductRequest

	req.AvailableQuantity = int32(entity.GetQuantityAvailable())
	req.Name = entity.GetName()
	req.Description = entity.GetDescription()
	req.Price = float32(entity.GetPrice())
	_, err := s.client.CreateProduct(context.Background(), &req)
	return err
}

func (s *productRepository) Delete(ctx *context.Context, entity entity.Product) error {
	return errors.ErrUnsupported
}
func (s *productRepository) FindAll(ctx *context.Context) ([]entity.Product, error) {
	var req proto.Empty
	res, err := s.client.ListProducts(context.Background(), &req)
	if err != nil {
		return nil, err
	}
	var products []entity.Product
	for _, ord := range res.Products {
		product := fromProtoToProductEntity(ord)
		products = append(products, product)
	}
	return products, err

}
func (s *productRepository) FindById(ctx *context.Context, id string) (entity.Product, error) {
	return nil, errors.ErrUnsupported
}

func (s *productRepository) Update(ctx *context.Context, entity entity.Product) error {
	return errors.ErrUnsupported
}
