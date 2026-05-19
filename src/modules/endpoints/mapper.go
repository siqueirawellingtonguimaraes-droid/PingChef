package endpoints

import (
	"time"
)

func EndpointToResponseDTO(endpoint Endpoint) ResponseEndpointDTO {
	eventsDTO := make([]ResponseEventDTO, len(endpoint.Events))
	for i, event := range endpoint.Events {
		eventsDTO[i] = ResponseEventDTO{
			ID:           event.ID.String(),
			EndpointID:   endpoint.ID.String(),
			Status:       string(event.Status),
			StatusCode:   event.StatusCode,
			ResponseTime: int(event.ResponseTime.Milliseconds()),
		}
	}
	return ResponseEndpointDTO{
		ID:       endpoint.ID.String(),
		Name:     endpoint.Name,
		URL:      endpoint.URL,
		Interval: int(endpoint.Interval.Seconds()),
		Events:   eventsDTO,
	}
}

func CreateToEndpoint(dto CreateEndpointDTO) (Endpoint, error) {
	return NewEndpoint(dto.Name, dto.URL, time.Second*time.Duration(dto.Interval))
}
