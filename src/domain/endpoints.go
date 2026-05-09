package domain

import (
	"time"

	"github.com/google/uuid"
)

type EndPoint struct {
	ID     uuid.UUID
	URL    string
	Ticker time.Duration
}
