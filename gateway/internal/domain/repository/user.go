package repository

import "github.com/superdev/ecommerce/gateway/internal/domain/entity"

type UserRepository interface {
	Repository[entity.User]
}
