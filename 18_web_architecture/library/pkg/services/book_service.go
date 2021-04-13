package services

import (
	models "github.com/marcoshuck/golang-course/18_web_architecture/library/pkg/model"
	"github.com/marcoshuck/golang-course/18_web_architecture/library/pkg/repositories"
)

type ReservationService interface {
	Borrow(isbn string) (*models.Reservation, error)
	Return(isbn string) (*models.Reservation, error)
}

type BookService interface {
	ReservationService
	Create(book models.Book) (*models.Book, error)
	List(limit, offset *int) ([]models.Book, error)
}

type bookService struct {
	repository repositories.BookRepository
}

func (b *bookService) Create(book models.Book) (*models.Book, error) {
	if err := book.Validate(); err != nil {
		return nil, err
	}

	created, err := b.repository.Create(book)
	if err != nil {
		return nil, err
	}

	return created, nil
}

func (b *bookService) List(limit, offset *int) ([]models.Book, error) {
	list, err := b.repository.List(limit, offset)

	if err != nil {
		return nil, err
	}

	return list, nil
}

func (b *bookService) Borrow(isbn string) (*models.Reservation, error) {
	if err := models.ValidateISBN(isbn); err != nil {
		return nil, err
	}

	book, err := b.repository.Get(isbn)
	if err != nil {
		return nil, err
	}

	reservation := book.ToReservation()

	reservation, err := b.reservationRepository.Create(reservation)
	if err != nil {
		return nil, err
	}

	return reservation, nil
}

func (b *bookService) Return(isbn string) (*models.Reservation, error) {
	if err := models.ValidateISBN(isbn); err != nil {
		return nil, err
	}
}

func NewBookService(repository repositories.BookRepository) BookService {
	return &bookService{
		repository: repository,
	}
}
