package server

import (
	"github.com/gin-gonic/gin"
)

func NewServer(modules []Module) *gin.Engine {

	router := gin.Default()

	apiGroup := router.Group("/api") // base group

	//Register all modules automatically
	for _, module := range modules {
		module.RegisterRoutes(apiGroup)
	}

	return router
}
