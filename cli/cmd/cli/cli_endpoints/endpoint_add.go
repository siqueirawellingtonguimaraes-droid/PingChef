package cli_endpoints

import (
	"PingChef/cli/src/modules/endpoints"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

func NewAddEndpointCmd(service *endpoints.EndpointService) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "add",
		Short: "Cadastra um novo endpoint",
		Long:  "Cadastra um novo endpoint para ser monitorado pelo PingChef",

		Run: func(cmd *cobra.Command, args []string) {
			name, err := cmd.Flags().GetString("name")
			if err != nil {
				log.Fatal(err)
			}

			url, err := cmd.Flags().GetString("url")
			if err != nil {
				log.Fatal(err)
			}

			interval, err := cmd.Flags().GetInt("interval")
			if err != nil {
				log.Fatal(err)
			}

			endpointDTO := endpoints.CreateEndpointDTO{
				Name:     name,
				URL:      url,
				Interval: interval,
			}

			resp, err := service.CreateEndpoint(endpointDTO)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("URL adicionada com sucesso\n")
			fmt.Printf("URL: %s", resp.URL)
			fmt.Printf("ID: %s", resp.ID)
		},
	}

	cmd.Flags().String("name", "", "Nome da rota")
	cmd.Flags().String("url", "", "URL do endpoint")
	cmd.Flags().Int("interval", 10, "Tempo para cada Ping")

	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("url")

	return cmd
}
