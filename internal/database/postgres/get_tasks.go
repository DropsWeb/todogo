package postgres

import "context"

func (conn *PgxConnect) GetTasks(ctx context.Context, user_id interface{}) ([]Task, error) {
	var tasks []Task

	q := `select 
			id, 
			title,
			description
		from 
			tasks 
		where user_id = $1`

	rows, err := conn.DB.Query(ctx, q, user_id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var task Task
		err := rows.Scan(
			&task.Id,
			&task.Title,
			&task.Description,
		)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}
