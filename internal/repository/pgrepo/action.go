package pgrepo

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/khussa1n/task-management/internal/entity"
)

func (p *Postgres) CreateAction(ctx context.Context, r *entity.Actions) (*entity.Actions, error) {
	action := new(entity.Actions)

	query := fmt.Sprintf(`INSERT INTO %s (action_name) VALUES ($1) RETURNING *;`, actionsTable)

	err := pgxscan.Get(ctx, p.Pool, action, query, r.ActionName)
	if err != nil {
		return nil, err
	}

	return action, nil
}

func (p *Postgres) GetAllActions(ctx context.Context) ([]entity.Actions, error) {
	var actions []entity.Actions

	query := fmt.Sprintf(`SELECT * FROM %s;`, actionsTable)

	err := pgxscan.Select(ctx, p.Pool, &actions, query)
	if err != nil {
		return nil, err
	}

	return actions, nil
}

func (p *Postgres) DeleteAction(ctx context.Context, id int64) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, actionsTable)

	_, err := p.Pool.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
