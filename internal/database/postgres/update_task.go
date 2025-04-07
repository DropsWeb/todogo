package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func (conn *PgxConnect) UpdateTask(ctx context.Context, tx pgx.Tx, title string, description string, taskId string, userId interface{}) error {
	q := `UPDATE tasks
		SET 
			title = $1,
			description = $2
		WHERE
			id = $3
		AND 
			user_id = $4
	`
	_, err := tx.Exec(ctx, q, title, description, taskId, userId)

	if err != nil {
		return err
	}

	return nil
}
