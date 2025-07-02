package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKeyEnv = os.Getenv("JWT_TOKEN_SECRET_KEY")
var secretKey = []byte(jwtKeyEnv)

type CustomClaims struct {
	Role string `json:"role"`
	jwt.RegisteredClaims
}

func CreateToken(username string, role string) (string, error) {
	// Method HS256 is a JWT signature algorithm which uses a shared secret (symmetric key) to sign and validate the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"role":     role,
		"exp":      time.Now().Add(time.Hour * 1).Unix(), //expiration time (1 hour)
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (*CustomClaims, error) {
	// Parse and verify token
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
