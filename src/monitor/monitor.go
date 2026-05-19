package monitor

import (
	ep "PingChef/src/modules/endpoints"
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/google/uuid"
)

type Monitor struct {
	repo ep.EndpointRepository
}

func NewMonitor(repo ep.EndpointRepository) *Monitor {
	return &Monitor{repo: repo}
}

func (m *Monitor) RunMonitor() {

	endpoints, err := m.repo.FindAllEndpoints()
	if err != nil {
		log.Fatal("Erro ao consultar endpoints:", err)
	}

	var wg sync.WaitGroup
	ctx := context.Background()

	for _, endpoint := range endpoints {

		wg.Add(1)

		go func(e ep.Endpoint) {
			defer wg.Done()

			if err := m.startMonitor(ctx, e); err != nil {
				log.Printf("Erro ao iniciar monitoramento para %s: %v", e.URL, err)
			}
		}(endpoint)
	}

	wg.Wait()

}

func (m *Monitor) startMonitor(ctx context.Context, endpoint ep.Endpoint) error {
	ticker := time.NewTicker(endpoint.Interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Monitoramento encerrado")
			return nil

		case <-ticker.C:
			resp := Ping(endpoint.URL)
			if resp.Error != nil {
				return resp.Error
			}

			new_event := ep.Event{
				ID:           uuid.New(),
				Status:       resp.Status,
				StatusCode:   resp.StatusCode,
				ResponseTime: resp.ElapsedTime,
			}

			if err := m.repo.SaveEvent(new_event, endpoint.ID.String()); err != nil {
				log.Printf("Erro ao salvar evento para %s: %v", endpoint.URL, err)
			}

			fmt.Printf("URL: %s - Status: %s, Tempo de resposta: %s\n", endpoint.URL, resp.Status, resp.ElapsedTime)
		}
	}
}
