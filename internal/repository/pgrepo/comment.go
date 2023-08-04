package pgrepo

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/khussa1n/task-management/internal/entity"
)

func (p *Postgres) CreateComment(ctx context.Context, c *entity.Comments) (*entity.Comments, error) {
	comment := new(entity.Comments)

	query := fmt.Sprintf(`
		INSERT INTO %s (
			user_id,
			task_id,
			created_date,
			comment,
			parent_comment_id
		)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING *;
	`, commentsTable)

	err := pgxscan.Get(ctx, p.Pool, comment, query,
		c.UserID, c.TaskID, c.CreatedDate, c.Comment, c.ParentCommentID)
	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (p *Postgres) GetAllComments(ctx context.Context, userID int64) ([]entity.Comments, error) {
	var comments []entity.Comments

	query := fmt.Sprintf(`
		SELECT * FROM %s
		WHERE user_id = $1
	`, commentsTable)

	err := pgxscan.Select(ctx, p.Pool, &comments, query, userID)
	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (p *Postgres) DeleteComment(ctx context.Context, id int64) error {
	query := fmt.Sprintf(`
		DELETE FROM %s
		WHERE id = $1;
	`, commentsTable)

	_, err := p.Pool.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
