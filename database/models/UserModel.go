package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
