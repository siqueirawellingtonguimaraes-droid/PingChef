package cli_endpoints

import (
	"PingChef/cli/src/modules/endpoints"

	"github.com/spf13/cobra"
)

func NewListEndpointCmd(service *endpoints.EndpointService) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "list",
		Short: "Lista os endpoints cadastrados",
		Long:  "Lista os endpoints cadastrados para serem monitorados pelo PingChef",

		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	return cmd

}

var routeListCmd = &cobra.Command{}
