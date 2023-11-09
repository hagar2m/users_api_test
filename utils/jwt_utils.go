package utils

import (
	// "context"
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"com.test.users_api_test/models"
	"github.com/golang-jwt/jwt"
)

func ValidateToken(w http.ResponseWriter, r *http.Request) (error, context.Context) {
	tokenString := getTokenString(w, r)
	token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return models.JwtKey, nil
	})

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return  fmt.Errorf("Error: %v", err), nil
	}

	if !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return  fmt.Errorf("Invalid token"), nil
	}

	if claims, ok := token.Claims.(*models.Claims); ok && token.Valid {
		ctx := context.WithValue(r.Context(), "email", claims.Email)
		ctx = context.WithValue(r.Context(), "userId", claims.UserId)
		return nil, ctx
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Printf("Invalid token claims\n")
		return  fmt.Errorf("nvalid token"), nil
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
