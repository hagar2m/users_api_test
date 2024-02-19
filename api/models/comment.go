package models

import "time"

type Comment struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"` // dynamic primary key
	UserID    uint      `gorm:"index;foreignKey:UserID" json:"user_id"`
	PostID    int       `gorm:"index;foreignKey:PostID" form:"post_id" json:"post_id"`
	UserName  string    `json:"user_name" form:"user_name"`                      // string for title
	Body      string    `json:"body" form:"body"`                                // string for body
	CreatedAt time.Time `gorm:"default:'2024-02-09 16:00:00'" json:"created_at"` // timestamp of creation
	Image     string    `json:"image_path" form:"-"`
}
