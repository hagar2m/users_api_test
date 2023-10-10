package handler

import (
	"gorm.io/gorm"
)

type GormDB struct {
	DB *gorm.DB
}
