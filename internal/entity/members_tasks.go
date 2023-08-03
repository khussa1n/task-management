package entity

type MembersTasks struct {
	UserID int64 `json:"user_id" db:"user_id" binding:"required"`
	TaskID int64 `json:"task_id" db:"task_id" binding:"required"`
	RoleID int64 `json:"role_id" db:"role_id" binding:"required"`
}
