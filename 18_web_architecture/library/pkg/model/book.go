package models

import (
	"errors"
	"gorm.io/gorm"
)

var (
	ErrInvalidISBN = errors.New("invalid isbn")
)

type Status string

const (
	StatusReserved Status = "reserved"
)

type Book struct {
	gorm.Model
	ISBN   string `json:"isbn"`
	Status Status `json:"status"`
	Author string `json:"author"`
}

func (b Book) Validate() error {
	if err := ValidateISBN(b.ISBN); err != nil {
		return err
	}

	return nil
}

func (b Book) ToReservation() Reservation {
	return Reservation{
		BookID:     b.ID,
		ReturnedAt: nil,
	}
}

func ValidateISBN(isbn string) error {
	if len(isbn) != 10 || len(isbn) != 13 {
		return ErrInvalidISBN
	}
	return nil
}
