package services

import (
	"github.com/DropsWeb/todogo/internal/database/postgres"
	"github.com/DropsWeb/todogo/pkg/swagger/models"
	"github.com/DropsWeb/todogo/pkg/swagger/ops"
	"github.com/go-openapi/runtime/middleware"
)

func (s *Server) GetTasks(params ops.GetTasksParams, principal interface{}) middleware.Responder {
	user_id, err := s.BearerAuth(params.HTTPRequest.Header.Get("Authorization"))
	if err != nil {
		return ops.NewPostAuthLoginUnauthorized()
	}
	tasks, err := s.ConnDb.GetTasks(params.HTTPRequest.Context(), user_id)

	if err != nil {
		return ops.NewGetTasksInternalServerError().WithPayload(err.Error())
	}

	result := convertToApiTasksCollection(tasks)

	return ops.NewGetTasksOK().WithPayload(result)
}

func convertToApiTasksCollection(data []postgres.Task) []*models.Task {

	var result []*models.Task

	for _, item := range data {
		task := &models.Task{
			ID:          item.Id,
			Title:       item.Title,
			Description: item.Description,
		}

		result = append(result, task)
	}

	return result
}
