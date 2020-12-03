package main

import (
	"fmt"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open()
	repo := NewGORMRepository(db)

	found, err := repo.GetByUsername("marcoshuck2")
	fmt.Println("User:", found)
	fmt.Println("Error:", err)
}
