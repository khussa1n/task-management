package entity

import "time"

type Tasks struct {
	ID           int64     `json:"id" db:"id"`
	UserID       string    `json:"user_id" db:"user_id" binding:"required"`
	CreatedDate  time.Time `json:"created_date" db:"created_date" binding:"required"`
	Description  string    `json:"description" db:"description" binding:"required"`
	StatusID     int64     `json:"status_id" db:"status_id" binding:"required"`
	DeadlineFrom time.Time `json:"deadline_from" db:"deadline_from" binding:"required"`
	DeadlineTo   time.Time `json:"deadline_to" db:"deadline_to" binding:"required"`
	PriorityID   int64     `json:"priority_id" db:"priority_id" binding:"required"`
	ParentTaskID int64     `json:"parent_task_id" db:"parent_task_id" binding:"required"`
}
