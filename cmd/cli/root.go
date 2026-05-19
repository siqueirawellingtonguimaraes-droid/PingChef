package cli

import (
	"PingChef/cmd/cli/cli_endpoints"
	"PingChef/src/modules/endpoints"

	"github.com/spf13/cobra"
)

func RootCMD(service *endpoints.EndpointService) *cobra.Command {
	var (
		cmd = &cobra.Command{
			Use:   "PingChef",
			Short: "Monitore se seu sistema está rodando",
		}
	)

	cmd.AddCommand(cli_endpoints.NewEndpointCmd(service))

	return cmd
}

// Execute executes the root command.
func Execute(service *endpoints.EndpointService) error {
	rootCMD := RootCMD(service)

	return rootCMD.Execute()
}
