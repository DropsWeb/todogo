package services

import (
	"time"

	"github.com/DropsWeb/todogo/pkg/swagger/models"
	"github.com/DropsWeb/todogo/pkg/swagger/ops"
	"github.com/go-openapi/runtime/middleware"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

func (s *Server) PostAuthLogin(params ops.PostAuthLoginParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	app_key := viper.GetString("APP_KEY")

	user, err := s.ConnDb.GetUser(ctx, *params.UserData.Email)

	if err != nil {
		return ops.NewPostAuthLoginInternalServerError().WithPayload(err.Error())
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(*params.UserData.Password)); err != nil {
		return &ops.PostAuthLoginNotFound{}
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"authorized": true,
		"sub":        user.Id,
		"email":      user.Email,
		"exp":        time.Now().Add(time.Hour * 2).Unix(), // Expires in 2 hours
	})

	token, err := claims.SignedString([]byte(app_key))
	if err != nil {
		return ops.NewPostAuthLoginInternalServerError()
	}

	payload := models.LoginSuccess{
		Success: true,
		Token:   token,
	}

	return ops.NewPostAuthLoginOK().WithPayload(&payload)
}
