package services

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

func (s *Server) BearerAuth(bearerHeader string) (interface{}, error) {
	bearerToken := strings.Split(bearerHeader, " ")[1]
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(bearerToken, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("error decoding token")
		}

		app_key := viper.GetString("APP_KEY")

		return []byte(app_key), nil
	})

	if err != nil {
		log.Errorf(err.Error())
		return "", err
	}
	if token.Valid {
		return claims["sub"].(string), nil
	}
	return "", errors.New("invalid token")
}
