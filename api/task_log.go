package api

import "github.com/khussa1n/task-management/internal/entity"

type TaskLogCreateRequest struct {
	entity.TaskLogs
}

type TaskLogUpdateRequest struct {
	entity.TaskLogs
}
