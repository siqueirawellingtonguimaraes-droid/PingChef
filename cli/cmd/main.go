package main

import (
	"PingChef/cli/cmd/cli"
	"PingChef/cli/src/config"
	"PingChef/cli/src/infra/db"
	"PingChef/cli/src/infra/repo"
	"PingChef/cli/src/modules/endpoints"
	"log"
)

func main() {
	config.InitDB()
	db, err := db.NewSQLiteConnection()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	repo := repo.NewSQLiteEndpointRepo(db)
	service := endpoints.NewEndpointService(repo)
	/* module := endpoints.NewModule(service) */

	cli.Execute(service)
}
