package models

import (
	"time"

	"gorm.io/gorm"
)

type List struct {
	ID          uint
	Title       string `validate:"required"`
	Description string `validate:"required"`
	UserId      int    `validate:"required"`
	User        User
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}
