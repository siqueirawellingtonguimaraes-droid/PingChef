package monitor

import (
	"PingChef/src/domain"
	"PingChef/src/modules/endpoints"
	"context"
	"net/http"
	"time"
)

func Ping(url string) domain.ResponsePing {
	// timeout da request
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// cria request com contexto
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return domain.ResponsePing{Status: endpoints.DOWN, ElapsedTime: 0, Error: err}
	}

	client := &http.Client{}

	start := time.Now()

	resp, err := client.Do(req)

	elapsed := time.Since(start)

	if err != nil {
		return domain.ResponsePing{Status: endpoints.DOWN, ElapsedTime: elapsed, Error: err}
	}
	defer resp.Body.Close()

	// considera UP apenas respostas 2xx e 3xx
	if resp.StatusCode >= 200 && resp.StatusCode < 400 {
		return domain.ResponsePing{Status: endpoints.UP, ElapsedTime: elapsed, Error: nil}
	}

	return domain.ResponsePing{Status: endpoints.DOWN, ElapsedTime: elapsed, Error: nil}
}
