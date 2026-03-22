package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/natz/go-lib-app/internal/middleware"
	"github.com/natz/go-lib-app/internal/response"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Register(c *gin.Context) {

	var dto RegisterDTO

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.Error(err)
		return
	}

	user, err := h.service.Register(dto)
	if err != nil {
		c.Error(err)
		return
	}

	response.Created(c, user)
}

func (h *Handler) Login(c *gin.Context) {

	var dto LoginDTO

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.Error(err)
		return
	}

	token, err := h.service.Login(dto)
	if err != nil {
		c.Error(err)
		return
	}

	response.Success(c, gin.H{
		"token": token,
	})
}

func (h *Handler) Me(c *gin.Context) {

	user := middleware.GetUser(c)

	response.Success(c, user)
}

// func (h *Handler) GetProfile(c *gin.Context) {

// 	user := c.Get("user")
