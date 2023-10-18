package main

import (
	"com.test.users_api_test/app"
	"com.test.users_api_test/routing"
)

func main() {
	confg := app.LoadEnv()
	appConfig := app.CreateNewSqlClient(confg)
	routing.StartRouting(appConfig)
}
