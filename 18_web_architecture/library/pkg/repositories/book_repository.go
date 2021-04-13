package repositories

import models "github.com/marcoshuck/golang-course/18_web_architecture/library/pkg/model"

type BookRepository interface {
	Create(book models.Book) (*models.Book, error)
	Get(isbn string) (*models.Book, error)
	List(limit, offset *int) ([]models.Book, error)
	Update(isbn string, b models.Book) error
}

func NewBookRepository() BookRepository {
	panic("not implemented")
}
