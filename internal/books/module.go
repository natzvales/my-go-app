package books

import (
	"github.com/gin-gonic/gin"
	"github.com/natz/go-lib-app/internal/container"
	"github.com/natz/go-lib-app/internal/server"
)

// BookModule holds the handler for the Book module
type BookModule struct {
	handler *BookHandler
}

func NewModule(c *container.Container) server.Module {

	repo := NewBookRepository(c.DB)
	service := NewBookService(repo)
	handler := NewBookHandler(service)

	return &BookModule{
		handler: handler,
	}
}

func (m *BookModule) RegisterRoutes(rg *gin.RouterGroup) {
	RegisterBookRoutes(rg, m.handler)
}

func init() {
	server.RegisterModule(NewModule)
}
