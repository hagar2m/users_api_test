package routing

import (
	"com.test.users_api_test/api/constants"
	"com.test.users_api_test/configs"
	"github.com/gin-gonic/gin"
)

func StartRouting() {
	router := gin.Default()

	nonTokenGroup := router.Group("/")
	NonTokenRoutes(nonTokenGroup)

	tokenGroup := router.Group("/auth")
	TokenRoutes(tokenGroup)

	router.Static(constants.UploadRoot, constants.UploadPath)

	portString := configs.GetServerPort()
	router.Run(":" + portString)
}
