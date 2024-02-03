package db

import (
	"fmt"
	"log"

	"com.test.users_api_test/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// type App struct {
// 	Router *gin.RouterGroup
// 	DB     *gorm.DB
// }

var DB *gorm.DB

func CreateNewSqlClient() {
	USER := configs.GetDBUsername()
	PASS := configs.GetDBPassword()
	DB_HOST := configs.GetDBHost()
	DB_PORT := configs.GetDBPort()
	DBNAME := configs.GetDBName()

	//"root:12345678@(127.0.0.1:3306)"
	createDBDsn := fmt.Sprintf("%s:%s@(%s:%s)", USER, PASS, DB_HOST, DB_PORT)
	database, err := gorm.Open(mysql.Open(createDBDsn), &gorm.Config{})

	_ = database.Exec("CREATE DATABASE IF NOT EXISTS " + DBNAME + ";")

	// root:12345678@(127.0.0.1:3306)/sys
	dsn := fmt.Sprintf("%s/%s?charset=utf8&parseTime=True&loc=Local", createDBDsn, DBNAME)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("can't connect to database", err)
		panic(err.Error())
	}

	DB = db
}
