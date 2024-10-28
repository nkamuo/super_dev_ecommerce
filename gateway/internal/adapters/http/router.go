package http

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/superdev/ecommerce/gateway/internal/adapters/http/handlers"
	"github.com/superdev/ecommerce/gateway/internal/adapters/http/middlewares"
)

func NewHTTPRouter(
	handlers []handlers.Handler,
	middlewares []middlewares.Middleware,
) (*gin.Engine, error) {

	engine := gin.Default()

	// Add middlewares
	for _, middleware := range middlewares {
		engine.Use(middleware.Handle())
	}
	for _, h := range handlers {
		methods := h.Methods()
		if len(methods) == 0 {
			return nil, errors.New("found a HTTP handler with empty method(verb) list. Please specify at least one method")
		} else if len(methods) == 1 {
			method := methods[0]
			switch strings.ToUpper(method) {
			case "GET":
				engine.GET(
					h.Pattern(),
					h.Handle,
				)
			case "PATCH":
				engine.PATCH(
					h.Pattern(),
					h.Handle,
				)
			case "PUT":
				engine.PUT(
					h.Pattern(),
					h.Handle,
				)

			case "POST":
				engine.POST(
					h.Pattern(),
					h.Handle,
				)
			case "DELETE":
				engine.DELETE(
					h.Pattern(),
					h.Handle,
				)
			}

		} else {
			engine.Match(
				h.Methods(),
				h.Pattern(),
				h.Handle,
			)
		}
	}

	return engine, nil
}
