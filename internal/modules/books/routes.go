package books

import (
	"github.com/gin-gonic/gin"
	"github.com/natz/go-lib-app/internal/middleware"
)

func RegisterBookRoutes(rg *gin.RouterGroup, handler *BookHandler) {

	books := rg.Group("/books")

	{
		books.GET("", handler.GetBooks)
		books.GET("/:id", middleware.ValidateUUIDParam("id"), handler.GetBook)
		books.Use(middleware.AuthMiddleware())
		books.POST("", middleware.RequireRole("librarian", "admin"), middleware.ValidateBody[CreateBookDTO](), handler.CreateBook)
		books.PUT("/:id", middleware.RequireRole("librarian", "admin"), middleware.ValidateUUIDParam("id"), middleware.ValidateBody[UpdateBookDTO](), handler.UpdateBook)
		books.DELETE("/:id", middleware.RequireRole("admin"), middleware.ValidateUUIDParam("id"), handler.DeleteBook)
	}
}
