package events

import (
	"github.com/spf13/cobra"
)

var eventRunAllCmd = &cobra.Command{
	Use:   "run-all",
	Short: "Inicia o monitoramento de todos os endpoints cadastrados",
	Long:  "Inicia o monitoramento de todos os endpoints cadastrados, criando eventos de monitoramento para cada um deles",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	EventCmd.AddCommand(eventRunAllCmd)
}
