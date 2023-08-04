package service

import (
	"context"
	"fmt"
	"github.com/khussa1n/task-management/internal/entity"
)

func (m *Manager) CreateTaskLog(ctx context.Context, tl *entity.TaskLogs) (*entity.TaskLogs, error) {
	taskLog, err := m.Repository.CreateTaskLog(ctx, tl)
	if err != nil {
		return nil, fmt.Errorf("can not create task_logs: %w", err)
	}

	return taskLog, nil
}

func (m *Manager) GetAllTaskLogsByTaskID(ctx context.Context, taskID int64) ([]entity.TaskLogs, error) {
	taskLogs, err := m.Repository.GetAllTaskLogsByTaskID(ctx, taskID)
	if err != nil {
		return nil, fmt.Errorf("can not get all tasks_logs: %w", err)
	}

	return taskLogs, nil
}

func (m *Manager) UpdateTaskLog(ctx context.Context, tl *entity.TaskLogs) (*entity.TaskLogs, error) {
	taskLog, err := m.Repository.UpdateTaskLog(ctx, tl)
	if err != nil {
		return nil, fmt.Errorf("can not update tasks_logs: %w", err)
	}

	return taskLog, nil
}
