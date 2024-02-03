package models

import "time"

type UserTable struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}


type Tabler interface {
	TableName() string
}

// TableName overrides the table name to `users` insted of `UserTable`
func (UserTable) TableName() string {
	return "users"
}