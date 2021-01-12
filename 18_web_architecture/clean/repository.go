package main

import (
	"errors"
	"gorm.io/gorm"
)

type Repository interface {
	Get(uuid string) (*User, error)
}

type userRepositorySQL struct {
	db *gorm.DB
}

func (r *userRepositorySQL) Get(uuid string) (*User, error) {
	var result User
	err := r.db.Model(&User{}).Where("uuid = ?", uuid).First(&result).Error

	if err != nil {
		return nil, err
	}
	return &result, nil
}

func NewSQLRepository(db *gorm.DB) Repository {
	return &userRepositorySQL{
		db: db,
	}
}

type userRepositoryInMemory struct {
	users []User
}

func (r *userRepositoryInMemory) Get(uuid string) (*User, error) {
	for _, u := range r.users {
		if u.UUID == uuid {
			return &u, nil
		}
	}
	return nil, errors.New("user not found")
}

func NewInMemoryRepository(users []User) Repository {
	return &userRepositoryInMemory{
		users: users,
	}
}
