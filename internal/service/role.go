package service

import (
	"context"
	"fmt"
	"github.com/khussa1n/task-management/internal/entity"
)

func (m *Manager) CreateRole(ctx context.Context, r *entity.Roles) (*entity.Roles, error) {
	role, err := m.Repository.CreateRole(ctx, r)
	if err != nil {
		return nil, fmt.Errorf("can not create roles: %w", err)
	}

	return role, nil
}

func (m *Manager) GetAllRoles(ctx context.Context) ([]entity.Roles, error) {
	roles, err := m.Repository.GetAllRoles(ctx)
	if err != nil {
		return nil, fmt.Errorf("can not get all roles: %w", err)
	}

	return roles, nil
}

func (m *Manager) DeleteRole(ctx context.Context, id int64) error {
	err := m.Repository.DeleteRole(ctx, id)
	if err != nil {
		return fmt.Errorf("can not delete roles: %w", err)
	}

	return nil
}
