package service

import (
	"context"
	"fmt"
	"github.com/khussa1n/task-management/internal/entity"
)

func (m *Manager) CreatePriority(ctx context.Context, r *entity.Priorities) (*entity.Priorities, error) {
	priority, err := m.Repository.CreatePriority(ctx, r)
	if err != nil {
		return nil, fmt.Errorf("can not create priority: %w", err)
	}

	return priority, nil
}

func (m *Manager) GetAllPriorities(ctx context.Context) ([]entity.Priorities, error) {
	priority, err := m.Repository.GetAllPriorities(ctx)
	if err != nil {
		return nil, fmt.Errorf("can not get all priorities: %w", err)
	}

	return priority, nil
}

func (m *Manager) DeletePriority(ctx context.Context, id int64) error {
	err := m.Repository.DeletePriority(ctx, id)
	if err != nil {
		return fmt.Errorf("can not Action priority: %w", err)
	}

	return nil
}
