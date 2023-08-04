package service

import (
	"context"
	"fmt"
	"github.com/khussa1n/task-management/internal/entity"
)

func (m *Manager) CreateEvent(ctx context.Context, e *entity.Events) (*entity.Events, error) {
	taskLog, err := m.Repository.CreateEvent(ctx, e)
	if err != nil {
		return nil, fmt.Errorf("can not create evets: %w", err)
	}

	return taskLog, nil
}

func (m *Manager) GetAllEventsByTaskID(ctx context.Context, taskID int64) ([]entity.Events, error) {
	taskLogs, err := m.Repository.GetAllEventsByTaskID(ctx, taskID)
	if err != nil {
		return nil, fmt.Errorf("can not get all evets: %w", err)
	}

	return taskLogs, nil
}
