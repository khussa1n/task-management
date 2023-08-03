package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/khussa1n/task-management/internal/entity"
	"github.com/khussa1n/task-management/pkg/util"
)

func (m *Manager) CreateUser(ctx context.Context, u *entity.Users) (*entity.Users, error) {
	hashedPassword, err := util.HashPassword(u.Password)
	if err != nil {
		return nil, err
	}

	password := u.Password
	u.Password = hashedPassword

	user, err := m.Repository.CreateUser(ctx, u)
	if err != nil {
		return nil, err
	}

	user.Password = password

	return user, nil
}

func (m *Manager) Login(ctx context.Context, email, password string) (string, error) {
	user, err := m.Repository.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", fmt.Errorf("user_service not found")
		}
		return "", fmt.Errorf("get user_service error: %w", err)
	}

	err = util.CheckPassword(password, user.Password)
	if err != nil {
		return "", fmt.Errorf("incorrect password: %w", err)
	}

	accessToken, err := m.Token.CreatToken(user.ID, m.Config.Token.TimeToLive)
	if err != nil {
		return "", fmt.Errorf("create token error: %w", err)
	}

	return accessToken, nil
}

func (m *Manager) VerifyToken(token string) (int64, error) {
	payload, err := m.Token.ValidateToken(token)
	if err != nil {
		return 0, fmt.Errorf("Validate token error: %w", err)
	}

	return payload.UserID, nil
}
