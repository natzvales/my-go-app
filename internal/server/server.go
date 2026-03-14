package server

import (
	"github.com/gin-gonic/gin"
	"github.com/natz/go-lib-app/internal/middleware"
)

func NewServer(modules []Module) *gin.Engine {

	router := gin.Default()

	//Global middlewares
	router.Use(middleware.Recovery())
	router.Use(middleware.Logger())
	router.Use(middleware.RequestID())
	router.Use(middleware.ErrorHandler())

	apiGroup := router.Group("/api") // base group

	//Register all modules automatically
	for _, module := range modules {
		module.RegisterRoutes(apiGroup)
	}

	return router
}
