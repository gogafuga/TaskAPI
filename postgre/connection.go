package postgre

import (
	"context"
	"errors"
	"os"

	"github.com/jackc/pgx/v5"
)

func ConnectionPSQl(ctx context.Context) (*pgx.Conn, error) {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		return nil, errors.New("DATABASE_URL is required")
	}
	conn, err := pgx.Connect(ctx, databaseURL)
	if err != nil {
		// Connection errors can contain credentials; do not expose them in logs.
		return nil, errors.New("database connection failed; check DATABASE_URL and PostgreSQL")
	}
	return conn, nil
}
