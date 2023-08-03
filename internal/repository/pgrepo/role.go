package pgrepo

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/khussa1n/task-management/internal/entity"
)

func (p *Postgres) CreateRole(ctx context.Context, r *entity.Roles) (*entity.Roles, error) {
	role := new(entity.Roles)

	query := fmt.Sprintf(`INSERT INTO %s (role_name) VALUES ($1) RETURNING *;`, rolesTable)

	err := pgxscan.Get(ctx, p.Pool, role, query, r.RoleName)
	if err != nil {
		return nil, err
	}

	return role, nil
}

func (p *Postgres) GetAllRoles(ctx context.Context) ([]entity.Roles, error) {
	var roles []entity.Roles

	query := fmt.Sprintf(`SELECT * FROM %s;`, rolesTable)

	err := pgxscan.Select(ctx, p.Pool, &roles, query)
	if err != nil {
		return nil, err
	}

	return roles, nil
}

func (p *Postgres) DeleteRole(ctx context.Context, id int64) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, rolesTable)

	_, err := p.Pool.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
