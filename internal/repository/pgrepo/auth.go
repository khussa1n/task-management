package pgrepo

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/khussa1n/task-management/internal/entity"
)

func (p *Postgres) CreateUser(ctx context.Context, u *entity.Users) (*entity.Users, error) {
	user := new(entity.Users)

	query := fmt.Sprintf(`
			INSERT INTO %s (
			                email, -- 1 
			                first_name, -- 2
			                last_name, -- 3
			                hashed_password -- 4
			                )
			VALUES ($1, $2, $3, $4) RETURNING *;
			`, usersTable)

	err := pgxscan.Get(ctx, p.Pool, user, query, u.Email, u.FirstName, u.LastName, u.Password)
	if err != nil {
		return nil, err
	}

	return user, nil
}
