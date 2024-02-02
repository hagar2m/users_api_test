package routing

import (
	"com.test.users_api_test/configs"
	"github.com/gin-gonic/gin"
)

func StartRouting() {
	router := gin.Default()

	nonTokenGroup := router.Group("/")
	NonTokenRoutes(nonTokenGroup)

	tokenGroup := router.Group("/auth")
	TokenRoutes(tokenGroup)

	portString := configs.GetPort()
	router.Run(":" + portString)
}
