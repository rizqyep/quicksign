package utils

import (
	"fmt"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"github.com/rizqyep/quicksign/domain"
)

func CreateJWTToken(user *domain.User) (string, error) {
	err := godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("failed load file environment")
	} else {
		fmt.Println("successfully read file environment")
	}

	expiry, _ := time.ParseDuration(os.Getenv("JWT_EXPIRY_HOUR"))
	secret := os.Getenv("JWT_SECRET")

	exp := &jwt.NumericDate{time.Now().Add(time.Hour * expiry)}
	claims := &domain.JwtClaims{
		Username: user.Username,
		Email:    user.Email,
		ID:       user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: exp,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed_token, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return signed_token, err
}

func IsAuthorized(requestToken string) (bool, error) {
	err := godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("failed load file environment")
	} else {
		fmt.Println("successfully read file environment")
	}

	secret := os.Getenv("JWT_SECRET")

	_, err = jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

func ParseTokenData(requestToken string) (jwt.MapClaims, error) {
	err := godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("failed load file environment")
	} else {
		fmt.Println("successfully read file environment")
	}

	secret := os.Getenv("JWT_SECRET")
	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return jwt.MapClaims{}, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	fmt.Println("CLAIMS : ", claims)
	if !ok && !token.Valid {
		return jwt.MapClaims{}, fmt.Errorf("Invalid Token")
	}

	return claims, nil
}
