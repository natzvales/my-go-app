package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/natz/go-lib-app/internal/middleware"
)

func RegisterRoutes(rg *gin.RouterGroup, handler *Handler, service *Service) {

	auth := rg.Group("/auth")

	auth.POST("/register", middleware.ValidateBody[RegisterDTO](), handler.Register)
	auth.POST("/login", middleware.ValidateBody[LoginDTO](), handler.Login)

	auth.GET("/me",
		middleware.AuthMiddleware(),
		handler.Me,
	)
}
