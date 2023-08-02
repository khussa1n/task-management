package pgrepo

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/khussa1n/task-management/internal/entity"
	"strings"
)

func (p *Postgres) GetUser(ctx context.Context, username string) (*entity.Users, error) {
	user := new(entity.Users)

	query := fmt.Sprintf(`
			SELECT user_id, username, first_name, last_name, hashed_password FROM %s WHERE username = $1
			`, usersTable)

	err := pgxscan.Get(ctx, p.Pool, user, query, strings.TrimSpace(username))
	if err != nil {
		return nil, err
	}

	return user, nil
}
