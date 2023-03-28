package domain

import jwt "github.com/golang-jwt/jwt/v4"

type JwtClaims struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	ID       int    `json:"id"`
	jwt.RegisteredClaims
}
