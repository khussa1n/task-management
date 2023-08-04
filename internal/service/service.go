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
	CreateTask(ctx context.Context, t *entity.Tasks) (*entity.Tasks, error)
	GetAllTasks(ctx context.Context, id int64) ([]entity.Tasks, error)
	UpdateTask(ctx context.Context, t *entity.Tasks) (*entity.Tasks, error)
	DeleteTask(ctx context.Context, id int64) error
}

type Status interface {
	CreateStatus(ctx context.Context, s *entity.Statuses) (*entity.Statuses, error)
	GetAllStatuses(ctx context.Context) ([]entity.Statuses, error)
	DeleteStatus(ctx context.Context, id int64) error
}

type Role interface {
	CreateRole(ctx context.Context, r *entity.Roles) (*entity.Roles, error)
	GetAllRoles(ctx context.Context) ([]entity.Roles, error)
	DeleteRole(ctx context.Context, id int64) error
}

type Priority interface {
	CreatePriority(ctx context.Context, p *entity.Priorities) (*entity.Priorities, error)
	GetAllPriorities(ctx context.Context) ([]entity.Priorities, error)
	DeletePriority(ctx context.Context, id int64) error
}

type Action interface {
	CreateAction(ctx context.Context, a *entity.Actions) (*entity.Actions, error)
	GetAllActions(ctx context.Context) ([]entity.Actions, error)
	DeleteAction(ctx context.Context, id int64) error
}

type Comment interface {
	CreateComment(ctx context.Context, c *entity.Comments) (*entity.Comments, error)
	GetAllComments(ctx context.Context, userID int64) ([]entity.Comments, error)
	DeleteComment(ctx context.Context, id int64) error
}

type TaskLog interface {
	CreateTaskLog(ctx context.Context, tl *entity.TaskLogs) (*entity.TaskLogs, error)
	GetAllTaskLogsByTaskID(ctx context.Context, taskID int64) ([]entity.TaskLogs, error)
	UpdateTaskLog(ctx context.Context, tl *entity.TaskLogs) (*entity.TaskLogs, error)
}

type Event interface {
	CreateEvent(ctx context.Context, e *entity.Events) (*entity.Events, error)
	GetAllEventsByTaskID(ctx context.Context, taskID int64) ([]entity.Events, error)
}

type Service interface {
	Authorization
	User
	Task
	Status
	Role
	Priority
	Action
	Comment
	TaskLog
	Event
}
