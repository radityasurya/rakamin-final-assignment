package models

import (
	"time"
)

type Photo struct {
	ID        uint      `gorm:"primaryKey" json:"id,omitempty"`
	Title     string    `json:"title" form:"title"`
	Caption   string    `json:"caption" form:"caption"`
	PhotoUrl  string    `json:"photo_url" form:"photo_url"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at,omitempty"`
}

type CreatePhotoRequest struct {
	Title     string    `json:"title"  binding:"required"`
	Caption   string    `json:"caption" binding:"required"`
	PhotoUrl  string    `json:"photo_url" binding:"required"`
	UserID    string    `json:"user_id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type UpdatePhoto struct {
	Title     string    `json:"title,omitempty"`
	Caption   string    `json:"caption,omitempty"`
	PhotoUrl  string    `json:"photo_url,omitempty"`
	UserID    string    `json:"user_id,omitempty"`
	CreateAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
