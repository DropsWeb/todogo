package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

func (conn *PgxConnect) CreateUser(ctx context.Context, tx pgx.Tx, params CreateUserParams) (int, error) {
	var UserID int

	password, err := generatePasswordHash(*params.Password)

	if err != nil {
		return 0, err
	}

	q := `
	INSERT INTO users (
		email,
		password
	)
		VALUES
		(
		$1,
		$2
		)
	`

	row := tx.QueryRow(ctx, q,
		params.Email,
		password,
	)

	row.Scan(&UserID)

	return UserID, nil
}

func generatePasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password due to error %w", err)
	}
	return string(hash), nil
}
