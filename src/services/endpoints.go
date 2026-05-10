package services

import (
	"PingChef/src/domain"
	"PingChef/src/infra"
	"PingChef/src/repositories"
	"fmt"

	"github.com/google/uuid"
)

func AddEndPoint(endpoint *domain.EndPoint) error {
	endpoint.ID = uuid.New()

	db, err := infra.ConnectDB()
	if err != nil {
		return err
	}

	defer db.Close()

	repo := repositories.NewRepository(db)

	if err = repo.Create(endpoint); err != nil {
		return err
	}

	fmt.Println("Novo endpoint adicionado:")
	fmt.Println("Nome:", endpoint.Name)
	fmt.Println("URL:", endpoint.URL)
	fmt.Println("Ticker:", endpoint.Interval)

	return nil
}

func RemoveEndPoint() error {
	return nil
}

func ListEndPoints() error {
	return nil
}
