package cli

import (
	"PingChef/cmd/cli/routes"

	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	cfgFile     string
	userLicense string

	rootCmd = &cobra.Command{
		Use:   "PingChef",
		Short: "Monitore se seu sistema está rodando",
	}
)

func init() {
	rootCmd.AddCommand(routes.RouteCmd)
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
