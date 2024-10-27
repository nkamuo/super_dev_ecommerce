package repository

import (
	"context"

	"github.com/superdev/ecommerce/gateway/internal/domain/entity"
)

type ProductRepository interface {
	Repository[entity.Product]
	CheckProductAvailability(ctx context.Context, id string) (*bool, error)
}
