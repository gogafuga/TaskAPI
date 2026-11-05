package postgre

import (
	"context"

	"github.com/jackc/pgx/v5"
)

var id int

func CreateTask(ctx context.Context, conn *pgx.Conn, task string) (int, error) {
	sqlQuery := `
	INSERT INTO tasks(title,completed)
	VALUES($1,FALSE)
	RETURNING id;
	`
	if err := conn.QueryRow(ctx, sqlQuery, task).Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
