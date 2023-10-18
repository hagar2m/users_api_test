package routing

import (
	"os"

	"com.test.users_api_test/app"
	"com.test.users_api_test/handler"
	"com.test.users_api_test/services"
	"github.com/gin-gonic/gin"
)

func StartRouting(appConfig *app.App) {
	userService := services.NewUserService(appConfig)
	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()

	routerWithoutToken := router.Group("/")
	appConfig.Router = routerWithoutToken
	InitNonTokenRoutes(appConfig, userHandler)

	// routerWithToken	 :=  AppRouting{
	// 	Router: router.Group("/t"),
	// }
	// InitTokenRoutes(routerWithToken, &gormDb)

	portString := os.Getenv("PORT")
	router.Run(":" + portString)
}
