package services

import (
	"github.com/DropsWeb/todogo/internal/database/postgres"
	"github.com/DropsWeb/todogo/pkg/swagger/ops"
	"github.com/go-openapi/runtime/middleware"
)

func (s *Server) PostAuthRegister(params ops.PostAuthRegisterParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	tx, err := s.ConnDb.DB.Begin(ctx)

	if err != nil {
		return ops.NewPostAuthRegisterInternalServerError().WithPayload(err.Error())
	}

	userParams := postgres.CreateUserParams(params.Register)

	_, err = s.ConnDb.CreateUser(ctx, tx, userParams)

	// defer tx.Rollback(ctx)

	if err != nil {
		return ops.NewPostAuthRegisterInternalServerError().WithPayload(err.Error())
	}

	err = tx.Commit(ctx)

	if err != nil {
		return ops.NewPostAuthRegisterInternalServerError().WithPayload(err.Error())
	}

	return ops.NewPostAuthRegisterNoContent()
}
