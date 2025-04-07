package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func (conn *PgxConnect) CreateTask(ctx context.Context, tx pgx.Tx, taskParams CreateTaskParams, userID interface{}) (string, error) {
	var taskId string

	q := `
		INSERT INTO tasks (
			title,
			description,
			user_id
		) 
		VALUES (
			$1,
			$2,
			$3
		)
		RETURNING id
	`

	row := tx.QueryRow(ctx, q,
		taskParams.Title,
		taskParams.Description,
		userID,
	)

	err := row.Scan(&taskId)

	if err != nil {
		return "", err
	}

	return taskId, nil
}
