package postgre

import (
	"TaskAPI/models"
	"context"

	"github.com/jackc/pgx/v5"
)

func GetTasks(ctx context.Context, conn *pgx.Conn) ([]models.Tasks, error) {
	var tasks []models.Tasks
	sqlQuery := `
	SELECT id,title,completed FROM tasks
	`
	rows, err := conn.Query(ctx, sqlQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		task := models.Tasks{}
		if err := rows.Scan(&task.Id, &task.Title, &task.Completed); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return tasks, nil
}
