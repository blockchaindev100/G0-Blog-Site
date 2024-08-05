package service

import (
	"errors"
	"os"
	"time"

	"github.com/blockchaindev100/Go-Blog-Site/models"
	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(user *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"is_admin": user.Is_Admin,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyToken(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		return false, err
	}
	is_admin := false
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if claims["exp"].(float64) < float64(time.Now().Unix()) {
			return false, errors.New("token expired")
		}
		if claims["is_admin"] == true {
			is_admin = true
		}
	}
	if !token.Valid {
		return false, errors.New("invalid token")
	}
	return is_admin, nil
}
