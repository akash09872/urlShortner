package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(username string) (string, error) {
	var jwtKey = []byte(os.Getenv("JWT_KEY"))
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(24 * time.Hour).Unix(),
		},
	)
	tokenString, err := token.SignedString(jwtKey)
	return tokenString, err
}
