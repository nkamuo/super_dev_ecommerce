package http

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

// import "net/http"

type Handler interface {
	Handle(c *gin.Context)
	Pattern() string
	Methods() []string
}

func AsHttpRoute(f any, anns ...fx.Annotation) any {
	anns = append(anns,
		fx.As(new(Handler)),
		fx.ResultTags(`group:"app.http.handler"`))

	return fx.Annotate(
		f,
		anns...,
	)
}
