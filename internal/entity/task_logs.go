package entity

import "time"

type TaskLogs struct {
	ID          int64      `json:"id" db:"id"`
	UserID      int64      `json:"user_id" db:"user_id"`
	TaskID      int64      `json:"task_id" db:"task_id" binding:"required"`
	StatusID    int64      `json:"status_id" db:"status_id" binding:"required"`
	BeginDate   time.Time  `json:"begin_date" db:"begin_date"`
	EndDate     *time.Time `json:"end_date" db:"end_date"`
	TotalHours  *int64     `json:"total_hours" db:"total_hours"`
	Description *string    `json:"description" db:"description"`
}
