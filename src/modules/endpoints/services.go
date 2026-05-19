package endpoints

type EndpointService struct {
	repo *EndpointRepository
}

func NewEndpointService(repo *EndpointRepository) *EndpointService {
	return &EndpointService{repo: repo}
}

func (s *EndpointService) CreateEndpoint(dto CreateEndpointDTO) (ResponseEndpointDTO, error) {
	return ResponseEndpointDTO{}, nil
}

func (s *EndpointService) GetEndpoint(id string) (ResponseEndpointDTO, error) {
	return ResponseEndpointDTO{}, nil
}

func (s *EndpointService) GetAllEndpoints() ([]ResponseEndpointDTO, error) {
	return []ResponseEndpointDTO{}, nil
}

func (s *EndpointService) UpdateEndpoint(id string, dto CreateEndpointDTO) (ResponseEndpointDTO, error) {
	return ResponseEndpointDTO{}, nil
}

func (s *EndpointService) DeleteEndpoint(id string) error {
	return nil
}
