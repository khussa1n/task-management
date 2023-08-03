package service

import (
	"context"
	"fmt"
	"github.com/khussa1n/task-management/internal/entity"
	"github.com/khussa1n/task-management/pkg/util"
)

func (m *Manager) GetUserByID(ctx context.Context, id int64) (*entity.Users, error) {
	user, err := m.Repository.GetUserByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("can not get user: %w", err)
	}

	return user, nil
}

func (m *Manager) UpdateUser(ctx context.Context, u *entity.Users) (*entity.Users, error) {
	hashedPassword, err := util.HashPassword(u.Password)
	if err != nil {
		return nil, err
	}
	newPassword := u.Password
	u.Password = hashedPassword

	user, err := m.Repository.UpdateUser(ctx, u)
	if err != nil {
		return nil, fmt.Errorf("can not update user: %w", err)
	}

	user.Password = newPassword

	return user, nil
}

func (m *Manager) DeleteUser(ctx context.Context, id int64) error {
	err := m.Repository.DeleteUser(ctx, id)
	if err != nil {
		return fmt.Errorf("can not delete user: %w", err)
	}

	return nil
}
