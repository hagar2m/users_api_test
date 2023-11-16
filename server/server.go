package server

import (
	"com.test.users_api_test/db"
	"com.test.users_api_test/routing"
)

func Start()  {
	dbUrl := db.GetDbUrlFromEnv()
	db.CreateNewSqlClient(dbUrl)
	routing.StartRouting()
}