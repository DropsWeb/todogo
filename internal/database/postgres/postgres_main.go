package postgres

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type PgxConnect struct {
	DB *pgxpool.Pool
}
