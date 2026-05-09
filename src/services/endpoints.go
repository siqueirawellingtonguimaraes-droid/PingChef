package services

import (
	"PingChef/src/domain"
	"fmt"

	"github.com/google/uuid"
)

func AddEndPoint(endpoint *domain.EndPoint) error {
	endpoint.ID = uuid.New()

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
