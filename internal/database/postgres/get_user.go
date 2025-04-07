package postgres

import (
	"context"
	"fmt"
)

func (conn *PgxConnect) GetUser(ctx context.Context, email string) (*User, error) {
	var user User

	q := `select 
			id, 
			email,
			password
		from 
			users 
		where email = $1`

	err := conn.DB.QueryRow(ctx, q, email).Scan(
		&user.Id,
		&user.Email,
		&user.Password,
	)

	if err != nil {
		return nil, err
	}

	fmt.Println(user, 123321)

	return &user, nil
}
