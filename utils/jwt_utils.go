package utils

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"com.test.users_api_test/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := getTokenString(c.Writer, c.Request)
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return models.JwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(*models.Claims); ok && token.Valid {
			c.Set("email", claims.Email)
			c.Set("userId", claims.UserId)
		} else {
			c.Writer.WriteHeader(http.StatusUnauthorized)
			fmt.Printf("Invalid token claims\n")
		}

		c.Next()
	}
}

func getTokenString(w http.ResponseWriter, r *http.Request) string {
	tokenString := r.Header.Get("Authorization")
	return strings.TrimPrefix(tokenString, "Bearer ")
}

func GenerateToken(loginUserData models.UserTable) (string, error) {
	expirationTime := time.Now().Add(24 * 5 * time.Hour).Unix()

	// Create the JWT claims, which includes the username and expiry time
	claims := &models.Claims{
		Email:  loginUserData.Email,
		UserId: loginUserData.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
		},
	}
	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	return token.SignedString(models.JwtKey)
}
