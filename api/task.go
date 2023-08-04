package api

import "github.com/khussa1n/task-management/internal/entity"

type TaskCreateRequest struct {
	entity.Tasks
}

type TaskUpdateRequest struct {
	entity.Tasks
}
