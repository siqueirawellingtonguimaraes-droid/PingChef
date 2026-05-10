package config

import (
	"PingChef/src/infra"
	"os"
)

func InitDB() error {
	db, err := infra.ConnectDB()

	if err != nil {
		return err
	}
	defer db.Close()

	sqlBytes, err := os.ReadFile("sql/db.sql")

	if err != nil {
		return err
	}

	query := string(sqlBytes)

	_, err = db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
