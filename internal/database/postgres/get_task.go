package postgres

import "context"

func (conn *PgxConnect) GetTask(ctx context.Context, user_id interface{}, task_id string) (*Task, error) {
	var task Task

	q := `select 
			id, 
			title,
			description
		from 
			tasks
		where id = $1
		and user_id = $2
		`

	err := conn.DB.QueryRow(ctx, q, task_id, user_id).Scan(
		&task.Id,
		&task.Title,
		&task.Description,
	)

	if err != nil {
		return nil, err
	}

	return &task, nil
}
