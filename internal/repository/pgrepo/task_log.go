package pgrepo

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/khussa1n/task-management/internal/entity"
)

func (p *Postgres) CreateTaskLog(ctx context.Context, tl *entity.TaskLogs) (*entity.TaskLogs, error) {
	taskLog := new(entity.TaskLogs)

	query := fmt.Sprintf(`
		INSERT INTO %s (
			user_id,
			task_id,
			status_id,
			begin_date,
			end_date,
			total_hours,
			description
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING *;
	`, taskLogsTable)

	err := pgxscan.Get(ctx, p.Pool, taskLog, query,
		tl.UserID, tl.TaskID, tl.StatusID, tl.BeginDate,
		tl.EndDate, tl.TotalHours, tl.Description)
	if err != nil {
		return nil, err
	}

	return taskLog, nil
}

func (p *Postgres) GetAllTaskLogsByTaskID(ctx context.Context, taskID int64) ([]entity.TaskLogs, error) {
	var taskLogs []entity.TaskLogs

	query := fmt.Sprintf(`
		SELECT * FROM %s
		WHERE task_id = $1
	`, taskLogsTable)

	err := pgxscan.Select(ctx, p.Pool, &taskLogs, query, taskID)
	if err != nil {
		return nil, err
	}

	return taskLogs, nil
}

func (p *Postgres) UpdateTaskLog(ctx context.Context, tl *entity.TaskLogs) (*entity.TaskLogs, error) {
	updatedTaskLog := new(entity.TaskLogs)

	query := fmt.Sprintf(`
		UPDATE %s
		SET
			user_id = $1,
			status_id = $2,
			begin_date = $3,
			end_date = $4,
			total_hours = $5,
			description = $6
		WHERE task_id = $7
		RETURNING *;
	`, taskLogsTable)

	err := pgxscan.Get(ctx, p.Pool, updatedTaskLog, query,
		tl.UserID, tl.StatusID, tl.BeginDate,
		tl.EndDate, tl.TotalHours, tl.Description, tl.TaskID)
	if err != nil {
		return nil, err
	}

	return updatedTaskLog, nil
}
