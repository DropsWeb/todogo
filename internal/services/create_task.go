package services

import (
	"github.com/DropsWeb/todogo/internal/database/postgres"
	"github.com/DropsWeb/todogo/pkg/swagger/ops"
	"github.com/go-openapi/runtime/middleware"
)

func (s *Server) PostTasks(params ops.PostTasksParams, principal interface{}) middleware.Responder {
	userID, err := s.BearerAuth(params.HTTPRequest.Header.Get("Authorization"))
	if err != nil {
		return ops.NewPostAuthLoginUnauthorized()
	}

	ctx := params.HTTPRequest.Context()

	tx, err := s.ConnDb.DB.Begin(ctx)

	if err != nil {
		return ops.NewPostTasksInternalServerError().WithPayload(err.Error())
	}

	defer tx.Rollback(ctx)

	taskParams := postgres.CreateTaskParams(params.CreateTask)

	_, err = s.ConnDb.CreateTask(ctx, tx, taskParams, userID)

	if err != nil {
		return ops.NewPostTasksInternalServerError().WithPayload(err.Error())
	}

	err = tx.Commit(ctx)

	if err != nil {
		return ops.NewPostTasksInternalServerError().WithPayload(err.Error())
	}

	return ops.NewPatchTasksIDCreated()
}
