package server

import (
	"github.com/gin-gonic/gin"
	"github.com/natz/go-lib-app/internal/container"
)

// Module defines a feature that can register its routes
type Module interface {
	RegisterRoutes(rg *gin.RouterGroup)
}

type ModuleFactory func(c *container.Container) Module
