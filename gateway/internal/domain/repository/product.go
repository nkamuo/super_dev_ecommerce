package repository

import "github.com/superdev/ecommerce/gateway/internal/domain/entity"

type ProductRepository interface {
	Repository[entity.Product]
}
