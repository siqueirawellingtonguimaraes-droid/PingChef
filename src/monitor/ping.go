package monitor

import (
	"PingChef/src/domain"
	"context"
	"net/http"
	"time"
)

func Ping(url string) (domain.Status, time.Duration, error) {
	// timeout da request
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// cria request com contexto
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return domain.DOWN, 0, err
	}

	client := &http.Client{}

	start := time.Now()

	resp, err := client.Do(req)

	elapsed := time.Since(start)

	if err != nil {
		return domain.DOWN, elapsed, err
	}
	defer resp.Body.Close()

	// considera UP apenas respostas 2xx e 3xx
	if resp.StatusCode >= 200 && resp.StatusCode < 400 {
		return domain.UP, elapsed, nil
	}

	return domain.DOWN, elapsed, nil
}
