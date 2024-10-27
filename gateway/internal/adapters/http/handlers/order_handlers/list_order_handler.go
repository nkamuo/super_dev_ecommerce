package orderhandlers

import (
	"fmt"
	_http "net/http"

	"github.com/gin-gonic/gin"
	"github.com/superdev/ecommerce/gateway/internal/adapters/http"
	"github.com/superdev/ecommerce/gateway/internal/config"
	"github.com/superdev/ecommerce/gateway/internal/domain/service"
)

type listOrderHandler struct {
	orderService service.OrderService
}

func NewListOrderHandler(
	orderService service.OrderService,
	conf *config.Config,
) http.Handler {
	return &listOrderHandler{
		orderService: orderService,
	}
}

func (s *listOrderHandler) Pattern() string {
	return "/orders/:id"
}

func (s *listOrderHandler) Methods() []string {
	return []string{"DELETE"}
}

func (s *listOrderHandler) Handle(c *gin.Context) {

	orders, err := s.orderService.ListOrders()
	if err != nil {
		c.JSON(_http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": fmt.Sprintf("Could not list Orders: %s", err.Error()),
		})
		return
	}

	c.JSON(_http.StatusOK, gin.H{
		"status":  "success",
		"message": ("Order listd successfully"),
		"data":    orders, //order,
	})

}
