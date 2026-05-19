package endpoints

import (
	"fmt"
	"net/url"
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID           uuid.UUID
	Status       string
	StatusCode   int
	ResponseTime int
}

type Endpoint struct {
	ID       uuid.UUID
	Name     string
	URL      string
	Interval time.Duration
	Events   []Event
}

func NewEndpoint(name, url string, interval time.Duration) (Endpoint, error) {
	if !isValidHTTPURL(url) {
		return Endpoint{}, fmt.Errorf("invalid HTTP URL")
	}

	return Endpoint{
		ID:       uuid.New(),
		Name:     name,
		URL:      url,
		Interval: interval,
		Events:   []Event{},
	}, nil
}

func (e *Endpoint) AddEvent(status string, statusCode, responseTime int) {
	event := Event{
		ID:           uuid.New(),
		Status:       status,
		StatusCode:   statusCode,
		ResponseTime: responseTime,
	}
	e.Events = append(e.Events, event)
}

func isValidHTTPURL(rawURL string) bool {
	u, err := url.Parse(rawURL)
	if err != nil {
		return false
	}

	if u.Scheme != "http" && u.Scheme != "https" {
		return false
	}

	if u.Host == "" {
		return false
	}

	return true
}
