package entity

import "time"

type Tasks struct {
	ID           int64     `json:"id" db:"id"`
	UserID       int64     `json:"user_id" db:"user_id"`
	CreatedDate  time.Time `json:"created_date" db:"created_date"`
	TaskName     string    `json:"task_name" db:"task_name" binding:"required"`
	Description  string    `json:"description" db:"description" binding:"required"`
	StatusID     *int64    `json:"status_id" db:"status_id"`
	DeadlineFrom time.Time `json:"deadline_from" db:"deadline_from"`
	DeadlineTo   time.Time `json:"deadline_to" db:"deadline_to"`
	PriorityID   *int64    `json:"priority_id" db:"priority_id"`
	ParentTaskID *int64    `json:"parent_task_id" db:"parent_task_id"`
}
