package main

import "gorm.io/gorm"

type Repository interface {
	Get(uuid string) (*User, error)
}

type userRepository struct {
	db *gorm.DB
}

func (r *userRepository) Get(uuid string) (*User, error) {
	var result User
	err := r.db.Model(&User{}).Where("uuid = ?", uuid).First(&result).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func NewSQLRepository(db *gorm.DB) Repository {
	return &userRepository{
		db: db,
	}
}
