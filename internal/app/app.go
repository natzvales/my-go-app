package app

import (
	"github.com/natz/go-lib-app/internal/books"
)

type App struct {
	BookHandler *books.BookHandler
	BookService *books.BookService
}
