package orderhandlers

import (
	"github.com/gin-gonic/gin"
	"github.com/superdev/ecommerce/gateway/internal/adapters/http"
	"github.com/superdev/ecommerce/gateway/internal/config"
)

type createOrderHandler struct {
}

func NewCreateOrderHandler(
	conf *config.Config,
) http.Handler {
	return &createOrderHandler{}
}

func (s *createOrderHandler) Pattern() string {
	return "/orders"
}

func (s *createOrderHandler) Methods() []string {
	return []string{"POST"}
}

func (s *createOrderHandler) Handle(c *gin.Context) {
}
