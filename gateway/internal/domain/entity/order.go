package entity

type Order interface {
	GetId() string
	GetItems() []OrderItem
}

type OrderItem interface {
	GetProductId() string
	GetPrice() int64
	GetQuantity() int32
}
