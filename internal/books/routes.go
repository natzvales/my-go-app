package books

import "github.com/gin-gonic/gin"

func RegisterBookRoutes(rg *gin.RouterGroup, handler *BookHandler) {

	books := rg.Group("/books")

	{
		books.GET("", handler.GetBooks)
		books.GET("/:id", handler.GetBook)
		books.POST("", handler.CreateBook)
		books.PUT("/:id", handler.UpdateBook)
		books.DELETE("/:id", handler.DeleteBook)
	}
}
