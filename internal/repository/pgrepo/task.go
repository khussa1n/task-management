package pgrepo

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/khussa1n/task-management/internal/entity"
)

func (p *Postgres) CreateTask(ctx context.Context, t *entity.Tasks) (*entity.Tasks, error) {
	task := new(entity.Tasks)

	query := fmt.Sprintf(`
			INSERT INTO %s (
			                user_id, -- 1 
			                created_date, -- 2
			                task_name, -- 3
			                description, -- 4
			                status_id, -- 5
			                deadline_from, -- 6
			                deadline_to, -- 7
			                priority_id, -- 8
			                parent_task_id -- 9
			                )
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING *;
			`, tasksTable)

	err := pgxscan.Get(ctx, p.Pool, task, query,
		t.UserID, t.CreatedDate, t.TaskName, t.Description,
		t.StatusID, t.DeadlineFrom, t.DeadlineTo, t.PriorityID, t.ParentTaskID)
	if err != nil {
		return nil, err
	}

	membersTask := new(entity.MembersTasks)

	membersQuery := fmt.Sprintf(`
			INSERT INTO %s (
			                user_id, -- 1 
			                task_id, -- 2
			                role_id -- 3
			                )
			VALUES ($1, $2, $3) RETURNING *;
			`, membersTasksTable)

	err = pgxscan.Get(ctx, p.Pool, membersTask, membersQuery, task.UserID, task.ID, 1)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (p *Postgres) GetAllTasks(ctx context.Context, userID int64) ([]entity.Tasks, error) {
	var tasks []entity.Tasks

	query := fmt.Sprintf(`
		SELECT t.id, t.user_id, t.created_date, t.task_name, t.description,
			   t.status_id, t.deadline_from, t.deadline_to, t.priority_id, t.parent_task_id
		FROM %s t
		INNER JOIN %s mt ON t.id = mt.task_id
		WHERE mt.user_id = $1
	`, tasksTable, membersTasksTable)

	err := pgxscan.Select(ctx, p.Pool, &tasks, query, userID)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}
