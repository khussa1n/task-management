package pgrepo

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/khussa1n/task-management/internal/entity"
)

func (p *Postgres) CreateEvent(ctx context.Context, e *entity.Events) (*entity.Events, error) {
	event := new(entity.Events)

	query := fmt.Sprintf(`
		INSERT INTO %s (
			user_id,
			task_id,
			action_id,
			created_date
		)
		VALUES ($1, $2, $3, $4)
		RETURNING *;
	`, eventsTable)

	err := pgxscan.Get(ctx, p.Pool, event, query,
		e.UserID, e.TaskID, e.ActionID, e.CreatedDate)
	if err != nil {
		return nil, err
	}

	return event, nil
}

func (p *Postgres) GetAllEventsByTaskID(ctx context.Context, taskID int64) ([]entity.Events, error) {
	var events []entity.Events

	query := fmt.Sprintf(`
		SELECT * FROM %s
		WHERE task_id = $1
	`, eventsTable)

	err := pgxscan.Select(ctx, p.Pool, &events, query, taskID)
	if err != nil {
		return nil, err
	}

	return events, nil
}
