package http

import (
	"go.uber.org/fx"
)

var Module = fx.Module("http_adapter",
	// fx.Invoke(func(*cobra.Command) {}),
	fx.Provide(
		NewHTTPRunCmd,
		NewHTTPServer,
		fx.Annotate(
			NewHTTPRouter,
			fx.ParamTags(`group:"app.http.handler"`, `group:"app.http.middleware"`),
		),
	),
)
