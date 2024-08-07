package service

import (
	"errors"
	"os"
	"time"

	logger "github.com/blockchaindev100/Go-Blog-Site/logger"
	"github.com/blockchaindev100/Go-Blog-Site/models"
	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(user *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.User_id,
		"is_admin": user.Is_Admin,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		logger.Logging().Error(err)
		return "", err
	}
	return tokenString, nil
}

func VerifyToken(tokenString string) (bool, string, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		logger.Logging().Error(err)
		return false, "", err
	}
	is_admin := false
	var user_id string
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if claims["exp"].(float64) < float64(time.Now().Unix()) {
			return false, user_id, errors.New("token expired")
		}
		if claims["is_admin"] == true {
			is_admin = true
		}
		if claims["user_id"] != "" {
			user_id = claims["user_id"].(string)
		}
	}
	if !token.Valid {
		logger.Logging().Error(errors.New("invalid token"))
		return false, user_id, errors.New("invalid token")
	}
	return is_admin, user_id, nil
}
