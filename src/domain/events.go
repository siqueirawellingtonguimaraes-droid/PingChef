package domain

import (
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
	URL          string
	Status       Status
	ResponseTime time.Duration
}
