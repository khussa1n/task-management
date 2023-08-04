package service

import (
	"context"
	"fmt"
	"github.com/khussa1n/task-management/internal/entity"
)

func (m *Manager) CreateAction(ctx context.Context, r *entity.Actions) (*entity.Actions, error) {
	action, err := m.Repository.CreateAction(ctx, r)
	if err != nil {
		return nil, fmt.Errorf("can not create Action: %w", err)
	}

	return action, nil
}

func (m *Manager) GetAllActions(ctx context.Context) ([]entity.Actions, error) {
	actions, err := m.Repository.GetAllActions(ctx)
	if err != nil {
		return nil, fmt.Errorf("can not get all Actions: %w", err)
	}

	return actions, nil
}

func (m *Manager) DeleteAction(ctx context.Context, id int64) error {
	err := m.Repository.DeleteAction(ctx, id)
	if err != nil {
		return fmt.Errorf("can not Action roles: %w", err)
	}

	return nil
}
