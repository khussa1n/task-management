package pgrepo

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/khussa1n/task-management/internal/entity"
)

func (p *Postgres) CreateStatus(ctx context.Context, s *entity.Statuses) (*entity.Statuses, error) {
	status := new(entity.Statuses)

	query := fmt.Sprintf(`INSERT INTO %s (status_name) VALUES ($1) RETURNING *;`, statusesTable)

	err := pgxscan.Get(ctx, p.Pool, status, query, s.StatusName)
	if err != nil {
		return nil, err
	}

	return status, nil
}

func (p *Postgres) GetAllStatuses(ctx context.Context) ([]entity.Statuses, error) {
	var statuses []entity.Statuses

	query := fmt.Sprintf(`SELECT * FROM %s;`, statusesTable)

	err := pgxscan.Select(ctx, p.Pool, &statuses, query)
	if err != nil {
		return nil, err
	}

	return statuses, nil
}

func (p *Postgres) DeleteStatus(ctx context.Context, id int64) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, statusesTable)

	_, err := p.Pool.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
