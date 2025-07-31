package entity

import (
	"time"
)

type Post struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" validate:"required,min=20"`
	Content   string    `json:"content" validate:"required,min=200"`
	Category  string    `json:"category" validate:"required,min=3"`
	Status    string    `json:"status" validate:"required,oneof=publish draft thrash"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
