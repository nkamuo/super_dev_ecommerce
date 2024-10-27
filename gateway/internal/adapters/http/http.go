package transport

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/superdev/ecommerce/gateway/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewHTTPServer(
	router *gin.Engine,
	logger *zap.Logger,
	conf *config.Config,
	lc fx.Lifecycle,
) (*http.Server, error) {

	port := conf.AppPort
	if 0 == port {
		port = 8080
	}
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}
	// Start the server when the lifecycle starts
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}
			logger.Debug(fmt.Sprintf("Starting HTTP server at %s", srv.Addr))
			go srv.Serve(ln)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Debug(fmt.Sprintf("Shutting down HTTP server at %s", srv.Addr))
			return srv.Shutdown(ctx)
		},
	})

	return srv, nil
}
