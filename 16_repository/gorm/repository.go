package main

import (
	"errors"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Repository interface {
	Create(user User) (*User, error)
	GetByUsername(username string) (*User, error)
}

type repositoryMemory struct {
	users []User
}

func (r *repositoryMemory) Create(user User) (*User, error) {
	r.users = append(r.users, user)
	return &user, nil
}

func (r *repositoryMemory) GetByUsername(username string) (*User, error) {
	var found User
	for _, u := range r.users {
		if u.Username == username {
			found = u
			return &found, nil
		}
	}
	return nil, errors.New("user not found")
}

func NewInMemoryRepository(initial []User) Repository {
	return &repositoryMemory{
		users: initial,
	}
}

type repositoryGORM struct {
	DB *gorm.DB
}

func (r *repositoryGORM) Create(user User) (*User, error) {
	if err := r.DB.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repositoryGORM) GetByUsername(username string) (*User, error) {
	var found User
	if err := r.DB.Where("username = ?", username).First(&found).Error; err != nil {
		return nil, err
	}
	return &found, nil
}

func NewGORMRepository(db *gorm.DB) Repository {
	return &repositoryGORM{
		DB: db,
	}
}
