package books

import "github.com/google/uuid"

type BookService struct {
	repo *BookRepository
}

func NewBookService(repo *BookRepository) *BookService {
	return &BookService{repo: repo}
}

func (s *BookService) GetAllBooks() ([]Book, error) {
	return s.repo.GetAll()
}

func (s *BookService) GetBook(id uuid.UUID) (Book, error) {
	return s.repo.GetByID(id)
}

func (s *BookService) CreateBook(dto CreateBookDTO) (Book, error) {
	book := Book{
		Title:  dto.Title,
		Author: dto.Author,
	}

	err := s.repo.Create(&book)

	return book, err
}

func (s *BookService) UpdateBook(id uuid.UUID, dto UpdateBookDTO) (Book, error) {
	book, err := s.repo.GetByID(id)
	if err != nil {
		return book, err
	}

	book.Title = dto.Title
	book.Author = dto.Author

	err = s.repo.Update(&book)

	return book, err
}

func (s *BookService) DeleteBook(id uuid.UUID) error {
	return s.repo.Delete(id)
}
