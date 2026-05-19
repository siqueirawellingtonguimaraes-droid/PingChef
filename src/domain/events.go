package domain

import (
	"PingChef/src/modules/endpoints"
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID           uuid.UUID
	URL          string
	Status       endpoints.Status
	ResponseTime time.Duration
}
