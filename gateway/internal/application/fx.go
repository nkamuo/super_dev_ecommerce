package application

import (
	"github.com/superdev/ecommerce/gateway/internal/application/service"
	"go.uber.org/fx"
)

var Module = fx.Module("http_adapter",
	// fx.Invoke(func(*cobra.Command) {}),
	fx.Provide(
		service.NewProductService,
		service.NewOrderService,
	),
)
