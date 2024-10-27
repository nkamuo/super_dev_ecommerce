package entity

type Product interface {
	GetId() string
	GetName() string
	GetDescription() string
	GetQuantityAvailable() int64
	GetPrice() int64
}
