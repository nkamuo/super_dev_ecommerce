package main

import (
	"github.com/superdev/ecommerce/gateway/internal/adapters/grpc"
	"github.com/superdev/ecommerce/gateway/internal/adapters/http"
	"github.com/superdev/ecommerce/gateway/internal/config"
	"github.com/superdev/ecommerce/gateway/internal/data"
	database "github.com/superdev/ecommerce/gateway/internal/infrastructure/db"
	"github.com/superdev/ecommerce/gateway/internal/infrastructure/logger"
	"go.uber.org/fx"
)

func main() {
	var app = fx.New(
		// --- APP CONFIG
		config.Module,
		// --- GRPC TRANSPORT
		grpc.Module,
		// -- HTTP TRANSPORT
		http.Module,
		fx.Provide(http.NewHTTPServerRunner),
		fx.Invoke(func(runner http.HttpServerRunner) {
		}),
		// --- DATA MODULE
		data.Module,

		// --- INFRASTRUCTURE
		database.Module,
		logger.Module,
	)

	app.Run()
}
