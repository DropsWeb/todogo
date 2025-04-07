package services

import (
	"github.com/DropsWeb/todogo/pkg/swagger/ops"
	"github.com/go-openapi/runtime/middleware"
)

func (s *Server) PatchTasksID(params ops.PatchTasksIDParams, principal interface{}) middleware.Responder {
	userId, err := s.BearerAuth(params.HTTPRequest.Header.Get("Authorization"))
	if err != nil {
		return ops.NewPostAuthLoginUnauthorized()
	}

	ctx := params.HTTPRequest.Context()

	tx, err := s.ConnDb.DB.Begin(ctx)

	if err != nil {
		return ops.NewPatchTasksIDInternalServerError().WithPayload(err.Error())
	}

	defer tx.Rollback(ctx)

	err = s.ConnDb.UpdateTask(ctx, tx, *params.UpdateTask.Title, *params.UpdateTask.Description, params.ID, userId)

	if err != nil {
		return ops.NewPatchTasksIDInternalServerError().WithPayload(err.Error())
	}

	err = tx.Commit(ctx)

	if err != nil {
		return ops.NewPatchTasksIDInternalServerError().WithPayload(err.Error())
	}

	return ops.NewPatchTasksIDCreated()
}
