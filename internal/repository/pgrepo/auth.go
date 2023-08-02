package pgrepo

import (
	"context"
	"fmt"
	"github.com/khussa1n/task-management/internal/entity"
)

func (p *Postgres) CreateUser(ctx context.Context, u *entity.Users) error {
	query := fmt.Sprintf(`
			INSERT INTO %s (
			                username, -- 1 
			                first_name, -- 2
			                last_name, -- 3
			                hashed_password -- 4
			                )
			VALUES ($1, $2, $3, $4)
			`, usersTable)

	_, err := p.Pool.Exec(ctx, query, u.Email, u.FirstName, u.LastName, u.Password)
	if err != nil {
		return err
	}

	return nil
}
