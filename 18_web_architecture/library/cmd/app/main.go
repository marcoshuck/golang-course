package main

import (
	"fmt"
	"github.com/marcoshuck/golang-course/18_web_architecture/library/pkg/repositories"
	"github.com/marcoshuck/golang-course/18_web_architecture/library/pkg/services"
)

func main() {
	r := repositories.NewBookRepository()
	s := services.NewBookService(r)

	reservation, err := s.Return("test")
	if err != nil {
		return
	}

	fmt.Println("Reservation:", reservation)
}
