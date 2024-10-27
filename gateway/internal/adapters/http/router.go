package transport

import (
	"github.com/gin-gonic/gin"
)

func NewHTTPRouter(
	handlers []Handler,
	middlewares []Middleware,
) (*gin.Engine, error) {

	engine := gin.Engine{}

	// Add middlewares
	for _, middleware := range middlewares {
		engine.Use(middleware.Handle())
	}
	for _, h := range handlers {
		engine.Match(
			h.Methods(),
			h.Pattern(),
			h.Handle,
		)
	}

	return &engine, nil
}
