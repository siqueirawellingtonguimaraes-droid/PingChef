package cli_endpoints

import (
	"PingChef/cli/src/modules/endpoints"

	"github.com/spf13/cobra"
)

func NewEndpointCmd(service *endpoints.EndpointService) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "route",
		Short: "Gerencia rotas monitoradas",
	}

	cmd.AddCommand(NewAddEndpointCmd(service))
	cmd.AddCommand(NewListEndpointCmd(service))

	return cmd
}
