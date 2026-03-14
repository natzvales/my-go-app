package books

import (
	"strconv"

	"github.com/gin-gonic/gin"
	appErrors "github.com/natz/go-lib-app/internal/errors"
	"github.com/natz/go-lib-app/internal/response"
)

type BookHandler struct {
	service *BookService
}

func NewBookHandler(service *BookService) *BookHandler {
	return &BookHandler{service: service}
}

func (h *BookHandler) GetBooks(c *gin.Context) {
	books, err := h.service.GetAllBooks()
	if err != nil {
		c.Error(err)
		return
	}
	response.Success(c, books)
}

func (h *BookHandler) GetBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(appErrors.New(400, "invalid ID"))
		return
	}

	book, err := h.service.GetBook(id)
	if err != nil {
		c.Error(err)
		return
	}

	response.Success(c, book)
}

func (h *BookHandler) CreateBook(c *gin.Context) {

	var dto CreateBookDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		// response.BadRequest(c, "invalid request body")
		c.Error(appErrors.New(400, err.Error()))
		return
	}

	book, err := h.service.CreateBook(dto)
	if err != nil {
		c.Error(err)
		return
	}

	response.Created(c, book)
}

func (h *BookHandler) UpdateBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(appErrors.New(400, "invalid ID"))
		return
	}

	var dto UpdateBookDTO

	if err := c.ShouldBindJSON(&dto); err != nil {
		// response.BadRequest(c, "invalid request body")
		c.Error(appErrors.New(400, err.Error()))
		return
	}

	book, err := h.service.UpdateBook(id, dto)
	if err != nil {
		// response.Internal(c)
		c.Error(err)
		return
	}

	response.Success(c, book)

}

func (h *BookHandler) DeleteBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// response.BadRequest(c, "invalid ID")
		c.Error(appErrors.New(400, "invalid id"))
		return
	}

	err = h.service.DeleteBook(id)

	if err != nil {
		c.Error(err)
		return
	}

	response.Success(c, gin.H{"message": "book deleted"})
}
