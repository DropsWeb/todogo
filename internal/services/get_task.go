package services

import (
	"github.com/DropsWeb/todogo/internal/database/postgres"
	"github.com/DropsWeb/todogo/pkg/swagger/models"
	"github.com/DropsWeb/todogo/pkg/swagger/ops"
	"github.com/go-openapi/runtime/middleware"
)

func (s *Server) GetTasksID(params ops.GetTasksIDParams, principal interface{}) middleware.Responder {
	user_id, err := s.BearerAuth(params.HTTPRequest.Header.Get("Authorization"))
	if err != nil {
		return ops.NewPostAuthLoginUnauthorized()
	}
	task, err := s.ConnDb.GetTask(params.HTTPRequest.Context(), user_id, params.ID)

	if err != nil {
		return ops.NewGetTasksIDNotFound()
	}

	result := convertToApiTaskResource(task)

	return ops.NewGetTasksIDOK().WithPayload(result)
}

func convertToApiTaskResource(data *postgres.Task) *models.Task {
	result := models.Task{
		ID:          data.Id,
		Title:       data.Title,
		Description: data.Description,
	}

	return &result
}
