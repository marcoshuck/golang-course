package main

import (
	"errors"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID     string `json:"uuid"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (u *User) ValidatePassword() error {
	if len(u.Password) == 0 {
		return errors.New("password too short")
	}
	if len(u.Password) > 64 {
		return errors.New("password too long")
	}
	return nil
}
