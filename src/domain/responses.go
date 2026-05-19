package domain

import (
	"PingChef/src/modules/endpoints"
	"time"
)

type ResponsePing struct {
	Status      endpoints.Status
	StatusCode  int
	ElapsedTime time.Duration
	Error       error
}
