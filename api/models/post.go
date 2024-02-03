package models

import (
	"time"
)

type Post struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`          // dynamic primary key
	UserID    uint      `gorm:"index" json:"user_id"`                        // foreign key to users table
	Title     string    `json:"title" form:"title"`                          // string for title
	Body      string    `json:"body" form:"body"`                            // string for body
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"` // timestamp of creation
	// image     FileResponse `gorm:"image"`
}