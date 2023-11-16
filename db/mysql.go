package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// type App struct {
// 	Router *gin.RouterGroup
// 	DB     *gorm.DB
// }

var DB *gorm.DB

func CreateNewSqlClient(DBURL string) {
	db, err := gorm.Open(mysql.Open(DBURL), &gorm.Config{})
	if err != nil {
		log.Fatal("can't connect to database")
	}
	DB = db
}
