package postgre

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CreateTable(ctx context.Context, conn *pgx.Conn) error {
	sqlQuery := `
	CREATE TABLE IF NOT EXISTS tasks(
	id SERIAL PRIMARY KEY,
	title TEXT NOT NULL,
	completed BOOLEAN NOT NULL DEFAULT FALSE
	);
	`
	_, err := conn.Exec(ctx, sqlQuery)
	return err
}
