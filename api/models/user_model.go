package models

import "time"

type UserTable struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Name      string    `json:"name"`
	Email     string    `json:"email" gorm:"type:varchar(255)"`
	Password  string    `json:"password"`
	CreatedAt time.Time `gorm:"default:'2024-02-09 16:00:00'"`
}


type Tabler interface {
	TableName() string
}

// TableName overrides the table name to `users` insted of `UserTable`
func (UserTable) TableName() string {
	return "users"
}