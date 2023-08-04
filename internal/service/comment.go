package service

import (
	"context"
	"fmt"
	"github.com/khussa1n/task-management/internal/entity"
	"time"
)

func (m *Manager) CreateComment(ctx context.Context, c *entity.Comments) (*entity.Comments, error) {
	c.CreatedDate = time.Now()
	comment, err := m.Repository.CreateComment(ctx, c)
	if err != nil {
		return nil, fmt.Errorf("can not create comment: %w", err)
	}

	return comment, nil
}

func (m *Manager) GetAllComments(ctx context.Context, userID int64) ([]entity.Comments, error) {
	comment, err := m.Repository.GetAllComments(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("can not get all comment: %w", err)
	}

	return comment, nil
}

func (m *Manager) DeleteComment(ctx context.Context, id int64) error {
	err := m.Repository.DeleteComment(ctx, id)
	if err != nil {
		return fmt.Errorf("can not delete comment: %w", err)
	}

	return nil
}
