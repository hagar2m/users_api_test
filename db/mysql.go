package db

import (
	"fmt"
	"log"

	"com.test.users_api_test/api/models"
	"com.test.users_api_test/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// type App struct {
// 	Router *gin.RouterGroup
// 	DB     *gorm.DB
// }

var DB *gorm.DB

func CreateNewSqlClient() {
	createdDBDsn := createDatabaseIfNotExist()
	db := openDataBase(createdDBDsn)
	makeDatabaseMigration(db)

	DB = db
}

func createDatabaseIfNotExist() string {
	USER := configs.GetDBUsername()
	PASS := configs.GetDBPassword()
	DB_HOST := configs.GetDBHost()
	DB_PORT := configs.GetDBPort()
	DBNAME := configs.GetDBName()

	//"root:12345678@(127.0.0.1:3306)"
	createDBDsn := fmt.Sprintf("%s:%s@(%s:%s)", USER, PASS, DB_HOST, DB_PORT)
	database, _ := gorm.Open(mysql.Open(createDBDsn), &gorm.Config{})
	database.Exec("CREATE DATABASE IF NOT EXISTS " + DBNAME + ";")
	return createDBDsn
}

func openDataBase(dsnName string) *gorm.DB {
	DBNAME := configs.GetDBName()

	// root:12345678@(127.0.0.1:3306)/sys
	dsn := fmt.Sprintf("%s/%s?charset=utf8&parseTime=True&loc=Local", dsnName, DBNAME)
	db, error := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if error != nil {
		log.Fatal("can't connect to database", error)
		panic(error.Error())
	}
	return db
}

func makeDatabaseMigration(db *gorm.DB) {
	err := db.AutoMigrate(&models.UserTable{}, &models.Post{}, &models.Comment{})
	if err != nil {
		log.Fatal("error in AutoMigrate with database", err)
	}
}
