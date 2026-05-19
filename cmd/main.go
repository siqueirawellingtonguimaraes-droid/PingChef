package main

import (
	"PingChef/cmd/cli"
	"PingChef/src/config"
	"PingChef/src/infra/db"
	"PingChef/src/infra/repo"
	"PingChef/src/modules/endpoints"
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
