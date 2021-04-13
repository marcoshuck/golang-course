package models

import (
	"gorm.io/gorm"
	"time"
)

type Reservation struct {
	gorm.Model
	BookID     uint
	ReturnedAt *time.Time
}
