package models

import (
	"time"
)

type Post struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`          // dynamic primary key
	UserID    uint      `gorm:"index" json:"user_id"`                        // foreign key to users table
	Title     string    `gorm:"type:varchar(255)" json:"title"`              // string for title
	Body      string    `gorm:"type:text" json:"body"`                       // string for body
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"` // timestamp of creation
}
