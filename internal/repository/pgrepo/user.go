package pgrepo

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/khussa1n/task-management/internal/entity"
	"strings"
)

func (p *Postgres) GetUserByEmail(ctx context.Context, email string) (*entity.Users, error) {
	user := new(entity.Users)

	query := fmt.Sprintf(`
			SELECT id, email, first_name, last_name, hashed_password FROM %s WHERE email = $1
			`, usersTable)

	err := pgxscan.Get(ctx, p.Pool, user, query, strings.TrimSpace(email))
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (p *Postgres) GetUserByID(ctx context.Context, id int64) (*entity.Users, error) {
	user := new(entity.Users)

	query := fmt.Sprintf(`
			SELECT id, email, first_name, last_name, hashed_password FROM %s WHERE id = $1
			`, usersTable)

	err := pgxscan.Get(ctx, p.Pool, user, query, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (p *Postgres) UpdateUser(ctx context.Context, u *entity.Users) (*entity.Users, error) {
	user := new(entity.Users)

	query := fmt.Sprintf(`
			UPDATE %s SET email = $2, first_name = $3, last_name = $4, hashed_password = $5  WHERE id = $1 RETURNING *;
			`, usersTable)

	err := pgxscan.Get(ctx, p.Pool, user, query, u.ID, u.Email, u.FirstName, u.LastName, u.Password)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (p *Postgres) DeleteUser(ctx context.Context, id int64) error {
	query := fmt.Sprintf(`
			DELETE FROM %s WHERE id = $1 RETURNING *;
			`, usersTable)

	_, err := p.Pool.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
