package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func (conn *PgxConnect) DeleteTask(ctx context.Context, tx pgx.Tx, taskId string) error {
	qs := `
		DELETE FROM tasks 
		WHERE 
			id = $1
	`
	_, err := tx.Exec(ctx, qs, taskId)

	if err != nil {
		return err
	}

	return nil
}
