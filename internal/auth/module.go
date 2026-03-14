package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/natz/go-lib-app/internal/container"
	"github.com/natz/go-lib-app/internal/server"
)

type Module struct {
	handler *Handler
	service *Service
}

func NewModule(c *container.Container) server.Module {

	// run migration
	// Migrate(c.DB)

	repo := NewRepository(c.DB)
	service := NewService(repo)
	handler := NewHandler(service)

	return &Module{
		handler: handler,
		service: service,
	}
}

func (m *Module) RegisterRoutes(rg *gin.RouterGroup) {
	RegisterRoutes(rg, m.handler, m.service)
}

func init() {
	server.RegisterModule(NewModule)
}
