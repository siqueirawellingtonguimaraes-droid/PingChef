package services

import (
	"PingChef/src/domain"
	"PingChef/src/infra"
	"PingChef/src/monitor"
	"context"
	"log"
	"sync"
	"time"
)

func RunMonitor() {
	db, err := infra.ConnectDB()
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}

	rows, err := db.Query("SELECT url, interval FROM endpoints")
	if err != nil {
		log.Fatal("Erro ao consultar endpoints:", err)
	}

	endpoints := []domain.EndPoint{}
	for rows.Next() {
		var url string
		var interval time.Duration
		if err := rows.Scan(&url, &interval); err != nil {
			log.Fatal("Erro ao ler dados dos endpoints:", err)
		}
		endpoint := domain.EndPoint{
			URL:      url,
			Interval: interval,
		}
		endpoints = append(endpoints, endpoint)
	}

	rows.Close()

	var wg sync.WaitGroup
	ctx := context.Background()

	for _, endpoint := range endpoints {

		wg.Add(1)

		go func(ep domain.EndPoint) {
			defer wg.Done()

			if err := monitor.StartMonitor(ctx, db, ep.URL, ep.Interval); err != nil {
				log.Printf("Erro ao iniciar monitoramento para %s: %v", ep.URL, err)
			}
		}(endpoint)
	}

	wg.Wait()

}
