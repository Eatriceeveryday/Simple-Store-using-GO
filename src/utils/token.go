package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(id string, privateKey string) (string, error) {

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}).SignedString([]byte(privateKey))

	if err != nil {
		return "", fmt.Errorf("create: sign token: %w", err)
	}

	return token, nil
}

func ValidateToken(token string, privateKey []byte) (string, error) {
	tkn, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return privateKey, nil
	})

	if err != nil {
		return "", err
	}

	if !tkn.Valid {
		return "", fmt.Errorf("invalid token")
	}

	claims := tkn.Claims.(jwt.MapClaims)
	exp := claims["exp"].(float64)
	if exp <= float64(time.Now().Unix()) {
		return "", fmt.Errorf("expired token")
	}
	return claims["id"].(string), nil
}
