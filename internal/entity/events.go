package entity

import "time"

type Events struct {
	ID          int64     `json:"id" db:"id"`
	UserID      int64     `json:"user_id" db:"user_id" binding:"required"`
	TaskID      int64     `json:"task_id" db:"task_id" binding:"required"`
	ActionID    int64     `json:"action_id" db:"action_id" binding:"required"`
	CreatedDate time.Time `json:"created_date" db:"created_date" binding:"required"`
}
