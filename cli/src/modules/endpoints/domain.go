package endpoints

import (
	"fmt"
	"net/url"
	"time"

	"github.com/google/uuid"
)

type Status string

const (
	UP   Status = "UP"
	DOWN Status = "DOWN"
)

type Event struct {
	ID           uuid.UUID
	Status       Status
	StatusCode   int
	ResponseTime time.Duration
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

func (e *Endpoint) AddEvent(event Event) {
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
