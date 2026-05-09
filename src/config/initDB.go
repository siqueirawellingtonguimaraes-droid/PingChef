package config

import "PingChef/src/infra"

func InitDB() error {
	db, err := infra.ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	return nil

}
