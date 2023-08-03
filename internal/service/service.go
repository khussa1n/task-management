package service

import (
	"context"
	"github.com/khussa1n/task-management/internal/entity"
)

type Authorization interface {
	CreateUser(ctx context.Context, u *entity.Users) (*entity.Users, error)
	Login(ctx context.Context, email, password string) (string, error)
	VerifyToken(token string) (int64, error)
}

type User interface {
	GetUserByID(ctx context.Context, id int64) (*entity.Users, error)
	UpdateUser(ctx context.Context, u *entity.Users) (*entity.Users, error)
	DeleteUser(ctx context.Context, id int64) error
}

type Service interface {
	Authorization
	User
}
