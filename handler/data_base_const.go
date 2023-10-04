package handler

import (
	"database/sql"

	"gorm.io/gorm"
)

type GormDB struct {
	DB *gorm.DB
}

type APIConfig struct {
	DB *sql.DB
}
