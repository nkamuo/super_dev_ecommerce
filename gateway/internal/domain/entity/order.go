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

func NewOrder(
	items []OrderItem,
) *order {
	return &order{
		items: items,
	}
}
func NewEmptyOrder() *order {
	return &order{}
}

func NewOrderItem(
	productId string,
	price int64,
	quantity int32,
) *orderItem {
	return &orderItem{
		productId: productId,
		price:     price,
		quantity:  quantity,
	}
}

func NewEmptyOrderItem() *orderItem {
	return &orderItem{}
}

type order struct {
	id    string
	items []OrderItem
}

func (s *order) GetId() string {
	return s.id
}

func (s *order) GetItems() []OrderItem {
	return s.items
}

func (s *order) AddItem(item OrderItem) {
	s.items = append(s.items, item)
}

//

type orderItem struct {
	productId string
	price     int64
	quantity  int32
}

func (s *orderItem) GetProductId() string {
	return s.productId
}
func (s *orderItem) SetProductId(productId string) {
	s.productId = productId
}

func (s *orderItem) GetPrice() int64 {
	return s.price
}
func (s *orderItem) SetPrice(price int64) {
	s.price = price
}

func (s *orderItem) SetQuantity(quantity int32) {
	s.quantity = quantity
}

func (s *orderItem) GetQuantity() int32 {
	return s.quantity
}
