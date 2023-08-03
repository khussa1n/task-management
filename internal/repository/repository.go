package repository

import (
	"context"
	"github.com/khussa1n/task-management/internal/entity"
)

type Authorization interface {
	CreateUser(ctx context.Context, u *entity.Users) (*entity.Users, error)
}

type User interface {
	GetUserByEmail(ctx context.Context, email string) (*entity.Users, error)
	GetUserByID(ctx context.Context, id int64) (*entity.Users, error)
	UpdateUser(ctx context.Context, u *entity.Users) (*entity.Users, error)
	DeleteUser(ctx context.Context, id int64) error
}

type Repository interface {
	Authorization
	User
}
