package transport

import (
	"net/http"

	"github.com/spf13/cobra"
	"github.com/superdev/ecommerce/gateway/command"
)

func NewHTTPRunCmd(
	srv *http.Server,
) command.GenericCommand {

	cmd := cobra.Command{
		Use:   "run",
		Short: "STSX - RUn Server Service",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	return command.NewGenericCmd(&cmd, "root:service:run")

}
