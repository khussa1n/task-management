package entity

import "time"

type TaskLogs struct {
	ID          int64     `json:"id" db:"id"`
	UserID      int64     `json:"user_id" db:"user_id" binding:"required"`
	TaskID      int64     `json:"task_id" db:"task_id" binding:"required"`
	BeginDate   time.Time `json:"begin_date" db:"begin_date" binding:"required"`
	EndDate     time.Time `json:"end_date" db:"end_date" binding:"required"`
	TotalHours  int64     `json:"total_hours" db:"total_hours" binding:"required"`
	Description string    `json:"description" db:"description" binding:"required"`
}
