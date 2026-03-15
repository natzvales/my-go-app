package books

import (
	"github.com/gin-gonic/gin"
	"github.com/natz/go-lib-app/internal/middleware"
)

func RegisterBookRoutes(rg *gin.RouterGroup, handler *BookHandler) {

	books := rg.Group("/books")

	{
		books.GET("", handler.GetBooks)
		books.GET("/:id", handler.GetBook)
		books.Use(middleware.AuthMiddleware())
		books.POST("", middleware.RequireRole("librarian", "admin"), handler.CreateBook)
		books.PUT("/:id", middleware.RequireRole("librarian", "admin"), handler.UpdateBook)
		books.DELETE("/:id", middleware.RequireRole("admin"), handler.DeleteBook)
	}
}
