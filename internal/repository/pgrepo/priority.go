package pgrepo

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/khussa1n/task-management/internal/entity"
)

func (p *Postgres) CreatePriority(ctx context.Context, r *entity.Priorities) (*entity.Priorities, error) {
	priority := new(entity.Priorities)

	query := fmt.Sprintf(`INSERT INTO %s (priority_name) VALUES ($1) RETURNING *;`, prioritiesTable)

	err := pgxscan.Get(ctx, p.Pool, priority, query, r.PriorityName)
	if err != nil {
		return nil, err
	}

	return priority, nil
}

func (p *Postgres) GetAllPriorities(ctx context.Context) ([]entity.Priorities, error) {
	var priorities []entity.Priorities

	query := fmt.Sprintf(`SELECT * FROM %s;`, prioritiesTable)

	err := pgxscan.Select(ctx, p.Pool, &priorities, query)
	if err != nil {
		return nil, err
	}

	return priorities, nil
}

func (p *Postgres) DeletePriority(ctx context.Context, id int64) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, prioritiesTable)

	_, err := p.Pool.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
