package db

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// type App struct {
// 	Router *gin.RouterGroup
// 	DB     *gorm.DB
// }

var DB *gorm.DB

func LoadEnv() string {
	godotenv.Load(".env")
	return os.Getenv("DB_URL")
}

func CreateNewSqlClient(DBURL string) {
	db, err := gorm.Open(mysql.Open(DBURL), &gorm.Config{})
	if err != nil {
		log.Fatal("can't connect to database")
	}
	DB = db
}
