package main

import (
	"com.test.users_api_test/database"
	"com.test.users_api_test/routing"
)

func main() {
	dbUrl := database.LoadEnv()
	database.CreateNewSqlClient(dbUrl)
	routing.StartRouting()
}
