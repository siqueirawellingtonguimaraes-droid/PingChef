package endpoints

type CreateEventDTO struct {
	EndpointID   string `json:"endpoint_id"`
	Status       string `json:"status"`
	StatusCode   int    `json:"status_code"`
	ResponseTime int    `json:"response_time"`
}

type CreateEndpointDTO struct {
	Name     string `json:"name"`
	URL      string `json:"url"`
	Interval int    `json:"interval"`
}

type ResponseEventDTO struct {
	ID           string `json:"id"`
	EndpointID   string `json:"endpoint_id"`
	Status       string `json:"status"`
	StatusCode   int    `json:"status_code"`
	ResponseTime int    `json:"response_time"`
}

type ResponseEndpointDTO struct {
	ID       string             `json:"id"`
	Name     string             `json:"name"`
	URL      string             `json:"url"`
	Interval int                `json:"interval"`
	Events   []ResponseEventDTO `json:"events"`
}

type UpdateEndpointDTO struct {
	Name     string `json:"name"`
	URL      string `json:"url"`
	Interval int    `json:"interval"`
}
