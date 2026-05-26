package endpoints

type EndpointRepository interface {
	SaveEndpoint(endpoint Endpoint) error
	SaveEvent(event Event, endpointID string) error
	FindByID(id string) (Endpoint, error)
	FindAllEndpoints() ([]Endpoint, error)
	FindAllEvents() ([]Event, error)
	DeleteEndpoint(id string) error
	UpdateEndpoint(endpoint Endpoint) error
}
