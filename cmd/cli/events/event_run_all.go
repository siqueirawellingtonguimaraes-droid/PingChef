package events

import (
	"PingChef/src/services"

	"github.com/spf13/cobra"
)

var eventRunAllCmd = &cobra.Command{
	Use:   "run-all",
	Short: "Inicia o monitoramento de todos os endpoints cadastrados",
	Long:  "Inicia o monitoramento de todos os endpoints cadastrados, criando eventos de monitoramento para cada um deles",
	RunE: func(cmd *cobra.Command, args []string) error {
		services.RunMonitor()
		return nil
	},
}

func init() {
	EventCmd.AddCommand(eventRunAllCmd)
}
