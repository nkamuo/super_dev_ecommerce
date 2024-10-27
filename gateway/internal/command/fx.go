package command

import (
	"github.com/spf13/cobra"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

var Module = fx.Module("command",
	fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
		return &fxevent.ZapLogger{Logger: log}
	}),
	fx.Invoke(func(*cobra.Command) {}),
	fx.Provide(
		// AsCommand(NewHelloCmd),
		// AsCommand(NewHiCmd),
		fx.Annotate(
			NewRootCmd,
			fx.ParamTags(`group:"app.commands"`),
		),
	),
)
