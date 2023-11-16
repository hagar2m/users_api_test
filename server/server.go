package server

import (
	"com.test.users_api_test/db"
	"com.test.users_api_test/routing"
	"com.test.users_api_test/configs"
)

func Start() {
	go configs.LoadViber()
	go routing.StartRouting()
	
	dbUrl := configs.GetDatabaseUrl()
	go db.CreateNewSqlClient(dbUrl)
}
