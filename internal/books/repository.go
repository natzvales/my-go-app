package books

import (
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	GetAll() ([]Book, error)
	GetByID(id int) (Book, error)
	Create(book *Book) error
	Update(book *Book) error
	Delete(id int) error
}

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{
		db: db,
	}
}

func (r *BookRepository) GetAll() ([]Book, error) {
	var books []Book
	result := r.db.Find(&books)
	return books, result.Error
}

func (r *BookRepository) GetByID(id int) (Book, error) {
	var book Book
	result := r.db.First(&book, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return book, errors.New("book not found")
	}
	return book, result.Error
}

func (r *BookRepository) Create(book *Book) error {
	return r.db.Create(book).Error
}

func (r *BookRepository) Update(book *Book) error {
	return r.db.Save(book).Error
}

func (r *BookRepository) Delete(id int) error {
	// result := r.db.Delete(&Book{}, id)
	// if result.RowsAffected == 0 {
	// 	return errors.New("book not found")
	// }
	// return result.Error
	return r.db.Delete(&Book{}, id).Error
}
