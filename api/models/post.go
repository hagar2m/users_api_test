package models

import (
	"time"
)

type Post struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"` // dynamic primary key
	UserID    uint      `gorm:"index;foreignKey:ID" json:"user_id"`
	Title     string    `json:"title" form:"title"`                              // string for title
	Body      string    `json:"body" form:"body"`                                // string for body
	CreatedAt time.Time `gorm:"default:'2024-02-09 16:00:00'" json:"created_at"` // timestamp of creation
	Image     string    `json:"image_path" form:"-"`
	// image     FileResponse `gorm:"image"`
}
