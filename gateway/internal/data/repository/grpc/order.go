package grpc

import (
	"fmt"
	"strconv"

	"github.com/superdev/ecommerce/gateway/internal/adapters/grpc/proto"
	"github.com/superdev/ecommerce/gateway/internal/domain/entity"
)

type GrpcOrder struct {
	ID    uint64
	Items []*GrpcOrderItem
}

func (us *GrpcOrder) GetId() string {
	return fmt.Sprintf("%d", us.ID)
}
func (us *GrpcOrder) GetItems() []entity.OrderItem {
	var items []entity.OrderItem
	for _, pItem := range us.Items {
		items = append(items, pItem)
	}
	return items
}

//

type GrpcOrderItem struct {
	ProductId uint64
	Price     int64
	Quantity  int32
}

func (us *GrpcOrderItem) GetProductId() string {
	return fmt.Sprintf("%d", us.ProductId)
}
func (us *GrpcOrderItem) GetPrice() int64 {
	return us.Price
}

func (us *GrpcOrderItem) GetQuantity() int32 {
	return us.Quantity
}

//

func fromProtoToOrderEntity(pOrder *proto.Order) entity.Order {
	var order GrpcOrder
	order.ID = uint64(pOrder.Id)
	var items []*GrpcOrderItem
	for _, pItem := range pOrder.Items {
		var item GrpcOrderItem
		item.ProductId = item.ProductId
		item.Quantity = pItem.Quantity
		items = append(items, &item)
	}
	order.Items = items
	return &order
}

func fromOrderEntity(order entity.Order) *GrpcOrder {
	if gOrder, ok := order.(*GrpcOrder); ok {
		return gOrder
	} else {
		ID, err := strconv.ParseUint(order.GetId(), 10, 64)
		if nil != err {
			panic(err)
		}
		gOrder.ID = ID
		return gOrder
	}
}

func toOrderEntity(order *GrpcOrder) entity.Order {
	return order
}
