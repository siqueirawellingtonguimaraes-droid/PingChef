package main

import (
	"PingChef/cmd/cli"
	"PingChef/src/config"
	"log"
)

func main() {
	if err := config.InitDB(); err != nil {
		log.Fatal(err)
	}
	cli.Execute()
}
