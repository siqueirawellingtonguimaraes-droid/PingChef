package routes

import (
	"PingChef/src/services"
	"log"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

var routeListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lista os endpoints cadastrados",
	Long:  "Lista os endpoints cadastrados para serem monitorados pelo PingChef",

	RunE: func(cmd *cobra.Command, args []string) error {
		endpoints, err := services.ListEndPoints()
		if err != nil {
			log.Fatal("Erro ao listar endpoints:", err)
		}

		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)

		t.AppendHeader(table.Row{
			"ID",
			"Name",
			"URL",
			"Interval",
		})

		for _, endpoint := range endpoints {
			t.AppendRow(table.Row{
				endpoint.ID,
				endpoint.Name,
				endpoint.URL,
				endpoint.Interval,
			})
		}
		t.Render()
		return nil
	},
}

func init() {
	RouteCmd.AddCommand(routeListCmd)
}
