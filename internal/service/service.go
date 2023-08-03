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

type Task interface {
	CreateTask(ctx context.Context, u *entity.Tasks) (*entity.Tasks, error)
	GetAllTasks(ctx context.Context, id int64) ([]entity.Tasks, error)
	//GetTaskByID(ctx context.Context, id int64) (*entity.Tasks, error)
	//UpdateTask(ctx context.Context, u *entity.Tasks) (*entity.Tasks, error)
	//DeleteTask(ctx context.Context, id int64) error
}

type Status interface {
	CreateStatus(ctx context.Context, s *entity.Statuses) (*entity.Statuses, error)
	GetAllStatuses(ctx context.Context) ([]entity.Statuses, error)
	DeleteStatus(ctx context.Context, id int64) error
}

type Role interface {
	CreateRole(ctx context.Context, s *entity.Roles) (*entity.Roles, error)
	GetAllRoles(ctx context.Context) ([]entity.Roles, error)
	DeleteRole(ctx context.Context, id int64) error
}

type Service interface {
	Authorization
	User
	Task
	Status
	Role
}
