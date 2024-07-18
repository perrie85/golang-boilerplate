package models

import (
	"time"

	"gorm.io/gorm"
)

type List struct {
	ID          uint           `json:"id"`
	Title       string         `validate:"required" json:"title"`
	Description string         `validate:"required" json:"description"`
	UserId      int            `validate:"required" json:"user_id"`
	User        User           `json:"user"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}
