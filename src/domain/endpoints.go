package domain

import (
	"time"

	"github.com/google/uuid"
)

type EndPoint struct {
	ID     uuid.UUID
	Name   string
	URL    string
	Ticker time.Duration
}
