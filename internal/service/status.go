package service

import (
	"context"
	"fmt"
	"github.com/khussa1n/task-management/internal/entity"
)

func (m *Manager) CreateStatus(ctx context.Context, s *entity.Statuses) (*entity.Statuses, error) {
	status, err := m.Repository.CreateStatus(ctx, s)
	if err != nil {
		return nil, fmt.Errorf("can not create status: %w", err)
	}

	return status, nil
}

func (m *Manager) GetAllStatuses(ctx context.Context) ([]entity.Statuses, error) {
	statuses, err := m.Repository.GetAllStatuses(ctx)
	if err != nil {
		return nil, fmt.Errorf("can not get status: %w", err)
	}

	return statuses, nil
}

func (m *Manager) DeleteStatus(ctx context.Context, id int64) error {
	err := m.Repository.DeleteStatus(ctx, id)
	if err != nil {
		return fmt.Errorf("can not delete status: %w", err)
	}

	return nil
}
