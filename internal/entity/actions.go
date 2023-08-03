package entity

type Actions struct {
	ID         int64  `json:"id" db:"id"`
	ActionName string `json:"action_name" db:"action_name" binding:"required"`
}
