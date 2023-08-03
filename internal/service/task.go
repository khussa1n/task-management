package service

import (
	"context"
	"fmt"
	"github.com/khussa1n/task-management/internal/entity"
	"time"
)

func (m *Manager) CreateTask(ctx context.Context, t *entity.Tasks) (*entity.Tasks, error) {
	t.CreatedDate = time.Now()
	task, err := m.Repository.CreateTask(ctx, t)
	if err != nil {
		return nil, fmt.Errorf("can not create task: %w", err)
	}

	return task, nil
}

func (m *Manager) GetAllTasks(ctx context.Context, id int64) ([]entity.Tasks, error) {
	tasks, err := m.Repository.GetAllTasks(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("can not get all tasks: %w", err)
	}

	return tasks, nil
}
