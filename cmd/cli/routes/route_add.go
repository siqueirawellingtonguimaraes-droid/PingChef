package routes

import (
	"PingChef/src/domain"
	"PingChef/src/services"
	"log"
	"time"

	"github.com/spf13/cobra"
)

var (
	routeName string
	routeURL  string
	interval  int
)

var routeAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Cadastra um novo endpoint",
	Long:  "Cadastra um novo endpoint para ser monitorado pelo PingChef",

	Run: func(cmd *cobra.Command, args []string) {

		endpoint := domain.EndPoint{
			Name:     routeName,
			URL:      routeURL,
			Interval: time.Second * time.Duration(interval),
		}

		if err := services.AddEndPoint(&endpoint); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	routeAddCmd.Flags().StringVar(
		&routeName,
		"name",
		"",
		"Nome da rota",
	)

	routeAddCmd.Flags().StringVar(
		&routeURL,
		"url",
		"",
		"URL do endpoint",
	)

	routeAddCmd.Flags().IntVar(
		&interval,
		"interval",
		10,
		"Tempo para cada Ping",
	)

	routeAddCmd.MarkFlagRequired("name")
	routeAddCmd.MarkFlagRequired("url")

	RouteCmd.AddCommand(routeAddCmd)
}
