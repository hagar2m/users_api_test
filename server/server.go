package server

import (

	"com.test.users_api_test/configs"
	"com.test.users_api_test/db"
	"com.test.users_api_test/routing"
)

func Start() {
	configs.LoadViber()

	dbUrl := configs.GetDatabaseUrl()
	db.CreateNewSqlClient(dbUrl)

	routing.StartRouting()

}
