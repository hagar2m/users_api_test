package models

import (
	"github.com/golang-jwt/jwt"
)

var JwtKey = []byte("my_secret_key")

type Claims struct {
	Email  string `json:"email"`
	UserId uint   `json:"user_id"`
	jwt.StandardClaims
}
