package monitor

import (
	"PingChef/src/domain"
	"PingChef/src/repositories"
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func StartMonitor(ctx context.Context, db *sql.DB, url string, interval time.Duration) error {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Monitoramento encerrado")
			return nil

		case <-ticker.C:
			status, responseTime, err := Ping(url)
			if err != nil {
				return err
			}

			new_event := domain.Event{
				ID:           uuid.New(),
				URL:          url,
				Status:       status,
				ResponseTime: responseTime,
			}

			repo := repositories.NewEventRepository(db)
			if err := repo.Create(new_event); err != nil {
				return err
			}

			fmt.Printf("URL: %s - Status: %s, Tempo de resposta: %s\n", url, status, responseTime)
		}
	}
}
