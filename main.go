package main

import (
	"com.test.users_api_test/db"
	"com.test.users_api_test/routing"
)

func main() {
	dbUrl := db.LoadEnv()
	db.CreateNewSqlClient(dbUrl)
	routing.StartRouting()
}
