package endpoints

import "github.com/gin-gonic/gin"

type Module struct {
	Service *EndpointService
}

func NewModule(service *EndpointService) *Module {
	return &Module{Service: service}
}

func (m *Module) Register(r *gin.Engine) {
	SetupEndpointsRoutes(r, m.Service)
}
