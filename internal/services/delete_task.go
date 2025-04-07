package services

import (
	"fmt"

	"github.com/DropsWeb/todogo/pkg/swagger/ops"
	"github.com/go-openapi/runtime/middleware"
)

func (s *Server) DeleteTasksID(params ops.DeleteTasksIDParams, principal interface{}) middleware.Responder {
	_, err := s.BearerAuth(params.HTTPRequest.Header.Get("Authorization"))
	if err != nil {
		return ops.NewPostAuthLoginUnauthorized()
	}

	ctx := params.HTTPRequest.Context()

	tx, err := s.ConnDb.DB.Begin(ctx)

	if err != nil {
		fmt.Println(err.Error())
	}

	err = s.ConnDb.DeleteTask(ctx, tx, params.ID)

	defer tx.Rollback(ctx)

	if err != nil {
		fmt.Println(err.Error())
	}

	err = tx.Commit(ctx)

	if err != nil {
		fmt.Println(err.Error())
	}

	return ops.NewDeleteTasksIDNoContent()
}
