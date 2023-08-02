package repository

import (
	"context"
	"github.com/khussa1n/task-management/internal/entity"
)

type Authorization interface {
	CreateUser(ctx context.Context, u *entity.Users) error
	Login(ctx context.Context, username, password string) (string, error)
}

type User interface {
	GetUser(ctx context.Context, username string) (*entity.Users, error)
	//UpdateUser(ctx context.Context, u *entity.User) error
	//DeleteUser(ctx context.Context, id int64) error
}

type Repository interface {
	Authorization
	User
}
