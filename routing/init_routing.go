package routing

import (
	"os"

	"github.com/gin-gonic/gin"
)

func StartRouting() {
	router := gin.Default()

	nonTokenGroup := router.Group("/")
	NonTokenRoutes(nonTokenGroup)

	tokenGroup := router.Group("/t")
	TokenRoutes(tokenGroup)

	portString := os.Getenv("PORT")
	router.Run(":" + portString)
}
