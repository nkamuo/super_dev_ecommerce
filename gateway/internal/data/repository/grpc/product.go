package grpc

import (
	"fmt"
	"strconv"

	"github.com/superdev/ecommerce/gateway/internal/domain/entity"
	"github.com/superdev/ecommerce/gateway/proto"
)

type GrpcProduct struct {
	ID                uint64
	Name              string
	Description       string
	Price             float32
	AvailableQuantity int32
}

func (us *GrpcProduct) GetId() string {
	return fmt.Sprintf("%d", us.ID)
}

func (us *GrpcProduct) GetName() string {
	return us.Name
}
func (us *GrpcProduct) GetDescription() string {
	return us.Description
}
func (us *GrpcProduct) GetQuantityAvailable() int64 {
	return int64(us.AvailableQuantity)
}
func (us *GrpcProduct) GetPrice() int64 {
	return int64(us.Price)
}

func fromProtoToProductEntity(pProduct *proto.Product) entity.Product {
	var product GrpcProduct
	product.ID = uint64(pProduct.Id)
	product.Name = pProduct.Name
	product.Description = pProduct.Description
	product.AvailableQuantity = pProduct.AvailableQuantity
	product.Price = pProduct.Price
	return &product
}

func fromProductEntity(product entity.Product) *GrpcProduct {
	if gProduct, ok := product.(*GrpcProduct); ok {
		return gProduct
	} else {
		ID, err := strconv.ParseUint(product.GetId(), 10, 64)
		if nil != err {
			panic(err)
		}
		gProduct.ID = ID
		return gProduct
	}
}

func toProductEntity(product *GrpcProduct) entity.Product {
	return product
}
